package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

func (k msgServer) CreateLot(goCtx context.Context, msg *types.MsgCreateLot) (*types.MsgCreateLotResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Assets) == 0 {
		return nil, types.ErrEmptyAssetsList
	}

	assets := make([]types.Assets, 0)

	for _, assetId := range msg.GetAssets() {
		asset, found := k.Keeper.GetAssets(ctx, assetId)
		if !found {
			return nil, types.ErrAssetNotFound
		}
		if asset.Owner != msg.Creator {
			return nil, types.ErrAssetAnotherOwner
		}
		if asset.LotId != 0 {
			return nil, types.ErrAssetAnotherLot
		}
		assets = append(assets, asset)
	}

	lot := types.Lots{
		LotId:      0,
		Owner:      msg.GetCreator(),
		StartTime:  0,
		FinishTime: 0,
		Opened:     false,
		Approved:   false,
		Bid:        sdk.NewCoin("token", sdk.NewInt(0)),
		BidOwner:   msg.GetCreator(),
		Assets:     msg.GetAssets(),
	}

	lotId := k.Keeper.AppendLots(ctx, lot)

	for _, asset := range assets {
		asset.LotId = lotId
		k.Keeper.SetAssets(ctx, asset)
	}

	return &types.MsgCreateLotResponse{LotId: lotId}, nil
}
