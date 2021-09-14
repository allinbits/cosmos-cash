package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/did/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document types.DidDocument) {
	k.Set(ctx, key, types.DidDocumentKey, document, k.Marshal)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (types.DidDocument, bool) {
	val, found := k.Get(ctx, key, types.DidDocumentKey, k.UnmarshalDidDocument)
	return val.(types.DidDocument), found
}

// UnmarshalDidDocument unmarshall a did document= and check if it is empty
// ad DID document is empty if contains no context
func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	data := types.DidDocument{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDIDDocument(&data)
}

func (k Keeper) SetDidMetadata(ctx sdk.Context, key []byte, meta types.DidMetadata) {
	k.Set(ctx, key, types.DidMetadataKey, meta, k.Marshal)
}

func (k Keeper) GetDidMetadata(ctx sdk.Context, key []byte) (types.DidMetadata, bool) {
	val, found := k.Get(ctx, key, types.DidMetadataKey, k.UnmarshalDidMetadata)
	return val.(types.DidMetadata), found
}

func (k Keeper) UnmarshalDidMetadata(value []byte) (interface{}, bool) {
	data := types.DidMetadata{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDIDMetadata(&data)
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.DidDocument:
		bytes = k.cdc.MustMarshal(&value)
	case types.DidMetadata:
		bytes = k.cdc.MustMarshal(&value)
	}
	return
}

// Unmarshal unmarshal a byte slice to a struct, return false in case of errors
func (k Keeper) Unmarshal(data []byte, val codec.ProtoMarshaler) bool {
	if len(data) == 0 {
		return false
	}
	if err := k.cdc.Unmarshal(data, val); err != nil {
		return false
	}
	return true
}

// GetAllDidDocumentsWithCondition retrieve a list of
// did document by some arbitrary criteria. The selector filter has access
// to both the did and it's metadata
func (k Keeper) GetAllDidDocumentsWithCondition(
	ctx sdk.Context,
	key []byte,
	didSelector func(did types.DidDocument) bool,
) (didDocs []types.DidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		did, _ := k.UnmarshalDidDocument(iterator.Value())
		didTyped := did.(types.DidDocument)

		if didSelector(didTyped) {
			didDocs = append(didDocs, didTyped)
		}
	}

	return didDocs
}

// GetAllDidDocuments returns all the DidDocuments
func (k Keeper) GetAllDidDocuments(ctx sdk.Context) []types.DidDocument {
	return k.GetAllDidDocumentsWithCondition(
		ctx,
		types.DidDocumentKey,
		func(did types.DidDocument) bool { return true },
	)
}

func (k Keeper) GetDidDocumentsByPubKey(ctx sdk.Context, pubkey cryptotypes.PubKey) []types.DidDocument {
	return k.GetAllDidDocumentsWithCondition(
		ctx,
		types.DidDocumentKey,
		func(did types.DidDocument) bool {
			return did.HasPublicKey(pubkey)
		},
	)
}
