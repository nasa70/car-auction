package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

// GetLotsCount get the total number of lots
func (k Keeper) GetLotsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LotsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLotsCount set the total number of lots
func (k Keeper) SetLotsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LotsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLots appends a lots in the store with a new id and update the count
func (k Keeper) AppendLots(
	ctx sdk.Context,
	lots types.Lots,
) uint64 {
	// Create the lots
	count := k.GetLotsCount(ctx)

	// Set the ID of the appended value
	lots.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsKey))
	appendedValue := k.cdc.MustMarshal(&lots)
	store.Set(GetLotsIDBytes(lots.Id), appendedValue)

	// Update lots count
	k.SetLotsCount(ctx, count+1)

	return count
}

// SetLots set a specific lots in the store
func (k Keeper) SetLots(ctx sdk.Context, lots types.Lots) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsKey))
	b := k.cdc.MustMarshal(&lots)
	store.Set(GetLotsIDBytes(lots.Id), b)
}

// GetLots returns a lots from its id
func (k Keeper) GetLots(ctx sdk.Context, id uint64) (val types.Lots, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsKey))
	b := store.Get(GetLotsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLots removes a lots from the store
func (k Keeper) RemoveLots(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsKey))
	store.Delete(GetLotsIDBytes(id))
}

// GetAllLots returns all lots
func (k Keeper) GetAllLots(ctx sdk.Context) (list []types.Lots) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Lots
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLotsIDBytes returns the byte representation of the ID
func GetLotsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLotsIDFromBytes returns ID in uint64 format from a byte array
func GetLotsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
