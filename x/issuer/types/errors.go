package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/issuer module sentinel errors
var (
	ErrIssuerFound = sdkerrors.Register(ModuleName, 1101, "issuer found")
)
