package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/did/types"
)

func (suite *KeeperTestSuite) TestHandleMsgCreateDidDocument() {
	var (
		req types.MsgCreateDidDocument
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"Pass: can create a an did",
			func() { req = *types.NewMsgCreateDidDocument("did:cosmos:cash:subject", nil, nil, "subject") },
			false,
		},
		{
			"FAIL: did doc validation fails",
			func() { req = *types.NewMsgCreateDidDocument("invalid did", nil, nil, "subject") },
			true,
		},
		{
			"FAIL: did already exists",
			func() {
				did := "did:cosmos:cash:subject"
				didDoc, _ := types.NewDidDocument(did)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgCreateDidDocument(did, nil, nil, "subject")
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()
			_, err := server.CreateDidDocument(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgUpdateDidDocument() {
	var (
		req types.MsgUpdateDidDocument
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: not found",
			func() {
				req = *types.NewMsgUpdateDidDocument("did:cosmos:cash:subject", nil, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
		{
			"FAIL: unauthorized",
			func() {

				did := "did:cosmos:cash:subject"
				didDoc, _ := types.NewDidDocument(did)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgUpdateDidDocument(didDoc.Id, []string{"did:cosmos:cash:controller"}, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
		{
			"PASS: nil controllers",
			func() {

				did := "did:cosmos:cash:subject"
				didDoc, _ := types.NewDidDocument(did, types.WithVerifications(
					types.NewVerification(
						types.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{types.Authentication},
						nil,
					),
				))
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgUpdateDidDocument(did, nil, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			false,
		},
		{
			"FAIL: invalid controllers",
			func() {
				didDoc, _ := types.NewDidDocument("did:cosmos:cash:subject", types.WithVerifications(
					types.NewVerification(
						types.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{types.Authentication},
						nil,
					),
				))
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				controllers := []string{
					"did:cosmos:cash:controller-1",
					"did:cosmos:cash:controller-2",
					"invalid",
				}

				req = *types.NewMsgUpdateDidDocument(didDoc.Id, controllers, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()

			_, err := server.UpdateDidDocument(sdk.WrapSDKContext(suite.ctx), &req)

			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgAddVerification() {
	var (
		req types.MsgAddVerification
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not add verification, did does not exist",
			func() { req = *types.NewMsgAddVerification("did:cosmos:cash:subject", nil, "subject") },
			true,
		},
		{
			"FAIL: can not add verification, unauthorized",
			func() {
				// setup
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.CapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cosmos:cash:subject#key-2",
						"did:cosmos:cash:subject",
						types.NewBlockchainAccountID("foochainid", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
					),
					[]string{types.Authentication},
					nil,
				)
				req = *types.NewMsgAddVerification(didDoc.Id, v, "not a key")
			},
			true,
		},
		{
			"FAIL: can not add verification, unauthorized, key mismatch",
			func() {
				// setup
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.CapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cosmos:cash:subject#key-2",
						"did:cosmos:cash:subject",
						types.NewBlockchainAccountID("foochainid", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
					),
					[]string{types.Authentication},
					nil,
				)
				req = *types.NewMsgAddVerification(didDoc.Id, v, "cash1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2")
			},
			true,
		},
		{
			"FAIL: can not add verification, invalid verification",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := types.NewVerification(
					types.NewVerificationMethod(
						"",
						"did:cosmos:cash:subject",
						types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
					),
					[]string{types.Authentication},
					nil,
				)
				req = *types.NewMsgAddVerification(didDoc.Id, v, signer)
			},
			true,
		},
		{
			"PASS: can add verification to did document",
			func() {
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cosmos:cash:subject#key-2",
						"did:cosmos:cash:subject",
						types.NewBlockchainAccountID("foochainid", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
					),
					[]string{types.Authentication},
					nil,
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddVerification(didDoc.Id, v, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()

			_, err := server.AddVerification(sdk.WrapSDKContext(suite.ctx), &req)

			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgSetVerificationRelationships() {
	var (
		req types.MsgSetVerificationRelationships
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not add verification relationship, did does not exist",
			func() {
				req = *types.NewMsgSetVerificationRelationships(
					"did:cosmos:cash:subject",
					"did:cosmos:cash:subject#key-1",
					[]string{types.Authentication},
					"subject",
				)
			},
			true,
		},
		{
			"FAIL: can not add verification relationship, unauthorized",
			func() {
				// setup
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.CapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cosmos:cash:subject",
					"did:cosmos:cash:subject#key-1",
					[]string{types.Authentication},
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			true,
		},
		{
			"FAIL: can not add verification relationship, invalid relationships",
			func() {
				// setup
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cosmos:cash:subject",
					"did:cosmos:cash:subject#key-1",
					nil,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			true,
		},
		{
			"FAIL: verification method does not exist ",
			func() {
				// setup
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cosmos:cash:subject",
					"did:cosmos:cash:subject#key-does-not-exists",
					[]string{types.Authentication, types.CapabilityInvocation},
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			true,
		},
		{
			"PASS: add a new relationship",
			func() {
				// setup
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cosmos:cash:subject",
					"did:cosmos:cash:subject#key-1",
					[]string{types.Authentication, types.CapabilityInvocation},
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()

			_, err := server.SetVerificationRelationships(sdk.WrapSDKContext(suite.ctx), &req)

			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgRevokeVerification() {
	var (
		req types.MsgRevokeVerification
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not revoke verification, did does not exist",
			func() {
				req = *types.NewMsgRevokeVerification("did:cosmos:cash:2222", "service-id", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
		{
			"FAIL: can not revoke verification, not found",
			func() {
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgRevokeVerification(didDoc.Id, "did:cosmos:cash:subject#not-existent", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
		{
			"FAIL: can not revoke verification, unauthorized",
			func() {
				signer := "controller-1"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.CapabilityDelegation},
							nil,
						),
					),
				)

				vmID := "did:cosmos:cash:subject#key-1"

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// controller-1 does not exists
				req = *types.NewMsgRevokeVerification(didDoc.Id, vmID, signer)
			},
			true,
		},
		{
			"PASS: can revoke verification",
			func() {
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgRevokeVerification(didDoc.Id,
					"did:cosmos:cash:subject#key-1",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("TestHandleMsgRevokeVerification#%v", i), func() {
			tc.malleate()

			_, err := server.RevokeVerification(sdk.WrapSDKContext(suite.ctx), &req)

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

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not add service, did does not exist",
			func() {
				service := types.NewService(
					"service-id",
					"NonUserCredential",
					"cash/multihash",
				)
				req = *types.NewMsgAddService("did:cosmos:cash:subject", service, "subject")
			},
			true,
		},
		{
			"FAIL: can not add service, did does not exist",
			func() {
				req = *types.NewMsgAddService("did:cosmos:cash:subject", nil, "subject")
			},
			true,
		},
		{
			"FAIL: cannot add service to did document (unauthorized, wrong relationship)",
			func() {
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								types.DID("did:cosmos:cash:subject"),
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.CapabilityInvocation, types.CapabilityDelegation},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"UserCredential",
					"cash/multihash",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			true,
		},
		{
			"FAIL: cannot add service to did document with an incorrect type",
			func() {
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"NonUserCredential",
					"cash/multihash",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			true,
		},
		{
			"FAIL: duplicated service",
			func() {
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
					types.WithServices(
						types.NewService(
							"service-id",
							"UserCredential",
							"cash/multihash",
						),
					),
				)

				service := types.NewService(
					"service-id",
					"UserCredential",
					"cash/multihash",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			true,
		},
		{
			"PASS: can add service to did document",
			func() {
				signer := "subject"
				didDoc, err := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewBlockchainAccountID("foochainid", signer),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				service := types.NewService(
					"service-id",
					"UserCredential",
					"cash/multihash",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()

			_, err := server.AddService(sdk.WrapSDKContext(suite.ctx), &req)

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

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not delete service, did does not exist",
			func() {
				req = *types.NewMsgDeleteService("did:cosmos:cash:2222", "service-id", "cash1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
		{

			"Pass: can delete service from did document",
			func() {
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#cash1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
					types.WithServices(
						types.NewService(
							"service-id",
							"UserCredential",
							"cash/multihash",
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, "service-id", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			false,
		},
		{
			"FAIL: cannot remove an invalid serviceID",
			func() {

				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				serviceID := ""

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, "cash1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
		{
			"FAIL: unauthorized (wrong relationship)",
			func() {
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"did:cosmos:cash:subject",
								types.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
							),
							[]string{types.CapabilityInvocation},
							nil,
						),
					),
				)

				serviceID := "service-id"

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, "cash1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf(tc.name), func() {
			tc.malleate()

			_, err := server.DeleteService(sdk.WrapSDKContext(suite.ctx), &req)

			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}