package keeper

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"

	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	dbm "github.com/tendermint/tm-db"
)

// Keeper test suit enables the keeper package to be tested
type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *simapp.SimApp
	keeper      Keeper
	didkeeper   didkeeper.Keeper
	queryClient types.QueryClient
}

// SetupTest creates a test suite to test the did
func (suite *KeeperTestSuite) SetupTest() {
	keyVc := sdk.NewKVStoreKey(types.StoreKey)
	memKeyVc := sdk.NewKVStoreKey(types.MemStoreKey)
	keyDidDocument := sdk.NewKVStoreKey(didtypes.StoreKey)
	memKeyDidDocument := sdk.NewKVStoreKey(didtypes.MemStoreKey)
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramtypes.StoreKey)
	memKeyParams := sdk.NewKVStoreKey(paramtypes.TStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyVc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyVc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyDidDocument, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyDidDocument, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, nil)

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
	}

	paramsKeeper := paramskeeper.NewKeeper(marshaler, nil, keyParams, memKeyParams)

	didKeeper := didkeeper.NewKeeper(
		marshaler,
		keyDidDocument,
		memKeyDidDocument,
	)

	accountKeeper := authkeeper.NewAccountKeeper(
		marshaler,
		keyAcc,
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	k := NewKeeper(
		marshaler,
		keyVc,
		memKeyVc,
		didKeeper,
		accountKeeper,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	types.RegisterQueryServer(queryHelper, k)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx, suite.keeper, suite.didkeeper, suite.queryClient = ctx, *k, *didKeeper, queryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestGenericKeeperSetAndGet() {
	testCases := []struct {
		msg string
		did types.VerifiableCredential
		// TODO: add mallate func and clean up test
		expPass bool
	}{
		//{
		//	"data stored successfully",
		//	types.NewUserVerifiableCredential(
		//		"did:cash:1111",
		//		"",
		//		time.Now(),
		//		types.NewUserCredentialSubject("", "root", true),
		//	),
		//	true,
		//},
	}
	for _, tc := range testCases {
		suite.keeper.Set(suite.ctx,
			[]byte(tc.did.Id),
			[]byte{0x01},
			tc.did,
			suite.keeper.MarshalVerifiableCredential,
		)
		suite.keeper.Set(suite.ctx,
			[]byte(tc.did.Id+"1"),
			[]byte{0x01},
			tc.did,
			suite.keeper.MarshalVerifiableCredential,
		)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.Get(
					suite.ctx,
					[]byte(tc.did.Id),
					[]byte{0x01},
					suite.keeper.UnmarshalVerifiableCredential,
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
