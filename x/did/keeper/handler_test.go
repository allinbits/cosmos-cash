package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/did/types"
)

func (suite *KeeperTestSuite) TestHandleMsgCreateDidDocument() {
	var (
		req types.MsgCreateDidDocument
	)

	handleFn := NewHandler(suite.keeper)

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
			_, err := handleFn(suite.ctx, &req)
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

	handleFn := NewHandler(suite.keeper)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
	}{
		{
			"FAIL: not found",
			func() {
				req = *types.NewMsgUpdateDidDocument("did:cosmos:cash:subject", nil, "subject")
			},
			true,
		},
		{
			"FAIL: unauthorized",
			func() {

				signer := "subject"
				did := "did:cosmos:cash:subject"
				didDoc, _ := types.NewDidDocument(did)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgUpdateDidDocument(didDoc.Id, []string{"did:cosmos:cash:controller"}, signer)
			},
			true,
		},
		{
			"PASS: nil controllers",
			func() {
				signer := "subject"
				did := "did:cosmos:cash:subject"
				didDoc, _ := types.NewDidDocument(did, types.WithVerifications(
					types.NewVerification(
						types.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"EcdsaSecp256k1RecoveryMethod2020",
							"did:cosmos:cash:subject",
							signer,
						),
						[]string{types.Authentication},
						nil,
					),
				))
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *types.NewMsgUpdateDidDocument(did, nil, signer)
			},
			false,
		},
		{
			"FAIL: invalid controllers",
			func() {
				signer := "subject"
				did := "did:cosmos:cash:subject"
				didDoc, _ := types.NewDidDocument(did, types.WithVerifications(
					types.NewVerification(
						types.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"EcdsaSecp256k1RecoveryMethod2020",
							"did:cosmos:cash:subject",
							signer,
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

				req = *types.NewMsgUpdateDidDocument(did, controllers, signer)
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
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
			"FAIL: can not add verification, did does not exist",
			func() { req = *types.NewMsgAddVerification("did:cosmos:cash:subject", nil, "subject") },
			true,
		},
		{
			"FAIL: can not add verification, unauthorized",
			func() {
				// setup
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
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
						"EcdsaSecp256k1RecoveryMethod2020",
						"did:cosmos:cash:subject",
						"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
					),
					[]string{types.Authentication},
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
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
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
						"EcdsaSecp256k1RecoveryMethod2020",
						"did:cosmos:cash:subject",
						"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
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
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								signer,
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				v := types.NewVerification(
					types.NewVerificationMethod(
						"did:cosmos:cash:subject#key-2",
						"EcdsaSecp256k1RecoveryMethod2020",
						"did:cosmos:cash:subject",
						"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
					),
					[]string{types.Authentication},
					nil,
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgAddVerification(didDoc.Id, v, signer)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
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
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
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
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
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
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								signer,
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
					[]string{types.CapabilityInvocation},
					signer,
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
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
			"FAIL: can not revoke verification, did does not exist",
			func() {
				req = *types.NewMsgRevokeVerification("did:cosmos:cash:2222", "service-id", "did:cosmos:cash:2222")
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
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgRevokeVerification(didDoc.Id, "did:cosmos:cash:subject#not-existent", "subject")
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
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.Authentication},
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
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								signer,
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				vmID := "did:cosmos:cash:subject#key-1"

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
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
			"FAIL: can not add service, did does not exist",
			func() {
				service := types.NewService(
					"service-id",
					"NonKYCCredential",
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
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.CapabilityInvocation, types.CapabilityDelegation},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"KYCCredential",
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
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				service := types.NewService(
					"service-id",
					"NonKYCCredential",
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
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.Authentication},
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
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								signer,
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
					"KYCCredential",
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
			"FAIL: can not delete service, did does not exist",
			func() { req = *types.NewMsgDeleteService("did:cosmos:cash:2222", "service-id", "did:cosmos:cash:2222") },
			true,
		},
		{

			"Pass: can delete service from did document",
			func() {
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								signer,
							),
							[]string{types.Authentication},
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

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, signer)
			},
			false,
		},
		{
			"FAIL: cannot remove an invalid serviceID",
			func() {
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.Authentication},
							nil,
						),
					),
				)

				serviceID := ""

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgDeleteService(didDoc.Id, serviceID, signer)
			},
			true,
		},
		{
			"FAIL: unauthorized (wrong relationship)",
			func() {
				signer := "subject"
				didDoc, _ := types.NewDidDocument(
					"did:cosmos:cash:subject",
					types.WithVerifications(
						types.NewVerification(
							types.NewVerificationMethod(
								"did:cosmos:cash:subject#key-1",
								"EcdsaSecp256k1RecoveryMethod2020",
								"did:cosmos:cash:subject",
								"027560af3387d375e3342a6968179ef3c6d04f5d33b2b611cf326d4708badd7770",
							),
							[]string{types.CapabilityInvocation},
							nil,
						),
					),
				)

				serviceID := "service-id"

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
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
