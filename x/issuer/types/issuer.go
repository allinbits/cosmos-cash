package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewIdentifier constructs a new Identifier
func NewIssuer(name string, token string, fee int32, state string, address string) (Issuer, error) {
	return Issuer{
		Name:    name,
		Token:   token,
		Fee:     fee,
		State:   state,
		Address: address,
	}, nil
}

// GetBytes is a helper for serialising
func (issuer Issuer) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&issuer))
}
