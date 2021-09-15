package keeper

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	vckeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential/keeper"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	server "github.com/cosmos/cosmos-sdk/server"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	dbm "github.com/tendermint/tm-db"
)

// Keeper test suit enables the keeper package to be tested
type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      Keeper
	queryClient types.QueryClient
	didkeeper   didkeeper.Keeper
	vckeeper    vckeeper.Keeper
}

// SetupTest creates a test suite to test the issuer
func (suite *KeeperTestSuite) SetupTest() {
	keyIssuer := sdk.NewKVStoreKey(types.StoreKey)
	memKeyIssuer := sdk.NewKVStoreKey(types.MemStoreKey)
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramtypes.StoreKey)
	memKeyParams := sdk.NewKVStoreKey(paramtypes.TStoreKey)
	keyVc := sdk.NewKVStoreKey(vctypes.StoreKey)
	memKeyVc := sdk.NewKVStoreKey(vctypes.MemStoreKey)
	keyDidDocument := sdk.NewKVStoreKey(didtypes.StoreKey)
	memKeyDidDocument := sdk.NewKVStoreKey(didtypes.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIssuer, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyIssuer, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyVc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyVc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyDidDocument, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyDidDocument, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, server.ZeroLogWrapper{log.Logger})

	interfaceRegistry := ct.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
		types.ModuleName:           {authtypes.Minter, authtypes.Burner},
	}

	allowedReceivingModAcc := map[string]bool{}

	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	paramsKeeper := paramskeeper.NewKeeper(marshaler, nil, keyParams, memKeyParams)

	AccountKeeper := authkeeper.NewAccountKeeper(
		marshaler,
		keyAcc,
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	BankKeeper := bankkeeper.NewBaseKeeper(
		marshaler,
		keyBank,
		AccountKeeper,
		paramsKeeper.Subspace(banktypes.ModuleName),
		blockedAddrs,
	)

	didKeeper := didkeeper.NewKeeper(
		marshaler,
		keyDidDocument,
		memKeyDidDocument,
	)

	vcKeeper := vckeeper.NewKeeper(
		marshaler,
		keyVc,
		memKeyVc,
		didKeeper,
	)

	k := NewKeeper(
		marshaler,
		keyIssuer,
		memKeyIssuer,
		BankKeeper,
		didKeeper,
		vcKeeper,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	types.RegisterQueryServer(queryHelper, k)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx, suite.keeper, suite.queryClient, suite.didkeeper, suite.vckeeper = ctx, *k, queryClient, *didKeeper, *vcKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestGenericKeeperSetAndGet() {
	testCases := []struct {
		msg     string
		issuer  types.Issuer
		expPass bool
	}{
		{
			"data stored successfully",
			types.Issuer{
				"context",
				10,
				"did:cash:1111",
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.keeper.Set(suite.ctx,
			[]byte(tc.issuer.IssuerDid),
			[]byte{0x01},
			tc.issuer,
			suite.keeper.MarshalIssuer,
		)
		suite.keeper.Set(suite.ctx,
			[]byte(tc.issuer.IssuerDid+"1"),
			[]byte{0x01},
			tc.issuer,
			suite.keeper.MarshalIssuer,
		)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.Get(
					suite.ctx,
					[]byte(tc.issuer.IssuerDid),
					[]byte{0x01},
					suite.keeper.UnmarshalIssuer,
				)
				suite.Require().True(found)

				iterator := suite.keeper.GetAll(
					suite.ctx,
					[]byte{0x01},
				)
				defer iterator.Close()

				var array []interface{}
				for ; iterator.Valid(); iterator.Next() {
					array = append(array, iterator.Value())
				}
				suite.Require().Equal(2, len(array))
			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}
