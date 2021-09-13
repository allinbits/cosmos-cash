package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/allinbits/cosmos-cash/x/regulator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # ibc/keeper/import
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
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
	vcKeepr types.VcKeeper,

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

// SetRegulatorsParams sets a protocol buffer object in the db with a prefixed key
func (k Keeper) SetRegulatorsParams(ctx sdk.Context,
	params *types.Regulators,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ParamStoreKeyRegulatorsParams, k.cdc.MustMarshal(params))
}
