package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identifier module sentinel errors
var (
	ErrIdentifierNotFound = sdkerrors.Register(ModuleName, 1100, "identifier not found")
	ErrIdentifierFound    = sdkerrors.Register(ModuleName, 1101, "identifier found")
)
