package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/carauction module sentinel errors
var (
	ErrAssetNotFound      = sdkerrors.Register(ModuleName, 1100, "asset by id not found")
	ErrAssetAlreadyExists = sdkerrors.Register(ModuleName, 1101, "an asset with the same id already exists")
	ErrEmptyAssetsList    = sdkerrors.Register(ModuleName, 1102, "lot must have not empty list of assets")
	ErrAssetAnotherOwner  = sdkerrors.Register(ModuleName, 1103, "the asset belongs to another owner")
	ErrAssetAnotherLot    = sdkerrors.Register(ModuleName, 1104, "the asset belongs to another lot")
)
