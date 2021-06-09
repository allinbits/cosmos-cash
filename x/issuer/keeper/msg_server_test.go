package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
			func() { req = *types.NewMsgBurnToken(999, "did:cash:1111") },
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

			identifierResp, err := server.BurnToken(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(identifierResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
