package keeper

import (
	"encoding/binary"
	"fmt"

	storetypes2 "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	availblob1 "github.com/vitwit/avail-da-module"
	"github.com/vitwit/avail-da-module/types"
)

const (
	READY_STATE     uint32 = 0
	PENDING_STATE   uint32 = 1
	IN_VOTING_STATE uint32 = 2
	FAILURE_STATE   uint32 = 3
)

func ParseStatus(status uint32) string {
	switch status {
	case READY_STATE:
		return "SUCCESS"
	case PENDING_STATE:
		return "PENDING"
	case IN_VOTING_STATE:
		return "IN_VOTING"
	case FAILURE_STATE:
		return "FAILUTE"
	default:
		return "UNKNOWN"
	}
}

func IsAlreadyExist(ctx sdk.Context, store storetypes2.KVStore, blocksRange types.Range) bool {
	pendingBlobStoreKey := availblob1.PendingBlobsStoreKey(blocksRange)
	blobStatus := store.Get(pendingBlobStoreKey)
	fmt.Println("blob status:", blobStatus, blobStatus == nil)
	if blobStatus == nil {
		return false
	}
	return true
}

func IsStateReady(store storetypes2.KVStore) bool {
	statusBytes := store.Get(availblob1.BlobStatusKey)
	fmt.Println("status bytes............", statusBytes)
	if statusBytes == nil || len(statusBytes) == 0 {
		return true
	}

	status := binary.BigEndian.Uint32(statusBytes)

	return status == READY_STATE
}

func GetStatusFromStore(store storetypes2.KVStore) uint32 {
	statusBytes := store.Get(availblob1.BlobStatusKey)

	if statusBytes == nil || len(statusBytes) == 0 {
		return READY_STATE
	}

	status := binary.BigEndian.Uint32(statusBytes)

	return status
}

func UpdateBlobStatus(ctx sdk.Context, store storetypes2.KVStore, status uint32) error {

	statusBytes := make([]byte, 4)

	binary.BigEndian.PutUint32(statusBytes, status)

	store.Set(availblob1.BlobStatusKey, statusBytes)
	return nil
}

func UpdateEndHeight(ctx sdk.Context, store storetypes2.KVStore, endHeight uint64) error {

	heightBytes := make([]byte, 8)

	binary.BigEndian.PutUint64(heightBytes, endHeight)

	store.Set(availblob1.NextHeightKey, heightBytes)
	return nil
}

func UpdateProvenHeight(ctx sdk.Context, store storetypes2.KVStore, endHeight uint64) error {

	heightBytes := make([]byte, 8)

	binary.BigEndian.PutUint64(heightBytes, endHeight)

	store.Set(availblob1.ProvenHeightKey, heightBytes)
	return nil
}
