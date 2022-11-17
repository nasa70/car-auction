package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/carauction module sentinel errors
var (
	ErrAssetNotFound      = sdkerrors.Register(ModuleName, 1100, "asset by id not found")
	ErrAssetAlreadyExists = sdkerrors.Register(ModuleName, 1101, "an asset with the same id already exists")
)
