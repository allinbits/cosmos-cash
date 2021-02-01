package identifier

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/allinbits/cosmos-cash/x/identifier/keeper"
	"github.com/allinbits/cosmos-cash/x/identifier/types"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	dbm "github.com/tendermint/tm-db"
)

func bootstrapHandler(t *testing.T) (sdk.Context, keeper.Keeper) {
	// TODO: can we use the KeeperTestSuite here

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

	k := keeper.NewKeeper(
		marshaler,
		keyIdentifier,
		memKeyIdentifier,
	)
	return ctx, *k
}

func TestHandleMsgCreateIdentifier(t *testing.T) {
	// TODO: add failure cases
	ctx, k := bootstrapHandler(t)

	handleFn := NewHandler(k)

	testCases := []struct {
		name string
		msg  sdk.Msg
	}{
		{
			"can create a an identifier",
			types.NewMsgCreateIdentifier("", nil, ""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(*testing.T) {
			_, err := handleFn(ctx, tc.msg)
			assert.NoError(t, err)
		})
	}
}
