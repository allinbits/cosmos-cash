package keeper

import (
	"context"
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryIdentifiers() {
	queryClient := suite.queryClient
	var req *types.QueryIdentifiersRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Pass: will return an empty array",
			func() {
				req = &types.QueryIdentifiersRequest{}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			identifiersResp, err := queryClient.Identifiers(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(identifiersResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryIdentifier() {
	queryClient := suite.queryClient
	var req *types.QueryIdentifierRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Fail: will fail because no id is provided",
			func() {
				req = &types.QueryIdentifierRequest{}
			},
			false,
		},
		{
			"Fail: will fail because no did is found",
			func() {
				req = &types.QueryIdentifierRequest{
					Id: "did:cash:1234",
				}
			},
			false,
		},
		// {
		// 	"Pass: will pass because a did is found",
		// 	func() {
		// 		suite.keeper.SetIdentifier(
		// 			suite.ctx,
		// 			[]byte("did:cash:1234"),
		// 			types.DidDocument{
		// 				"context",
		// 				"did:cash:1234",
		// 				nil,
		// 				nil,
		// 			},
		// 		)
		// 		req = &types.QueryIdentifierRequest{
		// 			Id: "did:cash:1234",
		// 		}
		// 	},
		// 	true,
		// },
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			identifiersResp, err := queryClient.Identifier(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(identifiersResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
