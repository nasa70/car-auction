package types_test

import (
	"testing"

	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
						AssetId: 1,
					},
					{
						AssetId: 2,
					},
				},
				AssetsCount: 2,
				AuctionInfo: &types.AuctionInfo{
					FirstInQueueLotId: 0,
					LastInQueueLotId:  0,
					FirstFinishTime:   0,
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated lots",
			genState: &types.GenesisState{
				LotsList: []types.Lots{
					{
						LotId: 1,
					},
					{
						LotId: 1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid lots count",
			genState: &types.GenesisState{
				LotsList: []types.Lots{
					{
						LotId: 1,
					},
				},
				LotsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated assets",
			genState: &types.GenesisState{
				AssetsList: []types.Assets{
					{
						AssetId: 1,
					},
					{
						AssetId: 1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid assets count",
			genState: &types.GenesisState{
				AssetsList: []types.Assets{
					{
						AssetId: 1,
					},
				},
				AssetsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated lotsQueue",
			genState: &types.GenesisState{
				LotsQueueList: []types.LotsQueue{
					{
						LotId: 1,
					},
					{
						LotId: 1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid lotsQueue count",
			genState: &types.GenesisState{
				LotsQueueList: []types.LotsQueue{
					{
						LotId: 1,
					},
				},
				LotsQueueCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
