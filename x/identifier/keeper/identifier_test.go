package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (suite *KeeperTestSuite) TestIdentifierKeeperSetAndGet() {
	testCases := []struct {
		msg     string
		didFn   func() types.DidDocument
		expPass bool
	}{
		{
			"data stored successfully",
			func() types.DidDocument {
				dd, _ := types.NewIdentifier("did:cash:subject")
				return dd
			},
			true,
		},
	}
	for _, tc := range testCases {
		dd := tc.didFn()

		suite.keeper.SetIdentifier(suite.ctx, []byte(dd.Id), dd)
		suite.keeper.SetIdentifier(suite.ctx, []byte(dd.Id+"1"), dd)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.GetIdentifier(
					suite.ctx,
					[]byte(dd.Id),
				)
				suite.Require().True(found)

				allEntities := suite.keeper.GetAllIdentifiers(
					suite.ctx,
				)
				suite.Require().Equal(2, len(allEntities))
			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}
