package keeper

import (
	"fmt"
	"time"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgSeverCreateIssuer() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgCreateIssuer

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: issuer can be created",
			func() {
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
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a license credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("seur", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// ACTUAL TEST
				req = *types.NewMsgCreateIssuer(
					vc.GetSubjectDID().String(),
					vc.Id,
					"seuro",
					100,
					suite.GetEMTiAddress().String(),
				)
			},
			true,
		},
		{
			"FAIL: signer not in provided did document",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin("seuro", circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgCreateIssuer(
					didDoc.Id,
					vcID,
					"seuro",
					100,
					"fail",
				)
			},
			false,
		},
		{
			"FAIL: verifiable credential not found in store",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred:2"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgCreateIssuer(
					didDoc.Id,
					vcID,
					"seuro",
					100,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer id in credential not correctly matching the provided did",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin("seuro", circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					"incorrect:did",
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgCreateIssuer(
					didDoc.Id,
					vc.Id,
					"seuro",
					100,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer already exists",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin("seuro", circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				issuer := types.Issuer{
					Token:     "seuro",
					Fee:       1,
					IssuerDid: didDoc.Id,
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				req = *types.NewMsgCreateIssuer(
					didDoc.Id,
					vc.Id,
					"seuro",
					100,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer token already exists",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin("seuro", circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				issuer := types.Issuer{
					Token:     "seuro",
					Fee:       1,
					IssuerDid: "didDoc.Id",
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				req = *types.NewMsgCreateIssuer(
					didDoc.Id,
					vc.Id,
					"seuro",
					100,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didResp, err := server.CreateIssuer(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didResp)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSeverBurnToken() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgBurnToken
	ctx := sdk.WrapSDKContext(suite.ctx) // ctx

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: issuer burns tokens",
			func() {
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
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a license credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("seur", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require an Issuer
				issuer := types.Issuer{
					Token:     "seuro",
					Fee:       1,
					IssuerDid: vc.GetSubjectDID().String(),
				}
				suite.keeper.SetIssuer(suite.ctx, issuer)
				// Require minted tokens
				amount, _ := sdk.ParseCoinNormalized("10seuro")
				mint := *types.NewMsgMintToken(
					vc.GetSubjectDID().String(),
					vc.Id,
					amount,
					suite.GetEMTiAddress().String(),
				)
				_, _ = server.MintToken(ctx, &mint)
				// ACTUAL TEST
				req = *types.NewMsgBurnToken(
					vc.GetSubjectDID().String(),
					vc.Id,
					amount,
					suite.GetEMTiAddress().String(),
				)
			},
			true,
		},
		{
			"FAIL: verifiable credential does not exist",
			func() {
				amount, _ := sdk.ParseCoinNormalized("999seuro")
				req = *types.NewMsgBurnToken(
					"did",
					"vcID",
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer did does not exist",
			func() {
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					"didDoc.Id",
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					"didDoc.Id",
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				amount, _ := sdk.ParseCoinNormalized("999seuro")
				req = *types.NewMsgBurnToken(
					"did:cosmos:cash",
					vc.Id,
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: token being burned is incorrect",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				issuer := types.Issuer{
					Token:     denom,
					Fee:       1,
					IssuerDid: didDoc.Id,
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				amount, _ := sdk.ParseCoinNormalized("900peuro")

				req = *types.NewMsgBurnToken(
					did,
					vcID,
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer does not have tokens to burn",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				issuer := types.Issuer{
					Token:     denom,
					Fee:       1,
					IssuerDid: didDoc.Id,
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				amount, _ := sdk.ParseCoinNormalized("10seuro")

				req = *types.NewMsgBurnToken(
					did,
					vcID,
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			didResp, err := server.BurnToken(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_MintToken() {
	// create the keeper
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgMintToken
	ctx := sdk.WrapSDKContext(suite.ctx)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: issuer mints tokens",
			func() {
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
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a license credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("seur", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require an Issuer
				issuer := types.Issuer{
					Token:     "seuro",
					Fee:       1,
					IssuerDid: vc.GetSubjectDID().String(),
				}
				suite.keeper.SetIssuer(suite.ctx, issuer)

				// ACTUAL TEST
				amount, _ := sdk.ParseCoinNormalized("999seuro")

				req = *types.NewMsgMintToken(
					vc.GetSubjectDID().String(),
					vc.Id,
					amount,
					suite.GetEMTiAddress().String(),
				)
			},
			true,
		},
		{
			"FAIL: verifiable credential does not exist",
			func() {
				amount, _ := sdk.ParseCoinNormalized("999seuro")
				req = *types.NewMsgMintToken(
					"did:cosmos:cash",
					"vcID",
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer did does not exist",
			func() {
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					"didDoc.Id",
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					"didDoc.Id",
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				amount, _ := sdk.ParseCoinNormalized("999seuro")
				req = *types.NewMsgMintToken(
					"did:cosmos:cash",
					vc.Id,
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: token being minted is incorrect",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				issuer := types.Issuer{
					Token:     denom,
					Fee:       1,
					IssuerDid: didDoc.Id}
				suite.keeper.SetIssuer(suite.ctx, issuer)
				amount, _ := sdk.ParseCoinNormalized("900peuro")

				req = *types.NewMsgMintToken(
					did,
					vcID,
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: reserve amount is less than minting amount",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					didDoc.Id,
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				issuer := types.Issuer{
					Token:     denom,
					Fee:       1,
					IssuerDid: didDoc.Id,
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				amount, _ := sdk.ParseCoinNormalized("1001seuro")

				req = *types.NewMsgMintToken(
					did,
					vcID,
					amount,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			didResp, err := server.MintToken(ctx, &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didResp)

			} else {
				suite.Require().Error(err)
				// TODO: check errors being returned are correct
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_PauseToken() {
	// create the keeper
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgPauseToken
	ctx := sdk.WrapSDKContext(suite.ctx)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: issuer paused token",
			func() {
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
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a license credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("seur", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require an Issuer
				issuer := types.Issuer{
					Token:     "seuro",
					Fee:       1,
					IssuerDid: vc.GetSubjectDID().String(),
				}
				suite.keeper.SetIssuer(suite.ctx, issuer)
				// ACTUAL TEST
				req = *types.NewMsgPauseToken(
					vc.GetSubjectDID().String(),
					vc.Id,
					suite.GetEMTiAddress().String(),
				)
			},
			true,
		},
		{
			"FAIL: verifiable credential does not exist",
			func() {
				req = *types.NewMsgPauseToken(
					"did:cosmos:cash",
					"vcID",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer did does not exist",
			func() {
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin(denom, circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					"didDoc.Id",
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					"didDoc.Id",
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				req = *types.NewMsgPauseToken(
					"did:cosmos:cash",
					vc.Id,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"FAIL: issuer id in credential not correctly matching the provided did",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
					didtypes.NewVerification(
						didtypes.NewVerificationMethod(
							"did:cosmos:cash:subject#key-1",
							"did:cosmos:cash:subject",
							didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
								didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
						),
						[]string{didtypes.Authentication},
						nil,
					),
				))
				circulationLimit, _ := sdk.NewIntFromString("1000")
				coin := sdk.NewCoin("seuro", circulationLimit)
				cs := vctypes.NewLicenseCredentialSubject(
					"incorrect:did",
					"MICAEMI",
					"IRL",
					"Another Financial Services Body (AFFB)",
					coin,
				)

				vc := vctypes.NewLicenseVerifiableCredential(
					vcID,
					didDoc.Id,
					time.Now(),
					cs,
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *types.NewMsgPauseToken(
					didDoc.Id,
					vc.Id,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
			},
			false,
		},
		{
			"PASS: issuer can pause a token",
			func() {
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
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a license credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("seur", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, suite.GetRegulatorAddress(),
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require an Issuer
				issuer := types.Issuer{
					Token:     "seuro",
					Fee:       1,
					IssuerDid: vc.GetSubjectDID().String(),
				}
				suite.keeper.SetIssuer(suite.ctx, issuer)
				// ACTUAL TEST
				req = *types.NewMsgPauseToken(
					vc.GetSubjectDID().String(),
					vc.Id,
					suite.GetEMTiAddress().String(),
				)
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			didResp, err := server.PauseToken(ctx, &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didResp)

			} else {
				suite.Require().Error(err)
				// TODO: check errors being returned are correct
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_IssueUserCredential() {
	// create the keeper
	server := NewMsgServerImpl(suite.keeper)
	var req vctypes.MsgIssueCredential
	ctx := sdk.WrapSDKContext(suite.ctx)

	testCases := []struct {
		msg      string
		malleate func()
		expErr   error
	}{
		{
			"PASS: can issue user credential",
			func() {
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
					suite.keyring, regulator,
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
					suite.keyring, regulator,
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require a license credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					regulatorDID.String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("seur", sdk.NewInt(1000)),
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, regulator,
					didtypes.NewVerificationMethodIDFromAddress(regulator.String()),
				)
				suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				// Require an Issuer (??)
				//issuer := types.Issuer{
				//	Token:     "seuro",
				//	Fee:       1,
				//	IssuerDid: vc.GetSubjectDID().String(),
				//}
				//suite.keeper.SetIssuer(suite.ctx, issuer)
				// ACTUAL TEST
				userCredentialIssuerAccount := suite.GetEMTiAddress()
				userCredentialIssuerAccountDID := didtypes.NewKeyDID(userCredentialIssuerAccount.String())
				vc = vctypes.NewUserVerifiableCredential(
					"user-credential-from-emti-to-alice",
					userCredentialIssuerAccountDID.String(),
					time.Now(),
					vctypes.NewUserCredentialSubject(
						didtypes.NewKeyDID(suite.GetAliceAddress().String()).String(),
						"some_garbled_zkp_stuff",
						true,
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, userCredentialIssuerAccount,
					didtypes.NewVerificationMethodIDFromAddress(userCredentialIssuerAccount.String()),
				)
				// ACTUAL TEST
				req = *vctypes.NewMsgIssueCredential(
					vc,
					userCredentialIssuerAccount.String(),
				)
			},
			nil,
		},
		{
			"FAIL: can issue user credential (not a licensed issuer)",
			func() {
				var vc vctypes.VerifiableCredential

				bobAccount := suite.GetBobAddress()
				// Require a license credential
				vc = vctypes.NewUserVerifiableCredential(
					"user-credential-from-bob-to-alice",
					didtypes.NewKeyDID(bobAccount.String()).String(),
					time.Now(),
					vctypes.NewUserCredentialSubject(
						didtypes.NewKeyDID(suite.GetAliceAddress().String()).String(),
						"some_garbled_zkp_stuff",
						true,
					),
				)
				vc, _ = vc.Sign(
					suite.keyring, bobAccount,
					didtypes.NewVerificationMethodIDFromAddress(bobAccount.String()),
				)
				// ACTUAL TEST
				req = *vctypes.NewMsgIssueCredential(
					vc,
					bobAccount.String(),
				)
			},
			types.ErrLicenseCredentialNotFound,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			didResp, err := server.IssueUserCredential(ctx, &req)
			if tc.expErr == nil {
				suite.NoError(err)
				suite.NotNil(didResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErr.Error())
			}
		})
	}
}
