package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (suite *KeeperTestSuite) TestHandleMsgCreateIdentifier() {
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

func (suite *KeeperTestSuite) TestHandleMsgAddAuthentication() {
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

func (suite *KeeperTestSuite) TestHandleMsgAddService() {
	var (
		req types.MsgAddService
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"can not add authentication, identifier does not exist",
			func() { req = *types.NewMsgAddService("did:cash:1111", nil, "did:cash:1111") },
			true,
		},
		{
			"can add authentication to did document",
			func() {
				service := types.NewService(
					"service-id",
					"VerifiableCredentials",
					"cash/multihash",
				)
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					nil,
					types.Services{&service},
				}
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
				req = *types.NewMsgAddService("did:cash:1111", &service, "did:cash:1111")
			},
			false,
		},
		// TODO: handle service == nil case
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
