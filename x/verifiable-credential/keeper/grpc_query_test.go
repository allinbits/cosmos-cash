package keeper

import (
	"context"
	"fmt"

	"github.com/allinbits/cosmos-cash/v3/x/verifiable-credential/types"
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
		//{
		//	"Pass: will pass because a vc is found",
		//	func() {
		//		cs := types.NewUserCredentialSubject(
		//			"accAddr",
		//			"root",
		//			true,
		//		)
		//
		//		vc := types.NewUserVerifiableCredential(
		//			"new-verifiable-cred-3",
		//			"accAddr",
		//			time.Now(),
		//			cs,
		//		)
		//		suite.keeper.SetVerifiableCredential(
		//			suite.ctx,
		//			[]byte(vc.Id),
		//			vc,
		//		)
		//		req = &types.QueryVerifiableCredentialRequest{
		//			VerifiableCredentialId: vc.Id,
		//		}
		//	},
		//	true,
		//},
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
