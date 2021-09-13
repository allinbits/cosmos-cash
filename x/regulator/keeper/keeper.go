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

// GetRegulatorsAddresses retrieve the genesis configured regulator addresses
func (k Keeper) GetRegulatorsAddresses(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.ParamStoreKeyRegulators)
	var r types.Regulators
	if err := k.cdc.Unmarshal(value, &r); err != nil {
		k.Logger(ctx).Error("error deserializing regulators:", err)
		return []string{}
	}
	return r.Addresses
}

// SetDidDocumentWithMeta commit a did document and the metadata to the persistent store. The DID of the
// document and the metadata is read from the did document
func (k Keeper) SetDidDocumentWithMeta(ctx sdk.Context, document didtypes.DidDocument, meta didtypes.DidMetadata) {
	k.didKeeper.SetDidDocumentWithMeta(ctx, document, meta)
}
