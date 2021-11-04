package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "github.com/allinbits/cosmos-cash/v2/x/did/types"
	vctypes "github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
)

// DidKeeper defines the expected did keeper functions
type DidKeeper interface {
	ResolveDid(ctx sdk.Context, did didtypes.DID) (doc didtypes.DidDocument, meta didtypes.DidMetadata, err error)
}

// VcKeeper defines the expected verfiable credentials keeper functions
type VcKeeper interface {
	GetVerifiableCredential(ctx sdk.Context, key []byte) (vctypes.VerifiableCredential, bool)
	SetVerifiableCredential(ctx sdk.Context, key []byte, vc vctypes.VerifiableCredential) error
	GetVerifiableCredentialWithType(ctx sdk.Context, subjectDID, vcType string) []vctypes.VerifiableCredential
}
