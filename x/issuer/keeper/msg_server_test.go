package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
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
		// TODO: uncomment when the latest version of the cosmos-sdk is released
		// Fixed by PR https://github.com/cosmos/cosmos-sdk/pull/9229
		//	{
		//		"PASS: issuer can be created",
		//		func() { req = *types.NewMsgCreateIssuer("seuro", 100, "did:cash:1111") },
		//		true,
		//	},
		{
			"FAIL: issuer already exists",
			func() {
				issuer := types.Issuer{
					Token:   "seuro",
					Fee:     1,
					Address: "did:cash:1111",
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				req = *types.NewMsgCreateIssuer("seuro", 100, "did:cash:1111")
			},
			false,
		},
		{
			"FAIL: issuer token already exists",
			func() {
				issuer := types.Issuer{
					Token:   "seuro",
					Fee:     1,
					Address: "did:cash:1112",
				}

				suite.keeper.SetIssuer(suite.ctx, issuer)
				req = *types.NewMsgCreateIssuer("seuro", 100, "did:cash:1111")
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

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"FAIL: issuer does not exist so tokens cannot be burned",

			func() { req = *types.NewMsgBurnToken(sdk.NewCoin("eeuro", sdk.NewInt(100)), "did:cash:1111") },
			false,
		},
		// TODO: uncomment when the latest version of the cosmos-sdk is released
		// Fixed by PR https://github.com/cosmos/cosmos-sdk/pull/9229
		//		{
		//			"PASS: issuer burns tokens",
		//			func() {
		//				issuer := *types.NewMsgCreateIssuer("seuro", 999, "did:cash:1111")
		//
		//				server.CreateIssuer(sdk.WrapSDKContext(suite.ctx), &issuer)
		//
		//				req = *types.NewMsgBurnToken(999, "did:cash:1111")
		//			},
		//			true,
		//		},
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
	// owner := "cosmos1m26ukcnpme38enptw85w2twcr8gllnj8anfy6a"
	ctx := sdk.WrapSDKContext(suite.ctx)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"FAIL: issuer does not exist so tokens cannot be burned",
			func() { req = *types.NewMsgMintToken(sdk.NewCoin("cash", sdk.NewInt(100)), "did:cash:1111") },
			false,
		},
		// TODO: add successful case when 0.43 is done
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
