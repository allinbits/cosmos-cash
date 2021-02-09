package identifier

import (
	"fmt"
	"github.com/stretchr/testify/suite"
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

// Keeper test suit enables the keeper package to be tested
type HandlerTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper keeper.Keeper
}

// SetupTest creates a test suite to test the identifier
func (suite *HandlerTestSuite) SetupTest() {
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

	suite.ctx, suite.keeper = ctx, *k
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) TestHandleMsgCreateIdentifier() {
	var (
		req types.MsgCreateIdentifier
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"can create a an identifier",
			func() { req = *types.NewMsgCreateIdentifier("did:cash:1111", nil, "did:cash:1111") },
			false,
		},
		{
			"identifier already exists",
			func() {
				auth := types.NewAuthentication(
					"did:cash:1111#keys-1",
					"sepk256",
					"did:cash:1111",
					"pubKey.Address().String()",
				)
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					types.Authentications{&auth},
					nil,
				}
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
				req = *types.NewMsgCreateIdentifier("did:cash:1111", nil, "did:cash:1111")
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", req), func() {
			tc.malleate()
			_, err := handleFn(suite.ctx, &req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *HandlerTestSuite) TestHandleMsgAddAuthentication() {
	var (
		req types.MsgAddAuthentication
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"can not add authentication, identifier does not exist",
			func() { req = *types.NewMsgAddAuthentication("did:cash:1111", nil, "did:cash:1111") },
			true,
		},
		{
			"can add authentication to did document",
			func() {
				auth := types.NewAuthentication(
					"did:cash:1111#keys-1",
					"sepk256",
					"did:cash:1111",
					"pubKey.Address().String()",
				)
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					types.Authentications{&auth},
					nil,
				}
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
				req = *types.NewMsgAddAuthentication("did:cash:1111", &auth, "did:cash:1111")
			},
			false,
		},
		// TODO: handle auth == nil case
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", req), func() {
			tc.malleate()
			_, err := handleFn(suite.ctx, &req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
