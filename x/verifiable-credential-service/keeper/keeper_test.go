package keeper

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tm-db"
)

// Keeper test suit enables the keeper package to be tested
type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      Keeper
	queryClient types.QueryClient
}

// SetupTest creates a test suite to test the identifier
func (suite *KeeperTestSuite) SetupTest() {
	keyVcs := sdk.NewKVStoreKey(types.StoreKey)
	memKeyVcs := sdk.NewKVStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyVcs, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyVcs, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, nil)

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	k := NewKeeper(
		marshaler,
		keyVcs,
		memKeyVcs,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	types.RegisterQueryServer(queryHelper, k)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx, suite.keeper, suite.queryClient = ctx, *k, queryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestGenericKeeperSetAndGet() {
	testCases := []struct {
		msg        string
		identifier types.VerifiableCredential
		// TODO: add mallate func and clean up test
		expPass bool
	}{
		{
			"data stored successfully",
			types.NewUserVerifiableCredential(
				"did:cash:1111",
				[]string{"context"},
				"",
				"",
				types.NewUserCredentialSubject("", "root", true),
				types.NewProof("", "", "", "", ""),
			),
			true,
		},
	}
	for _, tc := range testCases {
		suite.keeper.Set(suite.ctx,
			[]byte(tc.identifier.Id),
			[]byte{0x01},
			tc.identifier,
			suite.keeper.MarshalVerifiableCredential,
		)
		suite.keeper.Set(suite.ctx,
			[]byte(tc.identifier.Id+"1"),
			[]byte{0x01},
			tc.identifier,
			suite.keeper.MarshalVerifiableCredential,
		)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.Get(
					suite.ctx,
					[]byte(tc.identifier.Id),
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
