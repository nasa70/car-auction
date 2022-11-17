package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

const MaxDescriptionLenght = 255

func (k msgServer) CreateAsset(goCtx context.Context, msg *types.MsgCreateAsset) (*types.MsgCreateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Description) > MaxDescriptionLenght {
		msg.Description = string([]byte(msg.Description)[:MaxDescriptionLenght-1])
	}

	asset := types.Assets{
		AssetId:     0,
		Owner:       msg.Creator,
		Description: msg.Description,
		LotId:       0,
	}

	id := k.Keeper.AppendAssets(ctx, asset)

	return &types.MsgCreateAssetResponse{AssetId: id}, nil
}
