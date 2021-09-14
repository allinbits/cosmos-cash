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
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
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
					vc.Id,
					"seuro",
					100,
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
				)
			},
			true,
		},
		{
			"FAIL: signer not in provided did document",
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
				req = *types.NewMsgCreateIssuer(
					didDoc.Id,
					"any",
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
				vcID := "did:cosmos:cash:issuercred"
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
			"FAIL: issuer id not correctly matching the provided did",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
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
					"fail",
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
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
	ctx := sdk.WrapSDKContext(suite.ctx)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: issuer burns tokens",
			func() {
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
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

				mint := *types.NewMsgMintToken(
					did,
					vcID,
					amount,
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
				)
				_, _ = server.MintToken(ctx, &mint)

				req = *types.NewMsgBurnToken(
					did,
					vcID,
					amount,
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
				did := "did:cosmos:cash:subject"
				vcID := "did:cosmos:cash:issuercred"
				denom := "seuro"
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
				amount, _ := sdk.ParseCoinNormalized("999seuro")

				req = *types.NewMsgMintToken(
					did,
					vcID,
					amount,
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
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

				req = *types.NewMsgMintToken(
					did,
					vcID,
					amount,
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
							"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
							didtypes.DIDVMethodTypeCosmosAccountAddress,
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
					"cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a",
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
			}
		})
	}
}
