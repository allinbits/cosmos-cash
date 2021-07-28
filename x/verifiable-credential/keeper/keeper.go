package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// MarshalFn is a generic function to marshal interfaces
type MarshalFn func(value interface{}) []byte

// Keeper holds the application data
type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey
}

// NewKeeper create a new instance of a Keeper
func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (q Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (q Keeper) Set(ctx sdk.Context, key []byte, prefix []byte, i interface{}, marshal MarshalFn) {
	store := ctx.KVStore(q.storeKey)
	store.Set(append(prefix, key...), marshal(i))
}

// Get gets an item from the store by bytes
func (q Keeper) Get(ctx sdk.Context, key []byte, prefix []byte, unmarshal UnmarshalFn) (i interface{}, found bool) {
	store := ctx.KVStore(q.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// GetAll values from with a prefix from the store
func (q Keeper) GetAll(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(q.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}
