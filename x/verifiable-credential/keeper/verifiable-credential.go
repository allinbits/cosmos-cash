package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

func (q Keeper) SetVerifiableCredential(ctx sdk.Context, key []byte, vc types.VerifiableCredential) {
	q.Set(ctx, key, types.VerifiableCredentialKey, vc, q.MarshalVerifiableCredential)
}

func (q Keeper) GetVerifiableCredential(ctx sdk.Context, key []byte) (types.VerifiableCredential, bool) {
	val, found := q.Get(ctx, key, types.VerifiableCredentialKey, q.UnmarshalVerifiableCredential)
	return val.(types.VerifiableCredential), found
}

func (q Keeper) DeleteVerifiableCredentialFromStore(ctx sdk.Context, key []byte) {
	q.Delete(ctx, key, types.VerifiableCredentialKey)
}

func (q Keeper) UnmarshalVerifiableCredential(value []byte) (interface{}, bool) {
	vc := types.VerifiableCredential{}

	err := q.cdc.Unmarshal(value, &vc)
	if err != nil {
		return types.VerifiableCredential{}, false
	}

	if vc.Id == "" {
		return types.VerifiableCredential{}, false
	}

	return vc, true
}

func (q Keeper) MarshalVerifiableCredential(value interface{}) []byte {
	did := value.(types.VerifiableCredential)

	bytes, _ := q.cdc.Marshal(&did)

	return bytes
}

func (q Keeper) GetAllVerifiableCredentialsWithCondition(
	ctx sdk.Context,
	key []byte,
	vcSelector func(votes types.VerifiableCredential) bool,
) (vcs []types.VerifiableCredential) {
	iterator := q.GetAll(ctx, key)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		vc, _ := q.UnmarshalVerifiableCredential(iterator.Value())
		vcTyped := vc.(types.VerifiableCredential)
		if vcSelector(vcTyped) {
			vcs = append(vcs, vcTyped)
		}
	}

	return vcs
}

func (q Keeper) GetAllVerifiableCredentials(ctx sdk.Context) []types.VerifiableCredential {
	return q.GetAllVerifiableCredentialsWithCondition(
		ctx,
		types.VerifiableCredentialKey,
		func(vc types.VerifiableCredential) bool { return true },
	)
}
