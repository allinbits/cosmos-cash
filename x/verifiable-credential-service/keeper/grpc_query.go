package keeper

import (
	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
)

var _ types.QueryServer = Keeper{}
