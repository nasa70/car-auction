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

func (k Keeper) AssetsAll(c context.Context, req *types.QueryAllAssetsRequest) (*types.QueryAllAssetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var assetss []types.Assets
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	assetsStore := prefix.NewStore(store, types.KeyPrefix(types.AssetsKey))

	pageRes, err := query.Paginate(assetsStore, req.Pagination, func(key []byte, value []byte) error {
		var assets types.Assets
		if err := k.cdc.Unmarshal(value, &assets); err != nil {
			return err
		}

		assetss = append(assetss, assets)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAssetsResponse{Assets: assetss, Pagination: pageRes}, nil
}

func (k Keeper) Assets(c context.Context, req *types.QueryGetAssetsRequest) (*types.QueryGetAssetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	assets, found := k.GetAssets(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAssetsResponse{Assets: assets}, nil
}
