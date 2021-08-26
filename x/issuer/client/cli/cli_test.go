package cli_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/allinbits/cosmos-cash/x/issuer/client/cli"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	//"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	//sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/app"
	"github.com/allinbits/cosmos-cash/app/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	dbm "github.com/tendermint/tm-db"
)

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor(encodingCfg params.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return app.New(
			"cosmos-cash",
			val.Ctx.Logger,
			dbm.NewMemDB(), nil, true, make(map[int64]bool),
			val.Ctx.Config.RootDir,
			0,
			encodingCfg,
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// SetupSuite executes bootstrapping logic before all the tests, i.e. once before
// the entire suite, start executing.
func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
	cfg := network.DefaultConfig()
	types.RegisterInterfaces(cfg.InterfaceRegistry)
	cfg.AppConstructor = NewAppConstructor(app.MakeEncodingConfig())
	cfg.NumValidators = 2

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

// TearDownSuite performs cleanup logic after all the tests, i.e. once after the
// entire suite, has finished executing.
func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestGetCmdQueryIssuers() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
		respType  proto.Message
		expected  proto.Message
	}{
		{
			"json output",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			&types.QueryIssuersResponse{},
			&types.QueryIssuersResponse{},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryIssuers()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				s.Require().Equal(tc.expected.String(), tc.respType.String())

			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
