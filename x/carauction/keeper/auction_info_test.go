package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/testutil/nullify"
	"github.com/nasa70/car-auction/x/carauction/keeper"
	"github.com/nasa70/car-auction/x/carauction/types"
)

func createTestAuctionInfo(keeper *keeper.Keeper, ctx sdk.Context) types.AuctionInfo {
	item := types.AuctionInfo{}
	keeper.SetAuctionInfo(ctx, item)
	return item
}

func TestAuctionInfoGet(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	item := createTestAuctionInfo(keeper, ctx)
	rst, found := keeper.GetAuctionInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAuctionInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	createTestAuctionInfo(keeper, ctx)
	keeper.RemoveAuctionInfo(ctx)
	_, found := keeper.GetAuctionInfo(ctx)
	require.False(t, found)
}
