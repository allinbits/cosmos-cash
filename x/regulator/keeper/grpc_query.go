package keeper

import (
	"github.com/allinbits/cosmos-cash/v2/x/regulator/types"
)

var _ types.QueryServer = Keeper{}
