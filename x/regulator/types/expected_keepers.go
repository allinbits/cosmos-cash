package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "github.com/allinbits/cosmos-cash/v3/x/did/types"
	vctypes "github.com/allinbits/cosmos-cash/v3/x/verifiable-credential/types"
)

// DidKeeper defines the expected did keeper functions
type DidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (didtypes.DidDocument, bool)
	SetDidDocument(ctx sdk.Context, key []byte, didDoc didtypes.DidDocument)
	SetDidMetadata(ctx sdk.Context, key []byte, didMeta didtypes.DidMetadata)
}

// VcKeeper defines the expected verifiable credentials keeper functions
type VcKeeper interface {
	SetVerifiableCredential(ctx sdk.Context, key []byte, vc vctypes.VerifiableCredential) error
	GetVerifiableCredentialWithType(ctx sdk.Context, subjectDID, vcType string) []vctypes.VerifiableCredential
}
