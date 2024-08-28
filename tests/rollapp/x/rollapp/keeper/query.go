package keeper

import (
	"rollapp/x/rollapp/types"
)

var _ types.QueryServer = Keeper{}
