package keeper

import (
	"github.com/allinbits/cosmos-cash/x/regulator/types"
)

var _ types.QueryServer = Keeper{}
