package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

func (suite *KeeperTestSuite) TestIdentifierKeeperSetAndGet() {
	testCases := []struct {
		msg        string
		identifier types.DidDocument
		expPass    bool
	}{
		{
			"data stored successfully",
			types.DidDocument{
				"context",
				"did:cash:1111",
				nil,
				nil,
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.keeper.SetIdentifier(suite.ctx, []byte(tc.identifier.Id), tc.identifier)
		suite.keeper.SetIdentifier(suite.ctx, []byte(tc.identifier.Id+"1"), tc.identifier)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.GetIdentifier(
					suite.ctx,
					[]byte(tc.identifier.Id),
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
