package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/carauction module sentinel errors
var (
	ErrAssetNotFound        = sdkerrors.Register(ModuleName, 1100, "asset by id not found")
	ErrAssetAlreadyExists   = sdkerrors.Register(ModuleName, 1101, "an asset with the same id already exists")
	ErrEmptyAssetsList      = sdkerrors.Register(ModuleName, 1102, "lot must have not empty list of assets")
	ErrAssetAnotherOwner    = sdkerrors.Register(ModuleName, 1103, "the asset belongs to another owner")
	ErrAssetAnotherLot      = sdkerrors.Register(ModuleName, 1104, "the asset belongs to another lot")
	ErrLotNotFound          = sdkerrors.Register(ModuleName, 1105, "lot by id not found")
	ErrLotAlreadyExists     = sdkerrors.Register(ModuleName, 1106, "an lot with the same id already exists")
	ErrLotAnotherOwner      = sdkerrors.Register(ModuleName, 1107, "the lot belongs to another owner")
	ErrWrongBiddingDuration = sdkerrors.Register(ModuleName, 1108, "min bidding duration is one hour")
	ErrWrongBidValue        = sdkerrors.Register(ModuleName, 1109, "bid value must be >= 0")
)
