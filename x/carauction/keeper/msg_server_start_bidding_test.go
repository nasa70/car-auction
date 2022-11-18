package keeper_test

import (
	"testing"
	"time"

	"github.com/nasa70/car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestStartBidding(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithTwoAssetsAndOneLot(t)

	_, err := msgServer.StartBidding(context, &types.MsgStartBidding{
		Creator:    nasa,
		LotId:      1,
		FinishTime: uint64(time.Now().Add(2 * time.Hour).UTC().Unix()),
		Bid:        sdk.NewCoin("token", sdk.NewInt(100)),
	})
	require.NoError(t, err)

	lot, found := keeper.GetLots(sdk.UnwrapSDKContext(context), 1)
	require.True(t, found)
	require.Equal(t, nasa, lot.BidOwner)
}

func TestStartBiddingAnotherOwner(t *testing.T) {
	msgServer, _, context := setupMsgServerWithTwoAssetsAndOneLot(t)

	_, err := msgServer.StartBidding(context, &types.MsgStartBidding{
		Creator:    bob,
		LotId:      1,
		FinishTime: uint64(time.Now().Add(2 * time.Hour).UTC().Unix()),
		Bid:        sdk.NewCoin("token", sdk.NewInt(100)),
	})
	require.ErrorIs(t, err, types.ErrLotAnotherOwner)
}
