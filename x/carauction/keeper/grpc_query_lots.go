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

func (k Keeper) LotsAll(c context.Context, req *types.QueryAllLotsRequest) (*types.QueryAllLotsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lotss []types.Lots
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lotsStore := prefix.NewStore(store, types.KeyPrefix(types.LotsKey))

	pageRes, err := query.Paginate(lotsStore, req.Pagination, func(key []byte, value []byte) error {
		var lots types.Lots
		if err := k.cdc.Unmarshal(value, &lots); err != nil {
			return err
		}

		lotss = append(lotss, lots)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLotsResponse{Lots: lotss, Pagination: pageRes}, nil
}

func (k Keeper) Lots(c context.Context, req *types.QueryGetLotsRequest) (*types.QueryGetLotsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	lots, found := k.GetLots(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLotsResponse{Lots: lots}, nil
}
