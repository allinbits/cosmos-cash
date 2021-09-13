package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// DidKeeper defines the expected did keeper functions
type DidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (didtypes.DidDocument, bool)
	SetDidDocumentWithMeta(ctx sdk.Context, didDoc didtypes.DidDocument, didMeta didtypes.DidMetadata)
}

// VcKeeper defines the expected verfiable credentials keeper functions
type VcKeeper interface {
	GetVerifiableCredential(ctx sdk.Context, key []byte) (vctypes.VerifiableCredential, bool)
}
