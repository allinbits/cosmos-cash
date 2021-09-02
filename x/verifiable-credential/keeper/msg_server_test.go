package keeper

import (
	"fmt"
	"time"

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
					"accAddr",
					time.Now(),
					cs,
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
					"accAddr",
					time.Now(),
					cs,
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

func (suite *KeeperTestSuite) TestMsgSeverDeleteVerifableCredential() {
	server := NewMsgServerImpl(suite.keeper)
	var req types.MsgDeleteVerifiableCredential

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"correctly deletes vc",
			func() {
				cs := types.NewUserCredentialSubject(
					"accAddr",
					"root",
					true,
				)

				vc := types.NewUserVerifiableCredential(
					"new-verifiable-cred-3",
					"did:cash:1111",
					time.Now(),
					cs,
				)
				suite.keeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)

				req = *types.NewMsgDeleteVerifiableCredential(vc.Id, vc.Issuer, "did:cash:1111")
			},
			true,
		},
		{
			"vc does not exists",
			func() {
				req = *types.NewMsgDeleteVerifiableCredential(
					"new-verifiable-cred-3",
					"did:cosmos:cash:issuer",
					"did:cash:1111",
				)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			vcResp, err := server.DeleteVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(vcResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
