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
			"Pass: can create a an identifier",
			func() { req = *types.NewMsgCreateIdentifier("did:cash:subject", nil, nil, "subject") },
			false,
		},
		{
			"FAIL: did doc validation fails",
			func() { req = *types.NewMsgCreateIdentifier("invalid did", nil, nil, "subject") },
			true,
		},
		{
			"FAIL: identifier already exists",
			func() {
				did := "did:cash:subject"
				didDoc, _ := types.NewIdentifier(did)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgCreateIdentifier(did, nil, nil, "subject")
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

func (suite *KeeperTestSuite) TestHandleMsgUpdateIdentifier() {
	var (
		req types.MsgUpdateIdentifier
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: not found",
			func() {
				req = *types.NewMsgUpdateIdentifier("did:cash:subject", nil, "subject")
			},
			true,
		},
		{
			"FAIL: unauthorized",
			func() {

				signer := "subject"
				did := "did:cash:subject"
				didDoc, _ := types.NewIdentifier(did)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgUpdateIdentifier(didDoc.Id, []string{"did:cash:controller"}, signer)
			},
			true,
		},
		{
			"PASS: nil controllers",
			func() {
				signer := "subject"
				did := "did:cash:subject"
				didDoc, _ := types.NewIdentifier(did, types.WithVerifications(
					types.NewVerification(
						types.NewVerificationMethod(
							"did:cash:subject#key-1",
							"EcdsaSecp256k1RecoveryMethod2020",
							"did:cash:subject",
							"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
						),
						[]string{types.RelationshipAuthentication},
						nil,
					),
				))
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgUpdateIdentifier(did, nil, signer)
			},
			false,
		},
		{
			"FAIL: invalid controllers",
			func() {
				signer := "subject"
				did := "did:cash:subject"
				didDoc, _ := types.NewIdentifier(did, types.WithVerifications(
					types.NewVerification(
						types.NewVerificationMethod(
							"did:cash:subject#key-1",
							"EcdsaSecp256k1RecoveryMethod2020",
							"did:cash:subject",
							"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
						),
						[]string{types.RelationshipAuthentication},
						nil,
					),
				))
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)

				controllers := []string{
					"did:cash:controller-1",
					"did:cash:controller-2",
					"invalid",
				}

				req = *types.NewMsgUpdateIdentifier(did, controllers, signer)
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

func (suite *KeeperTestSuite) TestHandleMsgAddVerification() {
	var (
		req types.MsgAddVerification
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not add verification, identifier does not exist",
			func() { req = *types.NewMsgAddVerification("did:cash:subject", nil, "subject") },
			true,
		},
		{
			"FAIL: can not add verification, unauthorized",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipCapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cash:subject#key-2",
						"EcdsaSecp256k1RecoveryMethod2020",
						"did:cash:subject",
						"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
					),
					[]string{types.RelationshipAuthentication},
					nil,
				)
				req = *types.NewMsgAddVerification(didDoc.Id, v, signer)
			},
			true,
		},
		{
			"FAIL: can not add verification, invalid verification",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := types.NewVerification(
					types.NewVerificationMethod(
						"",
						"EcdsaSecp256k1RecoveryMethod2020",
						"did:cash:subject",
						"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
					),
					[]string{types.RelationshipAuthentication},
					nil,
				)
				req = *types.NewMsgAddVerification(didDoc.Id, v, signer)
			},
			true,
		},
		{
			"Pass: can add verification to did document",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)

				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cash:subject#key-2",
						"EcdsaSecp256k1RecoveryMethod2020",
						"did:cash:subject",
						"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
					),
					[]string{types.RelationshipAuthentication},
					nil,
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddVerification(didDoc.Id, v, signer)
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

