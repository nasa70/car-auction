package carauction_test

import (
	"testing"

	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/testutil/nullify"
	"github.com/nasa70/car-auction/x/carauction"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LotsList: []types.Lots{
			{
				LotId: 1,
			},
			{
				LotId: 2,
			},
		},
		LotsCount: 2,
		AssetsList: []types.Assets{
			{
				AssetId: 0,
			},
			{
				AssetId: 1,
			},
		},
		AssetsCount: 2,
		AuctionInfo: &types.AuctionInfo{
			FirstInQueueLotId: 37,
			LastInQueueLotId:  28,
			FirstFinishTime:   1,
		},
		LotsQueueList: []types.LotsQueue{
			{
				LotId: 1,
			},
			{
				LotId: 2,
			},
		},
		LotsQueueCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarauctionKeeper(t)
	carauction.InitGenesis(ctx, *k, genesisState)
	got := carauction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LotsList, got.LotsList)
	require.Equal(t, genesisState.LotsCount, got.LotsCount)
	require.ElementsMatch(t, genesisState.AssetsList, got.AssetsList)
	require.Equal(t, genesisState.AssetsCount, got.AssetsCount)
	require.Equal(t, genesisState.AuctionInfo, got.AuctionInfo)
	require.ElementsMatch(t, genesisState.LotsQueueList, got.LotsQueueList)
	require.Equal(t, genesisState.LotsQueueCount, got.LotsQueueCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
