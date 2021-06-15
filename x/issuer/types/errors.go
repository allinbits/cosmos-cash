package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/issuer module sentinel errors
var (
	ErrIssuerFound        = sdkerrors.Register(ModuleName, 1101, "issuer found")
	ErrUserFound          = sdkerrors.Register(ModuleName, 1102, "user found")
	ErrInvalidIssuerDenom = sdkerrors.Register(ModuleName, 1103, "invalid denom for issuer coin")
)
