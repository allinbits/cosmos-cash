package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
)

func (k Keeper) SetIssuer(ctx sdk.Context, issuer types.Issuer) {
	k.Set(ctx, []byte(issuer.IssuerDid), types.IssuerKey, issuer, k.MarshalIssuer)
	k.Set(ctx, []byte(issuer.Token), types.TokenKey, issuer, k.MarshalIssuer)
}

func (k Keeper) GetIssuer(ctx sdk.Context, key []byte) (types.Issuer, bool) {
	val, found := k.Get(ctx, key, types.IssuerKey, k.UnmarshalIssuer)
	return val.(types.Issuer), found
}

// GetIssuerByToken retrieve an issuer by it's key
func (k Keeper) GetIssuerByToken(ctx sdk.Context, key []byte) (types.Issuer, bool) {
	val, found := k.Get(ctx, key, types.TokenKey, k.UnmarshalIssuer)
	return val.(types.Issuer), found
}

func (k Keeper) UnmarshalIssuer(value []byte) (interface{}, bool) {
	issuer := types.Issuer{}

	err := k.cdc.Unmarshal(value, &issuer)
	if err != nil {
		return types.Issuer{}, false
	}

	if !didtypes.IsValidDID(issuer.IssuerDid) {
		return types.Issuer{}, false
	}

	return issuer, true
}

func (k Keeper) MarshalIssuer(value interface{}) []byte {
	issuer := value.(types.Issuer)

	bytes, _ := k.cdc.Marshal(&issuer)

	return bytes
}

func (k Keeper) GetAllIssuersWithCondition(
	ctx sdk.Context,
	key []byte,
	issuerSelector func(issuers types.Issuer) bool,
) (issuers []types.Issuer) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		issuer, _ := k.UnmarshalIssuer(iterator.Value())
		issuerTyped := issuer.(types.Issuer)
		if issuerSelector(issuerTyped) {
			issuers = append(issuers, issuerTyped)
		}
	}

	return issuers
}

func (k Keeper) GetAllIssuers(ctx sdk.Context) []types.Issuer {
	return k.GetAllIssuersWithCondition(
		ctx,
		types.IssuerKey,
		func(did types.Issuer) bool { return true },
	)
}
