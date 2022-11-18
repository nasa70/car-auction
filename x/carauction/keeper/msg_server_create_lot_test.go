package keeper_test

import (
	"testing"

	"github.com/nasa70/car-auction/x/carauction/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCreateLot(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateLot(context, &types.MsgCreateLot{
		Creator: nasa,
		Assets:  []uint64{2, 3},
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateLotResponse{
		LotId: 1,
	}, *createResponse)
}
