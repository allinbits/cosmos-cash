package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
			"Pass: can create a an identifier",
			func() { req = *types.NewMsgCreateIdentifier("did:cash:1111", nil, "did:cash:1111") },
			false,
		},
		{
			"Fail: identifier already exists",
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
			"Fail: can not add authentication, identifier does not exist",
			func() { req = *types.NewMsgAddAuthentication("did:cash:1111", nil, "did:cash:1111") },
			true,
		},
		{
			"Pass: can add authentication to did document",
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

func (suite *KeeperTestSuite) TestHandleMsgDeleteAuthentication() {
	var (
		req types.MsgDeleteAuthentication
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"Fail: can not delete service, identifier does not exist",
			func() { req = *types.NewMsgDeleteAuthentication("did:cash:2222", "service-id", "did:cash:2222") },
			true,
		},
		{
			"Fail: can not delete service, no services on identifier",
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
				req = *types.NewMsgDeleteAuthentication("did:cash:1111", "service-id", "did:cash:1111")
			},
			true,
		},
		{

			"Pass: can delete service from did document",
			func() {
				pubKeyBech32 := "cosmospub1addwnpepq2x59cyqkp59vh2ghxqwnzx9xnceas49x78nscfmymsaqahkjqx4k2jc54x"
				pubKey, _ := sdk.GetPubKeyFromBech32(
					sdk.Bech32PubKeyTypeAccPub,
					pubKeyBech32,
				)

				auth := types.NewAuthentication(
					"did:cash:1111#keys-1",
					"sepk256",
					pubKey.Address().String(),
					pubKey.Address().String(),
				)
				auth2 := types.NewAuthentication(
					"did:cash:1111#keys-1",
					"sepk256",
					pubKey.Address().String(),
					pubKey.Address().String(),
				)
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					types.Authentications{&auth, &auth2},
					nil,
				}
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
				req = *types.NewMsgDeleteAuthentication(
					"did:cash:1111",
					pubKeyBech32,
					pubKey.Address().String(),
				)
			},
			false,
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
			"Fail: can not add service, identifier does not exist",
			func() {
				req = *types.NewMsgAddService("did:cash:1111", nil, "did:cash:1111")
			},
			true,
		},
		{
			"Pass: can add service to did document",
			func() {
				auth := types.NewAuthentication(
					"did:cash:1111#keys-1",
					"sepk256",
					"did:cash:1111",
					"pubKey.Address().String()",
				)
				service := types.NewService(
					"service-id",
					"VerifiableCredentials",
					"cash/multihash",
				)
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					types.Authentications{&auth},
					nil,
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

func (suite *KeeperTestSuite) TestHandleMsgDeleteService() {
	var (
		req types.MsgDeleteService
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"Fail: can not delete service, identifier does not exist",
			func() { req = *types.NewMsgDeleteService("did:cash:2222", "service-id", "did:cash:2222") },
			true,
		},
		{

			"Pass: can delete service from did document",
			func() {
				auth := types.NewAuthentication(
					"did:cash:1111#keys-1",
					"sepk256",
					"did:cash:1111",
					"pubKey.Address().String()",
				)
				service := types.NewService(
					"service-id",
					"VerifiableCredentials",
					"cash/multihash",
				)
				identifier := types.DidDocument{
					"context",
					"did:cash:1111",
					types.Authentications{&auth},
					types.Services{&service},
				}
				suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
				req = *types.NewMsgDeleteService("did:cash:1111", "service-id", "did:cash:1111")
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
