package keeper

import (
	"github.com/noria-net/noria/x/coinmaster/types"
)

var _ types.QueryServer = Keeper{}
