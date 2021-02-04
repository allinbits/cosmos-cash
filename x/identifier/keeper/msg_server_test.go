package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgSeverCreateIdentifiers() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgCreateIdentifier

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"correctly creates identifier",
			func() { req = *types.NewMsgCreateIdentifier("did:cash:1111", nil, "did:cash:1111") },
			true,
		},
		{
			"identifier already exists",
			func() {
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					nil,
					nil,
				}
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)

				req = *types.NewMsgCreateIdentifier("did:cash:1111", nil, "did:cash:1111")
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			identifierResp, err := server.CreateIdentifier(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(identifierResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSeverAddAuthentication() {
	server := NewMsgServerImpl(suite.keeper)
	var (
		req types.MsgAddAuthentication
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"can not add authentication, identifier does not exist",
			func() { req = *types.NewMsgAddAuthentication("did:cash:1111", nil, "did:cash:1111") },
			false,
		},
		{
			"can add authentication to did document",
			func() {
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					nil,
					nil,
				}
				auth := types.NewAuthentication(
					"",
					"sepk256",
					"address.String()",
					"pubKey.Address().String()",
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
				req = *types.NewMsgAddAuthentication("did:cash:1111", &auth, "did:cash:1111")
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			authResp, err := server.AddAuthentication(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(authResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
