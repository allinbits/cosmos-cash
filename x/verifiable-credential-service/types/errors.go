package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/verifiable-credential-service module sentinel errors
var (
	ErrVerifiableCredentialNotFound = sdkerrors.Register(ModuleName, 1100, "vc not found")
	ErrVerifiableCredentialFound    = sdkerrors.Register(ModuleName, 1101, "vc found")
)
