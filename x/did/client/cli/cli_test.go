package cli_test

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/allinbits/cosmos-cash/x/did/client/cli"
	"github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

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

func name() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func addnewdiddoc(s *IntegrationTestSuite, identifier string, val *network.Validator) {
	clientCtx := val.ClientCtx
	args := []string{
		identifier,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := cli.NewCreateDidDocumentCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}
	response := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
}

func (s *IntegrationTestSuite) TestGetCmdQueryDidDocuments() {
	identifier := "123456789abcdefghijk"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			name() + "_1",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			&types.QueryDidDocumentsResponse{},
			func() {},
		},
		{
			name() + "_2",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			&types.QueryDidDocumentsResponse{},
			func() { addnewdiddoc(s, identifier, val) },
		},
	}

	firstloop := true
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.GetCmdQueryIdentifers()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			queryresponse := tc.respType.(*types.QueryDidDocumentsResponse)
			diddocs := queryresponse.GetDidDocuments()
			if firstloop {
				s.Require().Equal(len(diddocs), 0)
				firstloop = false
			} else {
				s.Require().Greater(len(diddocs), 0)
				s.Require().Equal(diddocs[0].Id, "did:cosmos:net:"+clientCtx.ChainID+":"+identifier)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdQueryDidDocument() {
	identifier := "123456789abcdefghijk"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
			codes.NotFound,
			&types.QueryDidDocumentResponse{},
			func() {},
		},
		{
			name(),
			codes.OK,
			&types.QueryDidDocumentResponse{},
			func() {
				args := []string{
					identifier,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				}
				cmd := cli.NewCreateDidDocumentCmd()
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				s.Require().NoError(err)
				response := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.GetCmdQueryIdentifer()
			identifiertoquery := "did:cosmos:net:" + clientCtx.ChainID + ":" + identifier
			args := []string{
				identifiertoquery,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				queryresponse := tc.respType.(*types.QueryDidDocumentResponse)
				diddoc := queryresponse.GetDidDocument()
				s.Require().Equal(identifiertoquery, diddoc.Id)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewCreateDidDocumentCmd() {

	identifier := "123456789abcdefghijk_"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
	}{
		{
			name(),
			[]string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				)},
			&sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			for i := 0; i < 3; i++ {
				cmd := cli.NewCreateDidDocumentCmd()
				tc.args[0] = identifier + fmt.Sprint(i)
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				//pull out the just created document
				cmd = cli.GetCmdQueryIdentifer()
				identifiertoquery := "did:cosmos:net:" + clientCtx.ChainID + ":" + tc.args[0]
				args_temp := []string{
					identifiertoquery,
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
				s.Require().NoError(err)
				response1 := &types.QueryDidDocumentResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response1))
				s.Require().Equal(response1.GetDidDocument().Id, identifiertoquery)

				//pull out the set of created documentsq
				cmd = cli.GetCmdQueryIdentifers()
				args_temp = []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
				out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
				s.Require().NoError(err)
				response2 := &types.QueryDidDocumentsResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response2))
				diddocs := response2.GetDidDocuments()
				s.Require().Equal(len(diddocs), i+1)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewUpdateDidDocumentCmd() {
	identifier1 := "123456789abcdefghijk"
	identifier2 := "cosmos1kslgpxklq75aj96cz3qwsczr95vdtrd3p0fslp"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	malonce := true

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			name(),
			[]string{
				identifier1,
				identifier2,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			&sdk.TxResponse{},
			func() {
				cmd := cli.NewCreateDidDocumentCmd()
				args_temp := []string{
					identifier1,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				}
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				response := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
				s.Require().True(malonce)
				malonce = false
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewUpdateDidDocumentCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			identifiertoquery := "did:cosmos:net:" + clientCtx.ChainID + ":" + tc.args[0]
			args_temp := []string{
				identifiertoquery,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
			s.Require().NoError(err)
			response := &types.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			controller := response.GetDidDocument().Controller
			s.Require().Equal(len(controller), 1)
			s.Require().Equal(controller[0], "did:cosmos:key:"+identifier2)
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddVerificationCmd() {
	identifier := "123456789abcdefghijk"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	malonce := true

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
			[]string{
				identifier,
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
			codes.OK,
			&sdk.TxResponse{},
			func() {
				cmd := cli.NewCreateDidDocumentCmd()
				args_temp := []string{
					identifier,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				}
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				response := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
				identifier = "did:cosmos:net:" + clientCtx.ChainID + ":" + identifier
				s.Require().True(malonce)
				malonce = false
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewAddVerificationCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(err)
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			identifiertoquery := "did:cosmos:net:" + clientCtx.ChainID + ":" + tc.args[0]
			args_temp := []string{
				identifiertoquery,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
			s.Require().NoError(err)
			response := &types.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			authentications := response.GetDidDocument().Authentication
			verificationmethods := response.GetDidDocument().VerificationMethod
			s.Require().Equal(2, len(authentications))
			s.Require().Equal(2, len(verificationmethods))
			for i := 0; i < 2; i++ {
				s.Require().Equal(authentications[i], verificationmethods[i].Id)
				s.Require().Equal(identifiertoquery, verificationmethods[i].Controller)
			}

			verificationmethod := verificationmethods[1]
			s.Require().Equal("F02126107837346bdbea51a56e230df6a3a4c2f687dcae04e36c4aacfa697299eac", verificationmethod.GetPublicKeyMultibase())
		})
	}
}

func (s *IntegrationTestSuite) TestNewSetVerificationRelationshipsCmd() {
	identifier := "123456789abcdefghijk"
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	malonce := true

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
	}{
		{
			name(),
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
			codes.OK,
			&sdk.TxResponse{},
			func() {
				cmd := cli.NewCreateDidDocumentCmd()
				args_temp := []string{
					identifier,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				}
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				response := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
				identifier = "did:cosmos:net:" + clientCtx.ChainID + ":" + identifier
				s.Require().True(malonce)
				malonce = false
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
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
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//check for update
			cmd = cli.GetCmdQueryIdentifer()
			identifiertoquery := "did:cosmos:net:" + clientCtx.ChainID + ":" + tc.args[0]
			args_temp := []string{
				identifiertoquery,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args_temp)
			s.Require().NoError(err)
			response := &types.QueryDidDocumentResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response))
			authentications := response.GetDidDocument().Authentication
			verificationmethods := response.GetDidDocument().VerificationMethod
			s.Require().Equal(2, len(authentications))
			s.Require().Equal(2, len(verificationmethods))
			for i := 0; i < 2; i++ {
				s.Require().Equal(authentications[i], verificationmethods[i].Id)
				s.Require().Equal(identifiertoquery, verificationmethods[i].Controller)
			}

			verificationmethod := verificationmethods[1]
			s.Require().Equal("F02126107837346bdbea51a56e230df6a3a4c2f687dcae04e36c4aacfa697299eac", verificationmethod.GetPublicKeyMultibase())

		})
	}
}

func (s *IntegrationTestSuite) TestNewRevokeVerificationCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
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
			codes.OK,
			&sdk.TxResponse{},
			func() {},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewRevokeVerificationCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(err)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewAddServiceCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
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
			codes.OK,
			&sdk.TxResponse{},
			func() {},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.NewAddServiceCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// TODO: optimize this
			err = s.network.WaitForNextBlock()
			s.Require().NoError(err)
			err = s.network.WaitForNextBlock()
			s.Require().NoError(err)
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewDeleteServiceCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr codes.Code
		respType  proto.Message
		malleate  func()
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
			codes.OK,
			&sdk.TxResponse{},
			func() {},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
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
			if tc.expectErr != codes.OK {
				s.Require().Error(err)
				s.Equal(tc.expectErr, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
