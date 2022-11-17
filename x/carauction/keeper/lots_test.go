package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/testutil/nullify"
	"github.com/nasa70/car-auction/x/carauction/keeper"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/stretchr/testify/require"
)

func createNLots(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Lots {
	items := make([]types.Lots, n)
	for i := range items {
		items[i].LotId = keeper.AppendLots(ctx, items[i])
	}
	return items
}

func TestLotsGet(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLots(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLots(ctx, item.LotId)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLotsRemove(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLots(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLots(ctx, item.LotId)
		_, found := keeper.GetLots(ctx, item.LotId)
		require.False(t, found)
	}
}

func TestLotsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLots(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLots(ctx)),
	)
}

func TestLotsCount(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLots(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLotsCount(ctx))
}
