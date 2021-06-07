package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (k Keeper) SetIdentifier(ctx sdk.Context, key []byte, document types.DidDocument) {
	k.Set(ctx, key, types.IdentifierKey, document, k.MarshalIdentifier)
}

func (k Keeper) GetIdentifier(ctx sdk.Context, key []byte) (types.DidDocument, bool) {
	val, found := k.Get(ctx, key, types.IdentifierKey, k.UnmarshalIdentifier)
	return val.(types.DidDocument), found
}

func (k Keeper) UnmarshalIdentifier(value []byte) (interface{}, bool) {
	document := types.DidDocument{}

	err := k.cdc.UnmarshalBinaryBare(value, &document)
	if err != nil {
		return types.DidDocument{}, false
	}

	if document.Context == "" {
		return types.DidDocument{}, false
	}

	return document, true
}

func (k Keeper) MarshalIdentifier(value interface{}) []byte {
	identifier := value.(types.DidDocument)

	bytes, _ := k.cdc.MarshalBinaryBare(&identifier)

	return bytes
}

func (k Keeper) GetAllIdentifiersWithCondition(
	ctx sdk.Context,
	key []byte,
	identifierSelector func(identifiers types.DidDocument) bool,
) (identifiers []types.DidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		identifier, _ := k.UnmarshalIdentifier(iterator.Value())
		identifierTyped := identifier.(types.DidDocument)
		if identifierSelector(identifierTyped) {
			identifiers = append(identifiers, identifierTyped)
		}
	}

	return identifiers
}

func (k Keeper) GetAllIdentifiers(ctx sdk.Context) []types.DidDocument {
	return k.GetAllIdentifiersWithCondition(
		ctx,
		types.IdentifierKey,
		func(did types.DidDocument) bool { return true },
	)
}
