package keeper

import (
	"fmt"
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (suite *KeeperTestSuite) TestMsgSeverActivateRegulator() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgActivate

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
				req = types.MsgActivate{
					Credentials: &rvc,
					Creator:     regulator.String(),
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
				req = types.MsgActivate{
					Credentials: &rvc,
					Creator:     regulator.String(),
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
				req = types.MsgActivate{
					Credentials: &rvc,
					Creator:     regulator.String(),
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
				req = types.MsgActivate{
					Credentials: &rvc,
					Creator:     regulator.String(),
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
				req = types.MsgActivate{
					Credentials: &rvc,
					Creator:     regulator.String(),
				}
			},
			types.ErrNotARegulator,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didResp, err := server.Activate(sdk.WrapSDKContext(suite.ctx), &req)
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
