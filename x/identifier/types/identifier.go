package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewIdentifier constructs a new Identifier
func NewIdentifier(operator sdk.AccAddress) (DidDocument, error) {
	return DidDocument{
		Context: operator.String(),
		Id:      operator.String(),
	}, nil
}

// GetBytes is a helper for serialising
func (did DidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
}
