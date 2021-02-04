package keeper

import (
	"github.com/allinbits/cosmos-cash/x/identifier/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	identiferSelector func(votes types.DidDocument) bool,
) (identifiers []types.DidDocument) {
	val := k.GetAll(ctx, key, k.UnmarshalIdentifier)

	for _, value := range val {
		identifer := value.(types.DidDocument)
		if identiferSelector(identifer) {
			identifiers = append(identifiers, value.(types.DidDocument))
		}
	}

	return identifiers
}

func (k Keeper) GetAllIdentifiers(ctx sdk.Context) []types.DidDocument {
	return k.GetAllIdentifiersWithCondition(
		ctx,
		types.IdentifierKey,
		func(votes types.DidDocument) bool { return true },
	)
}
