package relayer

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vitwit/avail-da-module/relayer/avail"
)

// SubmitDataToAvailClient submits data to the Avail client using an HTTP POST request.
func (r *Relayer) SubmitDataToAvailClient(data []byte, blocks []int64) (avail.BlockMetaData, error) {
	blockInfo, err := r.AvailDAClient.Submit(data)
	if err != nil {
		r.Logger.Error("Error while posting block(s) to Avail DA",
			"height_start", blocks[0],
			"height_end", blocks[len(blocks)-1],
			"appID", strconv.Itoa(r.AvailConfig.AppID),
		)

		return blockInfo, err
	}

	r.Logger.Info("Successfully posted block(s) to Avail DA",
		"height_start", blocks[0],
		"height_end", blocks[len(blocks)-1],
		"appID", strconv.Itoa(r.AvailConfig.AppID),
		"block_hash", blockInfo.BlockHash,
		"block_number", blockInfo.BlockNumber,
		"hash", blockInfo.Hash,
	)

	return blockInfo, nil
}

// IsDataAvailable is to query the avail light client and check if the data is made available at the given height
func (r *Relayer) IsDataAvailable(ctx sdk.Context, from, to, availHeight uint64) (bool, error) {
	var blocks []int64
	for i := from; i <= to; i++ {
		blocks = append(blocks, int64(i))
	}

	cosmosBlocksData := r.GetBlocksDataFromLocal(ctx, blocks)

	return r.AvailDAClient.IsDataAvailable(cosmosBlocksData, int(availHeight))
}

// GetBlock represents the data structure for a block with its associated transactions.
type GetBlock struct {
	BlockNumber      int      `json:"block_number"`
	DataTransactions []string `json:"data_transactions"`
}
