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

func createNLotsQueue(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LotsQueue {
	items := make([]types.LotsQueue, n)
	for i := range items {
		items[i].LotId = keeper.AppendLotsQueue(ctx, items[i])
	}
	return items
}

func TestLotsQueueGet(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLotsQueue(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLotsQueue(ctx, item.LotId)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLotsQueueRemove(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLotsQueue(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLotsQueue(ctx, item.LotId)
		_, found := keeper.GetLotsQueue(ctx, item.LotId)
		require.False(t, found)
	}
}

func TestLotsQueueGetAll(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLotsQueue(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLotsQueue(ctx)),
	)
}

func TestLotsQueueCount(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	items := createNLotsQueue(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLotsQueueCount(ctx))
}
