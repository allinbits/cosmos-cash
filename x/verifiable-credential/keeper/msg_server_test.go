package keeper

import (
	"fmt"
	"time"

	didtypes "github.com/allinbits/cosmos-cash/v3/x/did/types"
	"github.com/allinbits/cosmos-cash/v3/x/verifiable-credential/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgSeverDeleteVerifableCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgRevokeCredential

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		//{
		//	"PASS: correctly deletes vc",
		//	func() {
		//		// NEED ACCOUNTS HERE
		//		vc := types.NewUserVerifiableCredential(
		//			"new-verifiable-cred-3",
		//			didDoc.Id,
		//			time.Now(),
		//			types.NewUserCredentialSubject(
		//				"accAddr",
		//				"root",
		//				true,
		//			),
		//		)
		//		suite.keeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
		//
		//		req = *types.NewMsgRevokeVerifiableCredential(vc.Id, "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a")
		//	},
		//	true,
		//},
		{
			"FAIL: vc issuer and did id do not match",
			func() {
				did := "did:cosmos:cash:subject"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewBlockchainAccountID(suite.ctx.ChainID(), "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a"),
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

				req = *types.NewMsgRevokeVerifiableCredential(vc.Id, "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a")
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
							didtypes.NewBlockchainAccountID(suite.ctx.ChainID(), "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a"),
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
				req = *types.NewMsgRevokeVerifiableCredential(
					"new-verifiable-cred-3",
					"did:cash:1111",
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			vcResp, err := server.RevokeCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(vcResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
