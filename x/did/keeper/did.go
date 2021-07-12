package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/did/types"
)

func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document types.DidDocument) {
	k.Set(ctx, key, types.DidDocumentKey, document, k.MarshalDidDocument)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (types.DidDocument, bool) {
	val, found := k.Get(ctx, key, types.DidDocumentKey, k.UnmarshalDidDocument)
	return val.(types.DidDocument), found
}

func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	document := types.DidDocument{}

	err := k.cdc.UnmarshalBinaryBare(value, &document)
	if err != nil {
		return types.DidDocument{}, false
	}

	if len(document.Context) == 0 {
		return types.DidDocument{}, false
	}

	return document, true
}

func (k Keeper) MarshalDidDocument(value interface{}) []byte {
	did := value.(types.DidDocument)

	bytes, _ := k.cdc.MarshalBinaryBare(&did)

	return bytes
}

func (k Keeper) GetAllDidDocumentsWithCondition(
	ctx sdk.Context,
	key []byte,
	didSelector func(dids types.DidDocument) bool,
) (dids []types.DidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		did, _ := k.UnmarshalDidDocument(iterator.Value())
		didTyped := did.(types.DidDocument)
		if didSelector(didTyped) {
			dids = append(dids, didTyped)
		}
	}

	return dids
}

func (k Keeper) GetAllDidDocuments(ctx sdk.Context) []types.DidDocument {
	return k.GetAllDidDocumentsWithCondition(
		ctx,
		types.DidDocumentKey,
		func(did types.DidDocument) bool { return true },
	)
}
