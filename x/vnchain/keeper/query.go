package keeper

import (
	"vnchain/x/vnchain/types"
)

var _ types.QueryServer = Keeper{}
