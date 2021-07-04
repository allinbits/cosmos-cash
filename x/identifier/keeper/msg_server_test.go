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
			func() {
				req = *types.NewMsgCreateIdentifier(
					"did:cash:subject",
					[]*types.Verification{},
					[]*types.Service{},
					"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
				)
			},
			true,
		},
		{
			"identifier already exists",
			func() {

				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					),
					[]string{
						types.RelationshipAuthentication,
					},
					nil,
				)

				didDoc, _ := types.NewIdentifier("did:cash:subject", types.WithVerifications(v))

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgCreateIdentifier(
					"did:cash:1111",
					[]*types.Verification{v},
					[]*types.Service{},
					"did:cash:1111",
				)
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

func (suite *KeeperTestSuite) TestMsgSeverAddVerification() {
	server := NewMsgServerImpl(suite.keeper)
	var (
		req types.MsgAddVerification
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		// {
		// 	"can not add authentication, identifier does not exist",
		// 	func() { req = *types.NewMsgAddVerification("did:cash:1111", nil, "did:cash:1111") },
		// 	false,
		// },
		// {
		// 	"can add authentication to did document",
		// 	func() {
		// 		auth := types.NewAuthentication(
		// 			"did:cash:1111#keys-1",
		// 			"sepk256",
		// 			"did:cash:1111",
		// 			"pubKey.Address().String()",
		// 		)
		// 		identifier := types.DidDocument{
		// 			"context",
		// 			"did:cash:1111",
		// 			types.Authentications{&auth},
		// 			nil,
		// 		}
		// 		suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)
		// 		req = *types.NewMsgAddAuthentication("did:cash:1111", &auth, "did:cash:1111")
		// 	},
		// 	true,
		// },
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			authResp, err := server.AddVerification(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(authResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSeverAddService() {
	server := NewMsgServerImpl(suite.keeper)
	var (
		req types.MsgAddService
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"FAIL: can not add service, identifier does not exist",
			func() { req = *types.NewMsgAddService("did:cash:1111", nil, "did:cash:1111") },
			false,
		},
		{
			"FAIL: cannot add a service to did document with an incorrect type",
			func() {
				s := types.NewService(
					"did:cash:1111",
					"NonKYCCredential",
					"did:cash:1111",
				)

				didDoc, _ := types.NewIdentifier("did:cash:subject", types.WithServices(s))

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgAddService("did:cash:subject", s, "subject")
			},
			false,
		},
		{
			"PASS: can add service to did document",
			func() {
				// service := types.NewService(
				// 	"did:cash:1111",
				// 	"IssuerCredential",
				// 	"did:cash:1111",
				// )
				// identifier := types.DidDocument{
				// 	"context",
				// 	"did:cash:1111",
				// 	nil,
				// 	nil,
				// }
				// suite.keeper.SetIdentifier(suite.ctx, []byte(identifier.Id), identifier)

				// req = *types.NewMsgAddService("did:cash:1111", &service, "cash:cash:1111")
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			authResp, err := server.AddService(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(authResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

//TODO: test delete auth
//TODO: test delete service
