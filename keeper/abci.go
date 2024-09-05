package keeper

import (
	"bytes"
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StakeWeightedVotes struct {
	Votes              map[string]int64
	ExtendedCommitInfo abci.ExtendedCommitInfo
}

type ProofOfBlobProposalHandler struct {
	keeper *Keeper

	prepareProposalHandler sdk.PrepareProposalHandler
	processProposalHandler sdk.ProcessProposalHandler
	VoteExtHandler         VoteExtHandler
}

func NewProofOfBlobProposalHandler(
	k *Keeper,
	prepareProposalHandler sdk.PrepareProposalHandler,
	processProposalHandler sdk.ProcessProposalHandler,
	voteExtHandler VoteExtHandler,
) *ProofOfBlobProposalHandler {
	return &ProofOfBlobProposalHandler{
		keeper:                 k,
		prepareProposalHandler: prepareProposalHandler,
		processProposalHandler: processProposalHandler,
		VoteExtHandler:         voteExtHandler,
	}
}

func (h *ProofOfBlobProposalHandler) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	h.keeper.proposerAddress = req.ProposerAddress
	proposalTxs := req.Txs

	// h.voteExtHandler.ExtendVoteHandler()(nil,err)

	// resp, err := h.prepareProposalHandler(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }

	votes, err := h.aggregateVotes(ctx, req.LocalLastCommit)
	if err != nil {
		fmt.Println("error while aggregating votes", err)
		return nil, err
	}

	injectedVoteExtTx := StakeWeightedVotes{
		Votes:              votes,
		ExtendedCommitInfo: req.LocalLastCommit,
	}

	fmt.Println("votes..................", votes)
	fmt.Println("injectedVoteExtTx............", injectedVoteExtTx.ExtendedCommitInfo.String())

	bz, err := json.Marshal(injectedVoteExtTx)
	if err != nil {
		fmt.Println("failed to encode injected vote extension tx", "err", err)
		// return nil, errors.New("failed to encode injected vote extension tx")
	}

	proposalTxs = append(proposalTxs, bz)
	return &abci.ResponsePrepareProposal{
		Txs: proposalTxs,
	}, nil
}

func (h *ProofOfBlobProposalHandler) ProcessProposal(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	if len(req.Txs) == 0 {
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}

	var injectedVoteExtTx StakeWeightedVotes
	if err := json.Unmarshal(req.Txs[0], &injectedVoteExtTx); err != nil {
		fmt.Println("failed to decode injected vote extension tx", "err", err)
		// return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
	}

	fmt.Println("injected data is:..............", injectedVoteExtTx)
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil

}

func (k *Keeper) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) error {
	fmt.Println("coming hereee.........", ctx.BlockHeight())

	currentBlockHeight := ctx.BlockHeight()
	if !k.IsValidBlockToPostTODA(uint64(currentBlockHeight)) {
		return nil
	}

	// fmt.Printf("Ctx.........%+v\n", ctx)

	fmt.Println("block heighttt.........", ctx.BlockHeight(), ctx.ExecMode(), ctx.IsCheckTx(), ctx.IsReCheckTx())

	provenHeight := k.GetProvenHeightFromStore(ctx)
	fromHeight := provenHeight + 1
	fmt.Println("from height..", fromHeight)
	endHeight := min(fromHeight+uint64(k.MaxBlocksForBlob), uint64(ctx.BlockHeight())) //exclusive i.e [fromHeight, endHeight)
	fmt.Println("end height..", endHeight-1)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	err := k.SetBlobStatusPending(sdkCtx, fromHeight, endHeight-1)
	if err != nil {
		fmt.Println("error while setting blob status...", err)
		return nil
	}

	var blocksToSumit []int64

	for i := fromHeight; i < endHeight; i++ {
		blocksToSumit = append(blocksToSumit, int64(i))
	}

	fmt.Println("blocks to submittttttttt.........", blocksToSumit)

	// only proposar should should run the this
	if bytes.Equal(req.ProposerAddress, k.proposerAddress) {
		// update blob status to success
		// err = k.SetBlobStatusSuccess(sdkCtx, fromHeight, endHeight)
		// if err != nil {
		// 	return nil
		// }

		// Todo: run the relayer routine
		// relayer doesn't have to make submitBlob Transaction, it should just start DA submission
		k.relayer.PostBlocks(ctx, blocksToSumit, k.cdc, req.ProposerAddress)

	}

	return nil
}

func (k *Keeper) IsValidBlockToPostTODA(height uint64) bool {
	if uint64(height) <= uint64(1) {
		return false
	}

	if (height-1)%k.PublishToAvailBlockInterval != 0 {
		return false
	}

	return true
}

func (h *ProofOfBlobProposalHandler) aggregateVotes(ctx sdk.Context, ci abci.ExtendedCommitInfo) (map[string]int64, error) {
	from := h.keeper.GetStartHeightFromStore(ctx)
	to := h.keeper.GetEndHeightFromStore(ctx)

	pendingRangeKey := Key(from, to)
	votes := make(map[string]int64, 1)

	var totalStake int64

	for _, v := range ci.Votes {
		// Process only votes with BlockIDFlagCommit, indicating the validator committed to the block.
		// Skip votes with other flags (e.g., BlockIDFlagUnknown, BlockIDFlagNil).
		if v.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue
		}

		var voteExt VoteExtension
		if err := json.Unmarshal(v.VoteExtension, &voteExt); err != nil {
			h.VoteExtHandler.logger.Error("failed to decode vote extension", "err", err, "validator", fmt.Sprintf("%x", v.Validator.Address))
			//return nil, err
		}

		totalStake += v.Validator.Power

		for voteRange, isVoted := range voteExt.Votes {
			if voteRange != pendingRangeKey || !isVoted {
				continue
			}

			votes[voteRange] += v.Validator.Power
		}

	}
	return votes, nil
}
