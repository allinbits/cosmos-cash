package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/v3/x/did/types"
)

// DidKeeper defines the expected did keeper functions
type DidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (types.DidDocument, bool)
	ResolveDid(ctx sdk.Context, did types.DID) (doc types.DidDocument, meta types.DidMetadata, err error)
}

// AccountKeeper defines the functions from the account keeper
type AccountKeeper interface {
	GetPubKey(ctx sdk.Context, addr sdk.AccAddress) (cryptotypes.PubKey, error)
}
