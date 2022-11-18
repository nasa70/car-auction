package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/x/carauction"
	"github.com/nasa70/car-auction/x/carauction/keeper"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, *keeper.Keeper, context.Context) {
	k, ctx := keepertest.CarauctionKeeper(t)
	carauction.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), k, sdk.WrapSDKContext(ctx)
}

func setupMsgServerWithTwoAssetsAndOneLot(t testing.TB) (types.MsgServer, *keeper.Keeper, context.Context) {

	k, ctx := keepertest.CarauctionKeeper(t)
	carauction.InitGenesis(ctx, *k, *types.DefaultGenesis())
	msgServer := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	resp1, err := msgServer.CreateAsset(context, types.NewMsgCreateAsset(nasa, "asset-for-testing-1"))
	require.NoError(t, err)
	resp2, err := msgServer.CreateAsset(context, types.NewMsgCreateAsset(nasa, "asset-for-testing-2"))
	require.NoError(t, err)
	_, err = msgServer.CreateLot(context, types.NewMsgCreateLot(nasa, []uint64{resp1.AssetId, resp2.AssetId}))
	require.NoError(t, err)

	return keeper.NewMsgServerImpl(*k), k, sdk.WrapSDKContext(ctx)
}
