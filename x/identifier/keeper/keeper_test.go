package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	//"github.com/allinbits/cosmos-cash/app"
	"github.com/allinbits/cosmos-cash/x/identifier/types"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	dbm "github.com/tendermint/tm-db"
)

// nolint:deadcode,unused,varcheck
var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())
	priv2 = secp256k1.GenPrivKey()
	addr2 = sdk.AccAddress(priv2.PubKey().Address())

	valKey  = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())

	PKs = simapp.CreateTestPubKeys(500)
)

// MakeTestCtxAndKeeper returns a cosmos sdk ctx and a test identifier keeper
func MakeTestCtxAndKeeper(t *testing.T) (sdk.Context, Keeper) {
	keyIdentifier := sdk.NewKVStoreKey(types.StoreKey)
	memKeyIdentifier := sdk.NewKVStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentifier, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyIdentifier, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, nil)

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	k := NewKeeper(
		marshaler,
		keyIdentifier,
		memKeyIdentifier,
	)

	return ctx, *k
}

func TestGenericKeeperSetAndGet(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// Create test entity
	entityId := "did:cash:1111"
	entity, _ := types.NewIdentifier(
		entityId,
		nil,
	)

	// Set a value in the store
	keeper.Set(ctx, []byte(entity.Id), []byte{0x01}, entity, keeper.MarshalIdentifier)

	_, found := keeper.Get(
		ctx,
		[]byte(entity.Id),
		[]byte{0x01},
		keeper.UnmarshalIdentifier,
	)

	// Check the store to see if the entity was saved
	require.True(t, found)
}

func TestGenericKeeperSetAndGetAll(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// Create test entity
	entityId := "did:cash:1111"
	entity, _ := types.NewIdentifier(
		entityId,
		nil,
	)

	// Set a value in the store
	keeper.Set(ctx, []byte(entity.Id), []byte{0x01}, entity, keeper.MarshalIdentifier)

	allEntities := keeper.GetAll(
		ctx,
		[]byte{0x01},
		keeper.UnmarshalIdentifier,
	)

	// Check the store to see if the entity was saved
	require.Equal(t, 1, len(allEntities))
}
