package keeper

import (
	"github.com/allinbits/cosmos-cash/x/ibc-identifier/types"
)

var _ types.QueryServer = Keeper{}
