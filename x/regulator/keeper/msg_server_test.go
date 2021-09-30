package keeper

import (
	"fmt"
	"time"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgSeverActivateRegulator() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgIssueRegulatorCredential

	testCases := []struct {
		msg       string
		malleate  func()
		expectErr error
	}{

		{
			msg: "PASS: regulator can be activated (ephemeral did)",
			malleate: func() {
				// regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				rvc := vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				// sign the credentials
				vmID := fmt.Sprint(regulatorDID, "#", regulator.String())
				rvc, _ = rvc.Sign(suite.keyring, suite.GetRegulatorAddress(), vmID)
				// send the message
				req = types.MsgIssueRegulatorCredential{
					Credential: &rvc,
					Owner:      regulator.String(),
				}
			},
		},
		{
			"PASS: regulator can be activated (persisted did)",
			func() {
				// regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// 3rd party did
				customDID := didtypes.NewChainDID(suite.ctx.ChainID(), "a-custom-did-id")
				customDIDDoc, _ := didtypes.NewDidDocument(
					customDID.String(),
					didtypes.WithVerifications(
						didtypes.NewAccountVerification(
							customDID,
							suite.ctx.ChainID(),
							// here we use a 3rd party address as licensed to execute
							// regulator transactions
							suite.GetAliceAddress().String(),
							didtypes.AssertionMethod,
						),
					),
				)
				// store the did
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(customDIDDoc.Id), customDIDDoc)
				// regulator verifiable credentials
				rvc := vctypes.NewRegulatorVerifiableCredential(
					"alice-regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						customDID.String(),
						"alice regulator",
						"EU",
					),
				)

				// sign the credentials
				vmID := fmt.Sprint(regulatorDID, "#", regulator.String())
				rvc, _ = rvc.Sign(suite.keyring, suite.GetRegulatorAddress(), vmID)
				// send the message
				req = types.MsgIssueRegulatorCredential{
					Credential: &rvc,
					Owner:      regulator.String(),
				}
			},
			nil,
		},
		{
			"FAIL: pubkey not found (ephemeral did)",
			func() {
				// regulator
				regulator := suite.GetRegulatorUnknownAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				rvc := vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				// sign the credentials
				vmID := fmt.Sprint(regulatorDID, "#", regulator.String())
				rvc, _ = rvc.Sign(suite.keyring, suite.GetRegulatorAddress(), vmID)
				// send the message
				req = types.MsgIssueRegulatorCredential{
					Credential: &rvc,
					Owner:      regulator.String(),
				}
			},
			vctypes.ErrMessageSigner,
		},
		{
			"FAIL: did not found (persistent did)",
			func() {
				// regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// 3rd party did
				customDID := didtypes.NewChainDID(suite.ctx.ChainID(), "a-custom-did-id-not-saved")
				// regulator verifiable credentials
				rvc := vctypes.NewRegulatorVerifiableCredential(
					"custom-regulator-credential-no-did",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						customDID.String(),
						"custom regulator",
						"EU",
					),
				)
				// sign the credentials
				vmID := fmt.Sprint(regulatorDID, "#", regulator.String())
				rvc, _ = rvc.Sign(suite.keyring, suite.GetRegulatorAddress(), vmID)
				// send the message
				req = types.MsgIssueRegulatorCredential{
					Credential: &rvc,
					Owner:      regulator.String(),
				}
			},
			didtypes.ErrDidDocumentNotFound,
		},
		{
			"FAIL: not a regulator account",
			func() {
				// a non-regulator address
				regulator := suite.GetAliceAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				rvc := vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				// sign the credentials
				vmID := fmt.Sprint(regulatorDID, "#", regulator.String())
				rvc, _ = rvc.Sign(suite.keyring, suite.GetRegulatorAddress(), vmID)
				// send the message
				req = types.MsgIssueRegulatorCredential{
					Credential: &rvc,
					Owner:      regulator.String(),
				}
			},
			types.ErrNotARegulator,
		},
		{
			"FAIL: verifiable credential proof invalid",
			func() {
				// a non-regulator address
				regulator := suite.GetAliceAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				rvc := vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				// sign the credentials
				vmID := fmt.Sprint(regulatorDID, "#", regulator.String())
				rvc, _ = rvc.Sign(suite.keyring, suite.GetRegulatorUnknownAddress(), vmID)
				// send the message
				req = types.MsgIssueRegulatorCredential{
					Credential: &rvc,
					Owner:      regulator.String(),
				}
			},
			types.ErrNotARegulator,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didResp, err := server.IssueRegulatorCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expectErr == nil {
				suite.NoError(err)
				suite.NotNil(didResp)
			} else {
				suite.Require().Error(err)
				suite.Assert().Contains(err.Error(), tc.expectErr.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSeverIssueRegistrationCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgIssueRegistrationCredential

	testCases := []struct {
		msg       string
		malleate  func()
		expectErr error
	}{

		{
			msg:       "PASS: issuer can be registered",
			expectErr: nil,
			malleate: func() {
				var vc vctypes.VerifiableCredential
				// Requires an active regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				vc = vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Actual test
				vc = vctypes.NewRegistrationVerifiableCredential(
					"registraion-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegistrationCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"EU",
						"emti",
						"E-Money Token Issuer",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				req = types.MsgIssueRegistrationCredential{
					Credential: &vc,
					Owner:      regulator.String(),
				}
			},
		},
		{
			msg:       "FAIL: issuer can be registered (issuer not a regulator)",
			expectErr: types.ErrNotARegulator,
			malleate: func() {
				// Requires an active regulator
				regulator := suite.GetRegulatorUnknownAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// Actual test
				vc := vctypes.NewRegistrationVerifiableCredential(
					"registraion-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegistrationCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"EU",
						"emti",
						"E-Money Token Issuer",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				req = types.MsgIssueRegistrationCredential{
					Credential: &vc,
					Owner:      regulator.String(),
				}
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didResp, err := server.IssueRegistrationCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expectErr == nil {
				suite.NoError(err)
				suite.NotNil(didResp)
			} else {
				suite.Require().Error(err)
				suite.Assert().Contains(err.Error(), tc.expectErr.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSeverIssueLicenseCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgIssueLicenseCredential

	//
	//

	testCases := []struct {
		msg       string
		malleate  func()
		expectErr error
	}{

		{
			msg:       "PASS: issuer can be licensed",
			expectErr: nil,
			malleate: func() {
				var vc vctypes.VerifiableCredential
				// Requires an active regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				vc = vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a registration credential
				vc = vctypes.NewRegistrationVerifiableCredential(
					"registration-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegistrationCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"EU",
						"emti",
						"E-Money Token Issuer",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				// Actual test
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"EURO",
						"EU",
						"An authority",
						sdk.NewCoin("sEUR", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)

				req = types.MsgIssueLicenseCredential{
					Credential: &vc,
					Owner:      regulator.String(),
				}
			},
		},
		{
			msg:       "FAIL: issuer can be licensed (no registration)",
			expectErr: types.ErrNotRegistered,
			malleate: func() {
				var vc vctypes.VerifiableCredential
				// Requires an active regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				vc = vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Registration credential not issued
				// Actual test
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-alice",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetAliceAddress().String()).String(),
						"EURO",
						"EU",
						"An authority",
						sdk.NewCoin("sEUR", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)

				req = types.MsgIssueLicenseCredential{
					Credential: &vc,
					Owner:      regulator.String(),
				}
			},
		},
		{
			msg:       "FAIL: issuer can be licensed (vc issuer not a regulator)",
			expectErr: types.ErrNotARegulator,
			malleate: func() {
				var vc vctypes.VerifiableCredential
				// Requires an active regulator
				regulator := suite.GetRegulatorAddress()
				regulatorDID := didtypes.NewKeyDID(regulator.String())
				// regulator verifiable credentials
				vc = vctypes.NewRegulatorVerifiableCredential(
					"regulator-credential",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegulatorCredentialSubject(
						regulatorDID.String(),
						"The Regulator",
						"EU",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a registration credential
				vc = vctypes.NewRegistrationVerifiableCredential(
					"registration-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewRegistrationCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"EU",
						"emti",
						"E-Money Token Issuer",
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				// Actual test
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					didtypes.NewKeyDID(suite.GetRegulatorUnknownAddress().String()).String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"EURO",
						"EU",
						"An authority",
						sdk.NewCoin("sEUR", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorUnknownAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)

				req = types.MsgIssueLicenseCredential{
					Credential: &vc,
					Owner:      regulator.String(),
				}
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didResp, err := server.IssueLicenseCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expectErr == nil {
				suite.NoError(err)
				suite.NotNil(didResp)
			} else {
				suite.Require().Error(err)
				suite.Assert().Contains(err.Error(), tc.expectErr.Error())
			}
		})
	}
}
