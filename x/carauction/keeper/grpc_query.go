package keeper

import (
	"github.com/nasa70/car-auction/x/carauction/types"
)

var _ types.QueryServer = Keeper{}
