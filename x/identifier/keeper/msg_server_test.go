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
					"did:cash:subject",
					[]*types.Verification{v},
					[]*types.Service{},
					"subject",
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
		{
			"can not add authentication, identifier does not exist",
			func() { req = *types.NewMsgAddVerification("did:cash:subject", nil, "subject") },
			false,
		},
		{
			"can add authentication to did document",
			func() {

				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(types.NewVerification(
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
					)),
				)

				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cash:controller-1#key-2",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:controller-1",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					),
					[]string{
						types.RelationshipAuthentication,
					},
					nil,
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddVerification("did:cash:subject", v, "subject")
			},
			true,
		},
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
			func() { req = *types.NewMsgAddService("did:cash:subject", nil, "subject") },
			false,
		},
		{
			"FAIL: cannot add a service to did document with an incorrect type",
			func() {

				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(types.NewVerification(
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
					)),
				)

				s := types.NewService(
					"did:cash:1111",
					"NonKYCCredential",
					"did:cash:1111",
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgAddService("did:cash:subject", s, "subject")
			},
			false,
		},
		{
			"PASS: can add service to did document",
			func() {

				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(types.NewVerification(
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
					)),
				)

				s := types.NewService(
					"service:seuro",
					"IssuerCredential",
					"service:seuro",
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgAddService("did:cash:subject", s, "subject")
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
