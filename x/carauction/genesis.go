package carauction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/keeper"
	"github.com/nasa70/car-auction/x/carauction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the lots
	for _, elem := range genState.LotsList {
		k.SetLots(ctx, elem)
	}

	// Set lots count
	k.SetLotsCount(ctx, genState.LotsCount)
	// Set all the assets
	for _, elem := range genState.AssetsList {
		k.SetAssets(ctx, elem)
	}

	// Set assets count
	k.SetAssetsCount(ctx, genState.AssetsCount)
	// Set if defined
	//if genState.AuctionInfo != nil {
	k.SetAuctionInfo(ctx, *genState.AuctionInfo)
	//}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LotsList = k.GetAllLots(ctx)
	genesis.LotsCount = k.GetLotsCount(ctx)
	genesis.AssetsList = k.GetAllAssets(ctx)
	genesis.AssetsCount = k.GetAssetsCount(ctx)
	// Get all auctionInfo
	auctionInfo, found := k.GetAuctionInfo(ctx)
	if found {
		genesis.AuctionInfo = &auctionInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