func (suite *KeeperTestSuite) TestHandleMsgSetVerificationRelationships() {
	var (
		req types.MsgSetVerificationRelationships
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not add verification relationship, identifier does not exist",
			func() {
				req = *types.NewMsgSetVerificationRelationships(
					"did:cash:subject",
					"did:cash:subject#key-1",
					[]string{types.RelationshipAuthentication},
					"subject",
				)
			},
			true,
		},
		{
			"FAIL: can not add verification relationship, unauthorized",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipCapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cash:subject",
					"did:cash:subject#key-1",
					[]string{types.RelationshipAuthentication},
					signer,
				)
			},
			true,
		},
		{
			"FAIL: can not add verification relationship, invalid relationships",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cash:subject",
					"did:cash:subject#key-1",
					nil,
					signer,
				)
			},
			true,
		},
		{
			"PASS: add a new relationship",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *types.NewMsgSetVerificationRelationships(
					"did:cash:subject",
					"did:cash:subject#key-1",
					[]string{types.RelationshipCapabilityInvocation},
					signer,
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
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

func (suite *KeeperTestSuite) TestHandleMsgRevokeVerification() {
	var (
		req types.MsgRevokeVerification
	)

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: can not revoke verification, identifier does not exist",
			func() { req = *types.NewMsgRevokeVerification("did:cash:2222", "service-id", "did:cash:2222") },
			true,
		},
		{
			"FAIL: can not revoke verification, not found",
			func() {
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)
				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgRevokeVerification(didDoc.Id, "did:cash:subject#not-existent", "subject")
			},
			true,
		},
		{
			"FAIL: can not revoke verification, unauthorized",
			func() {
				signer := "controller-1"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)

				vmID := "did:cash:subject#key-1"

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				// controller-1 does not exists
				req = *types.NewMsgRevokeVerification(didDoc.Id, vmID, signer)
			},
			true,
		},
		{
			"Pass: can revoke verification",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)

				vmID := "did:cash:subject#key-1"

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgRevokeVerification(didDoc.Id, vmID, signer)
			},
			false,
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("TestHandleMsgRevokeVerification#%v", i), func() {
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
			"FAIL: can not add service, identifier does not exist",
			func() {
				service := types.NewService(
					"service-id",
					"NonKYCCredential",
					"cash/multihash",
				)
				req = *types.NewMsgAddService("did:cash:subject", service, "did:cash:subject")
			},
			true,
		},
		{
			"FAIL: can not add service, identifier does not exist",
			func() {
				req = *types.NewMsgAddService("did:cash:subject", nil, "did:cash:subject")
			},
			true,
		},
		{
			"FAIL: cannot add service to did document (unauthorized, wrong relationship)",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipCapabilityInvocation, types.RelationshipCapabilityDelegation},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"KYCCredential",
					"cash/multihash",
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			true,
		},
		{
			"FAIL: cannot add service to did document with an incorrect type",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"NonKYCCredential",
					"cash/multihash",
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			true,
		},
		{
			"FAIL: duplicated service",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
					types.WithServices(
						types.NewService(
							"service-id",
							"KYCCredential",
							"cash/multihash",
						),
					),
				)

				service := types.NewService(
					"service-id",
					"KYCCredential",
					"cash/multihash",
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			true,
		},
		{
			"Pass: can add service to did document",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"KYCCredential",
					"cash/multihash",
				)

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddService(didDoc.Id, service, signer)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
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
			"FAIL: can not delete service, identifier does not exist",
			func() { req = *types.NewMsgDeleteService("did:cash:2222", "service-id", "did:cash:2222") },
			true,
		},
		{

			"Pass: can delete service from did document",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
					types.WithServices(
						types.NewService(
							"service-id",
							"KYCCredential",
							"cash/multihash",
						),
					),
				)

				serviceID := "service-id"

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, signer)
			},
			false,
		},
		{
			"FAIL: cannot remove an invalid serviceID",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipAuthentication},
							nil,
						),
					),
				)

				serviceID := ""

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, signer)
			},
			true,
		},
		{
			"FAIL: unauthorized (wrong relationship)",
			func() {
				signer := "subject"
				didDoc, _ := types.NewIdentifier(
					"did:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.RelationshipCapabilityInvocation},
							nil,
						),
					),
				)

				serviceID := "service-id"

				suite.keeper.SetIdentifier(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, signer)
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
