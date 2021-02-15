package keeper

import (
	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetVerifiableCredential(ctx sdk.Context, key []byte, vc types.VerifiableCredential) {
	k.Set(ctx, key, types.VerifiableCredentialKey, vc, k.MarshalVerifiableCredential)
}

func (k Keeper) GetVerifiableCredential(ctx sdk.Context, key []byte) (types.VerifiableCredential, bool) {
	val, found := k.Get(ctx, key, types.VerifiableCredentialKey, k.UnmarshalVerifiableCredential)
	return val.(types.VerifiableCredential), found
}

func (k Keeper) UnmarshalVerifiableCredential(value []byte) (interface{}, bool) {
	vc := types.VerifiableCredential{}

	err := k.cdc.UnmarshalBinaryBare(value, &vc)
	if err != nil {
		return types.VerifiableCredential{}, false
	}

	if vc.Id == "" {
		return types.VerifiableCredential{}, false
	}

	return vc, true
}

func (k Keeper) MarshalVerifiableCredential(value interface{}) []byte {
	identifier := value.(types.VerifiableCredential)

	bytes, _ := k.cdc.MarshalBinaryBare(&identifier)

	return bytes
}

func (k Keeper) GetAllVerifiableCredentialsWithCondition(
	ctx sdk.Context,
	key []byte,
	vcSelector func(votes types.VerifiableCredential) bool,
) (vcs []types.VerifiableCredential) {
	iterator := k.GetAll(ctx, key)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		vc, _ := k.UnmarshalVerifiableCredential(iterator.Value())
		vcTyped := vc.(types.VerifiableCredential)
		if vcSelector(vcTyped) {
			vcs = append(vcs, vcTyped)
		}
	}

	return vcs
}

func (k Keeper) GetAllVerifiableCredentials(ctx sdk.Context) []types.VerifiableCredential {
	return k.GetAllVerifiableCredentialsWithCondition(
		ctx,
		types.VerifiableCredentialKey,
		func(vc types.VerifiableCredential) bool { return true },
	)
}
