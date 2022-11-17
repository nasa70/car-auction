package keeper_test

import (
	//"context"
	"testing"

	//keepertest "github.com/nasa70/car-auction/testutil/keeper"
	//"github.com/nasa70/car-auction/x/carauction"
	//"github.com/nasa70/car-auction/x/carauction/keeper"
	"github.com/nasa70/car-auction/x/carauction/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	nasa  = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

//func setupMsgServerCreateAsset(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
//	k, ctx := keepertest.CarauctionKeeper(t)
//	carauction.InitGenesis(ctx, *k, *types.DefaultGenesis())
//	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
//}

func TestCreateAsset(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateAsset(context, &types.MsgCreateAsset{
		Creator:     nasa,
		Description: "asset-1 description",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateAssetResponse{
		AssetId: 1,
	}, *createResponse)
}

func TestCreateTwoAssets(t *testing.T) {

	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateAsset(context, &types.MsgCreateAsset{
		Creator:     nasa,
		Description: "asset-1 description",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateAssetResponse{
		AssetId: 1,
	}, *createResponse)

	createResponse, err = msgServer.CreateAsset(context, &types.MsgCreateAsset{
		Creator:     nasa,
		Description: "asset-2 description",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateAssetResponse{
		AssetId: 2,
	}, *createResponse)
}
