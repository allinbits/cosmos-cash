package keeper

import (
	"github.com/allinbits/cosmos-cash/x/cosmoscash/types"
)

var _ types.QueryServer = Keeper{}
