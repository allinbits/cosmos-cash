package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identifier module sentinel errors
var (
	ErrIdentifierNotFound      = sdkerrors.Register(ModuleName, 1100, "identifier not found")
	ErrIdentifierFound         = sdkerrors.Register(ModuleName, 1101, "identifier found")
	ErrInvalidDIDFormat        = sdkerrors.Register(ModuleName, 1102, "input is not compliant with the DID specifications (crf. https://www.w3.org/TR/did-core/#did-syntax)")
	ErrInvalidDIDURLFormat     = sdkerrors.Register(ModuleName, 1103, "input is not compliant with the DID URL specifications (crf. https://www.w3.org/TR/did-core/#did-url-syntax)")
	ErrInvalidRFC3986UriFormat = sdkerrors.Register(ModuleName, 1104, "input is not compliant with the RFC3986 URI specifications (crf. https://datatracker.ietf.org/doc/html/rfc3986)")
)
