package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/testutil/nullify"
	"github.com/nasa70/car-auction/x/carauction/types"
)

func TestAuctionInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestAuctionInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAuctionInfoRequest
		response *types.QueryGetAuctionInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAuctionInfoRequest{},
			response: &types.QueryGetAuctionInfoResponse{AuctionInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.AuctionInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
