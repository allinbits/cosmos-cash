package keeper

import (
	//"context"
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (suite *KeeperTestSuite) TestMsgSeverIdentifiers() {
	// TODO: write tests for msg server

	//msgServer := suite.msgServer
	var msg *types.MsgCreateIdentifier

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"empty request",
			func() {
				auth := types.NewAuthentication(
					"accAddrBech32",
					"sepk256",
					"id",
					"pubKey.Address().String()",
				)

				msg = types.NewMsgCreateIdentifier(
					"id",
					types.Authentications{&auth},
					"accAddrBech32",
				)
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			fmt.Println(msg)
			// TODO: fix msgServer handler`

			//createIdentifierResp, err := msgServer.CreateIdentifier(context.Background(), msg)
			//if tc.expPass {
			//	suite.NoError(err)
			//	suite.NotNil(createIdentifierResp)

			//} else {
			//	suite.Require().Error(err)
			//}
		})
	}
}
