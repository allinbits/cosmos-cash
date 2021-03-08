package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	identifierkeeper "github.com/allinbits/cosmos-cash/x/identifier/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// UnmarshalFn is a generic function to unmarshal bytes
type MarshalFn func(value interface{}) []byte

type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	ik   identifierkeeper.Keeper
	vcsk vcskeeper.Keeper
	bk   bank.Keeper
}

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	ik identifierkeeper.Keeper,
	vcsk vcskeeper.Keeper,
	bk bank.Keeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		ik:       ik,
		vcsk:     vcsk,
		bk:       bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) Set(ctx sdk.Context,
	key []byte,
	prefix []byte,
	i interface{},
	marshal MarshalFn,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), marshal(i))
}

// Deletes a value form the sotre
func (k Keeper) Delete(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(prefix, key...))
}

// Get gets an item from the store by bytes
func (k Keeper) Get(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
	unmarshal UnmarshalFn,
) (i interface{}, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(
	ctx sdk.Context,
	prefix []byte,
) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}
