package keeper

import (
	"myblog/x/myblog/types"
)

var _ types.QueryServer = Keeper{}
