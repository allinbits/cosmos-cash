package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # ibc/keeper/import
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/tendermint/tendermint/libs/log"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
		bk        bank.Keeper
		didKeeper types.DidKeeper
		vcKeeper  types.VcKeeper
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	bk bank.Keeper,
	didKeeper types.DidKeeper,
	vcKeeper types.VcKeeper,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		bk:        bk,
		didKeeper: didKeeper,
		vcKeeper:  vcKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetRegulators sets a protocol buffer object in the db with a prefixed key
func (k Keeper) SetRegulators(ctx sdk.Context,
	params *types.Regulators,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ParamStoreKeyRegulators, k.cdc.MustMarshal(params))
}

// GetRegulators sets a protocol buffer object in the db with a prefixed key
func (k Keeper) GetRegulators(ctx sdk.Context) *types.Regulators {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.ParamStoreKeyRegulators)
	var r types.Regulators
	if err := k.cdc.Unmarshal(value, &r); err != nil {
		k.Logger(ctx).Error("error deserializing regulators:", err)
		return nil
	}
	return &r
}

// GetRegulatorsAddresses retrieve the genesis configured regulator addresses and dids
func (k Keeper) GetRegulatorsAddresses(ctx sdk.Context) []string {
	r := k.GetRegulators(ctx)
	if r == nil {
		return []string{}
	}
	return r.Addresses
}

// SetDidDocumentWithMeta commit a did document and the metadata to the persistent store. The DID of the
// document and the metadata is read from the did document
func (k Keeper) SetDidDocumentWithMeta(ctx sdk.Context, document didtypes.DidDocument, meta didtypes.DidMetadata) {
	key := []byte(document.Id)
	k.didKeeper.SetDidDocument(ctx, key, document)
	k.didKeeper.SetDidMetadata(ctx, key, meta)
}

// SetVerifiableCredential store verifiable credentials
func (k Keeper) SetVerifiableCredential(ctx sdk.Context, vc vctypes.VerifiableCredential) error {
	return k.vcKeeper.SetVerifiableCredential(ctx, []byte(vc.Id), vc)
}

// GetVerifiableCredential return a verifiable credential if exists
func (k Keeper) GetVerifiableCredential(ctx sdk.Context, vcID string) (vctypes.VerifiableCredential, bool) {
	return k.vcKeeper.GetVerifiableCredential(ctx, []byte(vcID))
}

// DeleteVerifiableCredential invokes the verifiable credential keeper to execute the deletion of credentials
func (k Keeper) DeleteVerifiableCredential(ctx sdk.Context, credentialID string) error {
	return k.vcKeeper.DeleteVerifiableCredentialFromStore(ctx, []byte(credentialID))
}
