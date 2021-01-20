package keeper

import (
	"github.com/allinbits/cosmos-cash/x/issuer/types"
)

var _ types.QueryServer = Keeper{}
