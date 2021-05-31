package keeper

import (
	"fmt"
	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
)

func (suite *KeeperTestSuite) TestVerifiableCredentialsKeeperSetAndGet() {
	testCases := []struct {
		msg string
		vc  types.VerifiableCredential
		// TODO: add mallate func and clean up test
		expPass bool
	}{
		{
			"data stored successfully",
			types.NewUserVerifiableCredential(
				"did:cash:1111",
				[]string{"context"},
				"",
				"",
				types.NewUserCredentialSubject("", "root", true),
				types.NewProof("", "", "", "", ""),
			),
			true,
		},
	}
	for _, tc := range testCases {
		suite.keeper.SetVerifiableCredential(
			suite.ctx,
			[]byte(tc.vc.Id),
			tc.vc,
		)
		suite.keeper.SetVerifiableCredential(
			suite.ctx,
			[]byte(tc.vc.Id+"1"),
			tc.vc,
		)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.GetVerifiableCredential(
					suite.ctx,
					[]byte(tc.vc.Id),
				)
				suite.Require().True(found)

				array := suite.keeper.GetAllVerifiableCredentials(
					suite.ctx,
				)

				suite.Require().Equal(2, len(array))
			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}
