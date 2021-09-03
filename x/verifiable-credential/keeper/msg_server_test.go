package keeper

import (
	"fmt"
	"time"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgSeverCreateVerifableCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgCreateVerifiableCredential

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"correctly creates vc",
			func() {
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					"accAddr",
					time.Now(),
					cs,
				)
				req = *types.NewMsgCreateVerifiableCredential(vc, "did:cash:1111")
			},
			true,
		},
		{
			"vc already exists",
			func() {
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					"accAddr",
					time.Now(),
					cs,
				)
				suite.keeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)

				req = *types.NewMsgCreateVerifiableCredential(vc, "did:cash:1111")
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			vcResp, err := server.CreateVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(vcResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSeverDeleteVerifableCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgDeleteVerifiableCredential

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: correctly deletes vc",
			func() {
				did := "did:cosmos:cash:subject"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.keeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgDeleteVerifiableCredential(vc.Id, vc.Issuer, "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a")
			},
			true,
		},
		{
			"FAIL: vc issuer and did id do not match",
			func() {
				did := "did:cosmos:cash:subject"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					"did:cosmos:cash:noone",
					time.Now(),
					cs,
				)
				suite.keeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgDeleteVerifiableCredential(vc.Id, vc.Issuer, "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a")
			},
			false,
		},
		{
			"FAIL: vc does not exist",
			func() {
				did := "did:cosmos:cash:subject"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
			},
			false,
		},
		{
			"FAIL: did does not exists",
			func() {
				req = *types.NewMsgDeleteVerifiableCredential(
					"new-verifiable-cred-3",
					"did:cosmos:cash:issuer",
					"did:cash:1111",
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			vcResp, err := server.DeleteVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(vcResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
