package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

const MaxDescriptionLenght = 255

func (k msgServer) CreateAsset(goCtx context.Context, msg *types.MsgCreateAsset) (*types.MsgCreateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	aInfo, found := k.Keeper.GetAuctionInfo(ctx)
	if !found {
		panic("AuctionInfo not found")
	}

	_, found = k.Keeper.GetAssets(ctx, aInfo.NextAssetId)
	if found {
		return nil, types.ErrAssetAlreadyExists
	}

	if len(msg.Description) > MaxDescriptionLenght {
		msg.Description = string([]byte(msg.Description)[:MaxDescriptionLenght-1])
	}

	asset := types.Assets{
		AssetId:     aInfo.NextAssetId,
		Owner:       msg.Creator,
		Description: msg.Description,
		LotId:       0,
	}

	aInfo.NextAssetId++
	k.Keeper.SetAuctionInfo(ctx, aInfo)
	k.Keeper.SetAssets(ctx, asset)

	return &types.MsgCreateAssetResponse{AssetId: asset.AssetId}, nil
}
