package keeper

import (
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetIssuer(ctx sdk.Context, key []byte, document types.Issuer) {
	k.Set(ctx, key, types.IssuerKey, document, k.MarshalIssuer)
}

func (k Keeper) GetIssuer(ctx sdk.Context, key []byte) (types.Issuer, bool) {
	val, found := k.Get(ctx, key, types.IssuerKey, k.UnmarshalIssuer)
	return val.(types.Issuer), found
}

func (k Keeper) UnmarshalIssuer(value []byte) (interface{}, bool) {
	issuer := types.Issuer{}

	err := k.cdc.UnmarshalBinaryBare(value, &issuer)
	if err != nil {
		return types.Issuer{}, false
	}

	if issuer.Address == "" {
		return types.Issuer{}, false
	}

	return issuer, true
}

func (k Keeper) MarshalIssuer(value interface{}) []byte {
	identifier := value.(types.Issuer)

	bytes, _ := k.cdc.MarshalBinaryBare(&identifier)

	return bytes
}

func (k Keeper) GetAllIssuersWithCondition(ctx sdk.Context, key []byte, identiferSelector func(votes types.Issuer) bool) (identifiers []types.Issuer) {
	val := k.GetAll(ctx, key, k.UnmarshalIssuer)

	for _, value := range val {
		identifer := value.(types.Issuer)
		if identiferSelector(identifer) {
			identifiers = append(identifiers, value.(types.Issuer))
		}
	}

	return identifiers
}

func (k Keeper) GetAllIssuers(ctx sdk.Context) []types.Issuer {
	return k.GetAllIssuersWithCondition(ctx, types.IssuerKey, func(votes types.Issuer) bool { return true })
}
