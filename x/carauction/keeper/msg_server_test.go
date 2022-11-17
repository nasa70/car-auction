package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/x/carauction"
	"github.com/nasa70/car-auction/x/carauction/keeper"
	"github.com/nasa70/car-auction/x/carauction/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CarauctionKeeper(t)
	carauction.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

//func setupMsgServerCreateAsset(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
//	k, ctx := keepertest.CarauctionKeeper(t)
//	carauction.InitGenesis(ctx, *k, *types.DefaultGenesis())
//	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
//
