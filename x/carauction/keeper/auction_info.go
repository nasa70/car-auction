package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

// SetAuctionInfo set auctionInfo in the store
func (k Keeper) SetAuctionInfo(ctx sdk.Context, auctionInfo types.AuctionInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKey))
	b := k.cdc.MustMarshal(&auctionInfo)
	store.Set([]byte{0}, b)
}

// GetAuctionInfo returns auctionInfo
func (k Keeper) GetAuctionInfo(ctx sdk.Context) (val types.AuctionInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuctionInfo removes auctionInfo from the store
func (k Keeper) RemoveAuctionInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKey))
	store.Delete([]byte{0})
}
