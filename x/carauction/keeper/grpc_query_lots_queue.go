package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nasa70/car-auction/x/carauction/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LotsQueueAll(c context.Context, req *types.QueryAllLotsQueueRequest) (*types.QueryAllLotsQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lotsQueues []types.LotsQueue
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lotsQueueStore := prefix.NewStore(store, types.KeyPrefix(types.LotsQueueKey))

	pageRes, err := query.Paginate(lotsQueueStore, req.Pagination, func(key []byte, value []byte) error {
		var lotsQueue types.LotsQueue
		if err := k.cdc.Unmarshal(value, &lotsQueue); err != nil {
			return err
		}

		lotsQueues = append(lotsQueues, lotsQueue)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLotsQueueResponse{LotsQueue: lotsQueues, Pagination: pageRes}, nil
}

func (k Keeper) LotsQueue(c context.Context, req *types.QueryGetLotsQueueRequest) (*types.QueryGetLotsQueueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	lotsQueue, found := k.GetLotsQueue(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLotsQueueResponse{LotsQueue: lotsQueue}, nil
}
