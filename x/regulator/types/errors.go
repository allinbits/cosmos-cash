package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/regulator module sentinel errors
var (
	ErrNotARegulator = sdkerrors.Register(ModuleName, 1100, "address is not registered as regulator")
	ErrAlreadyActive = sdkerrors.Register(ModuleName, 1101, "the regulator credential has already been issued for this address")
	// this line is used by starport scaffolding # ibc/errors
)
