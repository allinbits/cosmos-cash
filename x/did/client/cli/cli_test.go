package cli_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/allinbits/cosmos-cash/v2/x/did/client/cli"
	"github.com/allinbits/cosmos-cash/v2/x/did/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/v2/app"
	"github.com/allinbits/cosmos-cash/v2/app/params"
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

// TODO: write more test cases for happy paths

func (s *IntegrationTestSuite) TestGetCmdQueryDidDocuments() {
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
			&types.QueryDidDocumentsResponse{},
			&types.QueryDidDocumentsResponse{},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryIdentifers()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				s.Require().Equal(tc.expected.String(), tc.respType.String())

			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdQueryDidDocument() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
		respType  proto.Message
		expected  proto.Message
	}{
		{
			"DID Document not found",
			[]string{
				"not-existing-did",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			true,
			&types.QueryDidDocumentResponse{},
			&types.QueryDidDocumentResponse{},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryIdentifer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				s.Require().Equal(tc.expected.String(), tc.respType.String())

			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewCreateDidDocumentCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"valid transaction",
			[]string{
				"123456789abcdefghijk",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewCreateDidDocumentCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddVerificationCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"invalid transaction",
			[]string{
				"123456789abcdefghijk",
				`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AhJhB4NzRr2+pRpW4jDfajpML2h9yuBONsSqz6aXKZ6s"}`,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewAddVerificationCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// TODO: optimize this
			err = s.network.WaitForNextBlock()
			s.Require().NoError(err)
			err = s.network.WaitForNextBlock()
			s.Require().NoError(err)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewSetVerificationRelationshipsCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"invalid transaction",
			[]string{
				"123456789abcdefghijk",
				"cosmos1qsu9djdhcp052vyr44vax0ku99s86h85xu8n7r",
				fmt.Sprintf("--relationship=%s", types.CapabilityDelegation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
		{
			"can add transaction",
			[]string{
				"123456789abcdefghijk",
				"cosmos1qsu9djdhcp052vyr44vax0ku99s86h85xu8n7r",
				fmt.Sprintf("--relationship=%s", types.CapabilityDelegation),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewSetVerificationRelationshipCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(err)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewRevokeVerificationCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"invalid transaction",
			[]string{
				"123456789abcdefghijk",
				"cosmos1qsu9djdhcp052vyr44vax0ku99s86h85xu8n7r", // TODO when using generated ids this is difficult to test
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewRevokeVerificationCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(err)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddServiceCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"invalid transaction",
			[]string{
				"123456789abcdefghijk",
				"service:seuro",
				"DIDComm",
				"service:euro/SIGNATURE",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewAddServiceCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// TODO: optimize this
			err = s.network.WaitForNextBlock()
			s.Require().NoError(err)
			err = s.network.WaitForNextBlock()
			s.Require().NoError(err)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewDeleteServiceCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"invalid transaction",
			[]string{
				"123456789abcdefghijk",
				"service:seuro",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewDeleteServiceCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(err)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
