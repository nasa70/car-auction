package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartBidding = "start_bidding"

var _ sdk.Msg = &MsgStartBidding{}

func NewMsgStartBidding(creator string, lotId uint64, finishTime uint64, bid sdk.Coin) *MsgStartBidding {
	return &MsgStartBidding{
		Creator:    creator,
		LotId:      lotId,
		FinishTime: finishTime,
		Bid:        bid,
	}
}

func (msg *MsgStartBidding) Route() string {
	return RouterKey
}

func (msg *MsgStartBidding) Type() string {
	return TypeMsgStartBidding
}

func (msg *MsgStartBidding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartBidding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartBidding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
