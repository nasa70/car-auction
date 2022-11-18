package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateLot = "create_lot"

var _ sdk.Msg = &MsgCreateLot{}

func NewMsgCreateLot(creator string, assets []uint64) *MsgCreateLot {
	return &MsgCreateLot{
		Creator: creator,
		Assets:  assets,
	}
}

func (msg *MsgCreateLot) Route() string {
	return RouterKey
}

func (msg *MsgCreateLot) Type() string {
	return TypeMsgCreateLot
}

func (msg *MsgCreateLot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
