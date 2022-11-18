package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
)

const minBiddingDuration = time.Hour

func (k msgServer) StartBidding(goCtx context.Context, msg *types.MsgStartBidding) (*types.MsgStartBiddingResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	lot, found := k.Keeper.GetLots(ctx, msg.LotId)
	if !found {
		return nil, types.ErrLotNotFound
	}
	if lot.Owner != msg.Creator {
		return nil, types.ErrLotAnotherOwner
	}
	if len(lot.Assets) == 0 {
		return nil, types.ErrEmptyAssetsList
	}
	if lot.Opened {
		return &types.MsgStartBiddingResponse{}, nil
	}
	if time.Unix(int64(msg.GetFinishTime()), 0).Sub(time.Now().UTC()) < minBiddingDuration {
		return nil, types.ErrWrongBiddingDuration
	}
	if msg.GetBid().IsNegative() {
		return nil, types.ErrWrongBidValue
	}

	lot.Opened = true
	lot.StartTime = uint64(time.Now().UTC().Unix())
	lot.FinishTime = msg.GetFinishTime()
	lot.Bid = msg.GetBid()
	lot.BidOwner = msg.GetCreator()
	k.Keeper.SetLots(ctx, lot)

	return &types.MsgStartBiddingResponse{}, nil
}
