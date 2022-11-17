package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAsset = "create_asset"

var _ sdk.Msg = &MsgCreateAsset{}

func NewMsgCreateAsset(creator string, description string) *MsgCreateAsset {
	return &MsgCreateAsset{
		Creator:     creator,
		Description: description,
	}
}

func (msg *MsgCreateAsset) Route() string {
	return RouterKey
}

func (msg *MsgCreateAsset) Type() string {
	return TypeMsgCreateAsset
}

func (msg *MsgCreateAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
