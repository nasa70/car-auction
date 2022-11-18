package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/nasa70/car-auction/testutil/keeper"
	"github.com/nasa70/car-auction/testutil/nullify"
	"github.com/nasa70/car-auction/x/carauction/types"
)

func TestLotsQueueQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLotsQueue(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLotsQueueRequest
		response *types.QueryGetLotsQueueResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLotsQueueRequest{Id: msgs[0].LotId},
			response: &types.QueryGetLotsQueueResponse{LotsQueue: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetLotsQueueRequest{Id: msgs[1].LotId},
			response: &types.QueryGetLotsQueueResponse{LotsQueue: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetLotsQueueRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LotsQueue(wctx, tc.request)
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

func TestLotsQueueQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CarauctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLotsQueue(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLotsQueueRequest {
		return &types.QueryAllLotsQueueRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LotsQueueAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LotsQueue), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LotsQueue),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LotsQueueAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LotsQueue), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LotsQueue),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LotsQueueAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LotsQueue),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LotsQueueAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
