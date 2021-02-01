package keeper

import (
	"context"
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryValidators() {
	// TODO: add more test cases
	queryClient := suite.queryClient
	var req *types.QueryIdentifiersRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"empty request",
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
