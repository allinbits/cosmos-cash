package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryVerifiableCredentials() {
	queryClient := suite.queryClient
	var req *types.QueryVerifiableCredentialsRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Pass: will return an empty array",
			func() {
				req = &types.QueryVerifiableCredentialsRequest{}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didsResp, err := queryClient.VerifiableCredentials(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryVerifiableCredential() {
	queryClient := suite.queryClient
	var req *types.QueryVerifiableCredentialRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Fail: will fail because no id is provided",
			func() {
				req = &types.QueryVerifiableCredentialRequest{}
			},
			false,
		},
		{
			"Fail: will fail because no vc is found",
			func() {
				req = &types.QueryVerifiableCredentialRequest{
					VerifiableCredentialId: "vc:cash:1234",
				}
			},
			false,
		},
		{
			"Pass: will pass because a vc is found",
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
				suite.keeper.SetVerifiableCredential(
					suite.ctx,
					[]byte(vc.Id),
					vc,
				)
				req = &types.QueryVerifiableCredentialRequest{
					VerifiableCredentialId: vc.Id,
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didsResp, err := queryClient.VerifiableCredential(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryValidateVerifiableCredential() {
	queryClient := suite.queryClient
	var req *types.QueryValidateVerifiableCredentialRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Fail: will fail because no id is provided",
			func() {
				req = &types.QueryValidateVerifiableCredentialRequest{}
			},
			false,
		},
		{
			"Fail: will fail because no vc is found",
			func() {
				req = &types.QueryValidateVerifiableCredentialRequest{
					VerifiableCredentialId: "vc:cash:1234",
				}
			},
			false,
		},
		{
			"Pass: will pass because a vc is found and valid",
			func() {
				// NOTE: The signature is hardcoded here for simplicity, if changing this test in the
				// future please find a better way of generting the signature
				issuerPubkey := "cosmospub1addwnpepqg86fqwehwcjndtcluzs32eel0m4lcsghx6dkxrreqyrey4w7aqju0ks4l8"
				signature := "HjxPB1hv/iFjnvA5c3GGSfxi8YyzaKb8qqvBk8yBPa1DPG1VV/JzowVlTSyaO3YBBxAlb6dpRfeP8SfkHjcNQQ=="

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

				p := types.NewProof(
					"sepk256",
					fmt.Sprintf("%s", "tm"),
					"assertionMethod",
					"accAddrBech32"+"#keys-1",
					signature,
				)
				vc.Proof = &p
				suite.keeper.SetVerifiableCredential(
					suite.ctx,
					[]byte(vc.Id),
					vc,
				)
				req = &types.QueryValidateVerifiableCredentialRequest{
					VerifiableCredentialId: vc.Id,
					IssuerPubkey:           issuerPubkey,
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didsResp, err := queryClient.ValidateVerifiableCredential(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
