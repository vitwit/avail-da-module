package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/vitwit/avail-da-module/x/cada/keeper"
	availtypes "github.com/vitwit/avail-da-module/x/cada/types"
)

const (
	OpWeightMsgUpdateBlobStatusRequest = "op_weight_msg_update_blob_status"

	DefaultWeightMsgUpdateStatusRequest = 100
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, txConfig client.TxConfig,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var weightMsgUpdateBlobStatusRequest int
	appParams.GetOrGenerate(OpWeightMsgUpdateBlobStatusRequest, &weightMsgUpdateBlobStatusRequest, nil, func(_ *rand.Rand) {
		weightMsgUpdateBlobStatusRequest = DefaultWeightMsgUpdateStatusRequest
	})

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgUpdateBlobStatusRequest,
			SimulateUpdateBlobStatus(txConfig, cdc, ak, bk, k),
		),
	}
}

func SimulateUpdateBlobStatus(txConfig client.TxConfig, cdc codec.JSONCodec, ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper, _ keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simaAccount, _ := simtypes.RandomAcc(r, accs)

		fromBlock := r.Uint64()%100 + 1
		toBlock := fromBlock + r.Uint64()%10

		isSuccess := r.Intn(2) == 1

		availHeight := r.Uint64()%100 + 1

		blockRange := availtypes.Range{
			From: fromBlock,
			To:   toBlock,
		}

		msg := availtypes.NewMsgUpdateBlobStatus(simaAccount.Address.String(), blockRange, availHeight, isSuccess)

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}
