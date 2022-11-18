package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

// GetLotsQueueCount get the total number of lotsQueue
func (k Keeper) GetLotsQueueCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LotsQueueCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLotsQueueCount set the total number of lotsQueue
func (k Keeper) SetLotsQueueCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LotsQueueCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLotsQueue appends a lotsQueue in the store with a new id and update the count
func (k Keeper) AppendLotsQueue(
	ctx sdk.Context,
	lotsQueue types.LotsQueue,
) uint64 {
	// Create the lotsQueue
	count := k.GetLotsQueueCount(ctx)

	// Set the ID of the appended value
	lotsQueue.LotId = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsQueueKey))
	appendedValue := k.cdc.MustMarshal(&lotsQueue)
	store.Set(GetLotsQueueIDBytes(lotsQueue.LotId), appendedValue)

	// Update lotsQueue count
	k.SetLotsQueueCount(ctx, count+1)

	return count
}

// SetLotsQueue set a specific lotsQueue in the store
func (k Keeper) SetLotsQueue(ctx sdk.Context, lotsQueue types.LotsQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsQueueKey))
	b := k.cdc.MustMarshal(&lotsQueue)
	store.Set(GetLotsQueueIDBytes(lotsQueue.LotId), b)
}

// GetLotsQueue returns a lotsQueue from its id
func (k Keeper) GetLotsQueue(ctx sdk.Context, id uint64) (val types.LotsQueue, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsQueueKey))
	b := store.Get(GetLotsQueueIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLotsQueue removes a lotsQueue from the store
func (k Keeper) RemoveLotsQueue(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsQueueKey))
	store.Delete(GetLotsQueueIDBytes(id))
}

// GetAllLotsQueue returns all lotsQueue
func (k Keeper) GetAllLotsQueue(ctx sdk.Context) (list []types.LotsQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsQueueKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LotsQueue
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLotsQueueIDBytes returns the byte representation of the ID
func GetLotsQueueIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLotsQueueIDFromBytes returns ID in uint64 format from a byte array
func GetLotsQueueIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
