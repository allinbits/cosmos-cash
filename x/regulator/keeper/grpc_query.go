package keeper

import (
	"github.com/allinbits/cosmos-cash/v3/x/regulator/types"
)

var _ types.QueryServer = Keeper{}
