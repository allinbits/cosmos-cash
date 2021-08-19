package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgSeverCreateVerifableCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgCreateVerifiableCredential

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"correctly creates vc",
			func() {
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					[]string{"VerifiableCredential", "KYCCredential"},
					"accAddr",
					fmt.Sprintf("%s", "currentTome"),
					cs,
					types.NewProof("", "", "", "", ""),
				)
				req = *types.NewMsgCreateVerifiableCredential(vc, "did:cash:1111")
			},
			true,
		},
		{
			"vc already exists",
			func() {
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					[]string{"VerifiableCredential", "KYCCredential"},
					"accAddr",
					fmt.Sprintf("%s", "currentTome"),
					cs,
					types.NewProof("", "", "", "", ""),
				)
				suite.keeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)

				req = *types.NewMsgCreateVerifiableCredential(vc, "did:cash:1111")
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			vcResp, err := server.CreateVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(vcResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
