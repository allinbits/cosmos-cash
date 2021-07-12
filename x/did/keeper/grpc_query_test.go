package keeper

import (
	"context"
	"fmt"

	"github.com/allinbits/cosmos-cash/x/did/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryDidDocuments() {
	queryClient := suite.queryClient
	var req *types.QueryDidDocumentsRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Pass: will return an empty array",
			func() {
				req = &types.QueryDidDocumentsRequest{}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didsResp, err := queryClient.DidDocuments(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryDidDocument() {
	queryClient := suite.queryClient
	var req *types.QueryDidDocumentRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Fail: will fail because no id is provided",
			func() {
				req = &types.QueryDidDocumentRequest{}
			},
			false,
		},
		{
			"Fail: will fail because no did is found",
			func() {
				req = &types.QueryDidDocumentRequest{
					Id: "did:cash:1234",
				}
			},
			false,
		},
		{
			"Pass: will pass because a did is found",
			func() {

				dd, _ := types.NewDidDocument("did:cash:1234")

				suite.keeper.SetDidDocument(
					suite.ctx,
					[]byte(dd.Id),
					dd,
				)
				req = &types.QueryDidDocumentRequest{
					Id: "did:cash:1234",
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didsResp, err := queryClient.DidDocument(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
