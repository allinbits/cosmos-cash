package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identifier module sentinel errors
var (
	ErrIdentifierNotFound      = sdkerrors.Register(ModuleName, 1100, "did document not found")
	ErrIdentifierFound         = sdkerrors.Register(ModuleName, 1101, "did document found")
	ErrInvalidDIDFormat        = sdkerrors.Register(ModuleName, 1102, "is not compliant with the DID specifications (crf. https://www.w3.org/TR/did-core/#did-syntax)")
	ErrInvalidDIDURLFormat     = sdkerrors.Register(ModuleName, 1103, "is not compliant with the DID URL specifications (crf. https://www.w3.org/TR/did-core/#did-url-syntax)")
	ErrInvalidRFC3986UriFormat = sdkerrors.Register(ModuleName, 1104, "is not compliant with the RFC3986 URI specifications (crf. https://datatracker.ietf.org/doc/html/rfc3986)")
	ErrEmptyRelationships      = sdkerrors.Register(ModuleName, 1105, "a verification method should have at least one verification relationship. (cfr. https://www.w3.org/TR/did-core/#verification-relationships)")
	ErrUnauthorized            = sdkerrors.Register(ModuleName, 1106, "the transaction signer doesn't have the authorization to modify the did document")
	ErrInvalidState            = sdkerrors.Register(ModuleName, 1107, "the requested action is not applicable on the resource")
	ErrInvalidInput            = sdkerrors.Register(ModuleName, 1108, "input is invalid")
)
