package cli_test

import (
	"fmt"
	"testing"

	"github.com/allinbits/cosmos-cash/v2/app"
	"github.com/allinbits/cosmos-cash/v2/app/params"
	didcli "github.com/allinbits/cosmos-cash/v2/x/did/client/cli"
	"github.com/allinbits/cosmos-cash/v2/x/issuer/client/cli"
	"github.com/allinbits/cosmos-cash/v2/x/issuer/types"
	regcli "github.com/allinbits/cosmos-cash/v2/x/regulator/client/cli"
	regulatortypes "github.com/allinbits/cosmos-cash/v2/x/regulator/types"
	vccli "github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/client/cli"
	vctypes "github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
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

// SetupSuite executes bootstrapping logic before all the tests,
//i.e. once before the entire suite, start executing.

func (s *IntegrationTestSuite) SetupSuite() {

	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()

	types.RegisterInterfaces(cfg.InterfaceRegistry)

	cfg.AppConstructor = NewAppConstructor(app.MakeEncodingConfig())

	cfg.NumValidators = 2

	regulatorsAddresses := []string{"cosmos1pca4v0qd66w5dqatwu2w6am33jfcwngy07j6ex"}

	regulatorGenState := regulatortypes.DefaultGenesis(regulatorsAddresses...)

	regulatorGenStateBz, err := cfg.Codec.MarshalJSON(regulatorGenState)

	cfg.GenesisState[regulatortypes.ModuleName] = regulatorGenStateBz

	/*check state set start
	regulatorsAddresses = []string {}
	regulatorGenState = regulatortypes.DefaultGenesis(regulatorsAddresses...)
	err = cfg.Codec.UnmarshalJSON(cfg.GenesisState[regulatortypes.ModuleName], regulatorGenState)
	check state set end*/

	s.cfg = cfg

	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)

	s.Require().NoError(err)

	regulatorPrivateKey := `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 713F87B32C5B666CC4DC3728A830570C
type: secp256k1

k1KIzcz2NvPPib3IPDN7BGkeqJ0qt7bVkUk6CKceK3gW063kevI22+CUo31SKpM+
sOJlyRAJHjgm2aL6Nu1YvJfHuzz0hGKh1YNBxR8=
=s0e8
-----END TENDERMINT PRIVATE KEY-----`

	issuerPrivateKey := `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 6F301D9A4289DB0576B2D0BFDABFAEB7
type: secp256k1

NYCjo99V0sRnXZmEIWcKv2kZZsbMNMQmwHwVM2I4KF5INaFD5yevGyE5o1VBbAe/
pJYczoHm+YuvraF2Ecm8JHQwEf/phLIzx71blGI=
=Z3Cw
-----END TENDERMINT PRIVATE KEY-----`

	userPrivateKey := `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 95BDE423EAC7CFFBDFB9735832984EE9
type: secp256k1

5olJGlYhBLGPnW65otoEMlNYvdygoakQvebcY1OsUa0g/CSqI6yl01YQzk9/6bJj
PgEOZyH2lWJMtUyFIIMaL8Fs7ig26tRZ2WIATnM=
=e334
-----END TENDERMINT PRIVATE KEY-----`

	val := s.network.Validators[0]

	clientCtx := val.ClientCtx

	err = clientCtx.Keyring.ImportPrivKey("regulator", regulatorPrivateKey, "12345678")

	s.Require().NoError(err)

	err = clientCtx.Keyring.ImportPrivKey("issuer", issuerPrivateKey, "12345678")

	s.Require().NoError(err)

	err = clientCtx.Keyring.ImportPrivKey("user", userPrivateKey, "12345678")

	s.Require().NoError(err)

	args := []string{
		val.Address.String(),
		"cosmos1pca4v0qd66w5dqatwu2w6am33jfcwngy07j6ex",
		"1000stake",
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := bankcli.NewSendTxCmd()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}

	args = []string{
		val.Address.String(),
		"cosmos1vtcyr6c4d5sg2cmyx4jrfqlxzuy6esjteezwqm",
		"1000stake",
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd = bankcli.NewSendTxCmd()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}

	args = []string{
		val.Address.String(),
		"cosmos12875rn2j3chnyc4fmqmrs0n2wlt4sxl4a04ga6",
		"1000stake",
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd = bankcli.NewSendTxCmd()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}
}

// TearDownSuite performs cleanup logic after all the tests, i.e. once after the
// entire suite, has finished executing.
func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func addNewDidDoc(s *IntegrationTestSuite, val *network.Validator, identifier string, from string) {

	clientCtx := val.ClientCtx
	args := []string{
		identifier,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := didcli.NewCreateDidDocumentCmd()
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

func activateRegulator(s *IntegrationTestSuite, val *network.Validator) {

	clientCtx := val.ClientCtx

	args := []string{
		"FCA",
		"UK",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "regulator"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := regcli.ActivateCmd()
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

func createNewIssuerCredentials(s *IntegrationTestSuite, val *network.Validator, issuerDid string, shortName string, longName string, currency string, limit string) {

	clientCtx := val.ClientCtx

	args := []string{
		"registration-credential-" + issuerDid,
		"did:cosmos:key:cosmos1pca4v0qd66w5dqatwu2w6am33jfcwngy07j6ex",
		"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
		"UK",
		shortName,
		longName,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "regulator"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := regcli.IssueRegistrationCredentialCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}
	response := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())

	//license
	args = []string{
		"license-credential-" + issuerDid,
		"did:cosmos:key:cosmos1pca4v0qd66w5dqatwu2w6am33jfcwngy07j6ex",
		"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
		"license-credential-subject-type",
		"UK",
		"FCA",
		currency,
		limit,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "regulator"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd = regcli.IssueLicenseCredentialCmd()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	// wait for blocks
	for i := 0; i < 2; i++ {
		netError := s.network.WaitForNextBlock()
		s.Require().NoError(netError)
	}
	response = &sdk.TxResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())

	cmd = vccli.GetCmdQueryVerifiableCredentials()
	args = []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	addedResponse := &vctypes.QueryVerifiableCredentialsResponse{}
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())

}

func createNewIssuer(s *IntegrationTestSuite, val *network.Validator, issuerDid string, tokenName string, fee string) {

	clientCtx := val.ClientCtx
	args := []string{
		"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
		"license-credential-" + issuerDid,
		tokenName,
		fee,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	cmd := cli.NewCreateIssuerCmd()
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

func (s *IntegrationTestSuite) TestGetCmdQueryIssuers() {

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	issuerDid := "issuer1_did"

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			"PASS: query all issuers",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			&types.QueryIssuersResponse{},
			func() {
			},
		},
		{
			"PASS: query all issuers after creating 1 new issuer",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			&types.QueryIssuersResponse{},
			func() {
				activateRegulator(s, val)
				addNewDidDoc(s, val, issuerDid, "issuer")
				createNewIssuerCredentials(s, val, issuerDid, "issuer_short_name", "issuer_long_name", "GBP", "1000000")
				createNewIssuer(s, val, issuerDid, "token1", "1")
			},
		},
	}

	var first bool = true
	var size = 0
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			cmd := cli.GetCmdQueryIssuers()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			queryResponse := tc.respType.(*types.QueryIssuersResponse)
			issuers := queryResponse.GetIssuers()
			if first {
				first = false
				size = len(issuers)
			} else {
				s.Require().Greater(len(issuers), 0)
				s.Require().Equal(size+1, len(issuers))
				found := false
				for _, issuer := range issuers {
					if issuer.IssuerDid == "did:cosmos:net:"+clientCtx.ChainID+":"+issuerDid {
						found = true
						break
					}
				}
				s.Require().True(found)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewCreateIssuerCmd() {

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	issuerDid := "issuer2_did"

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			"PASS: create new issuer",
			[]string{
				"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
				"license-credential-" + issuerDid,
				"token2",
				"1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
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
				activateRegulator(s, val)
				addNewDidDoc(s, val, issuerDid, "issuer")
				createNewIssuerCredentials(s, val, issuerDid, "issuer_short_name", "issuer_long_name", "GBP", "1000000")
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()

			//pull out the set of created issuers
			cmd := cli.GetCmdQueryIssuers()
			args := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			response := &types.QueryIssuersResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			issuers := response.GetIssuers()
			size := len(issuers)

			cmd = cli.NewCreateIssuerCmd()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			//pull out the set of created issuers
			cmd = cli.GetCmdQueryIssuers()
			args = []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			issuers = response.GetIssuers()
			s.Require().Equal(size+1, len(issuers))
			found := false
			for _, issuer := range issuers {
				if issuer.IssuerDid == "did:cosmos:net:"+clientCtx.ChainID+":"+issuerDid {
					found = true
					break
				}
			}
			s.Require().True(found)
		})
	}
}

func (s *IntegrationTestSuite) TestNewIssueUserVerifiableCredentialCmd() {

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	issuerDid := "issuer_did"
	userDid := "user3_did"

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			"PASS: issue user verifiable credential",
			[]string{
				"did:cosmos:net:" + clientCtx.ChainID + ":" + userDid,
				"user-credential-" + userDid,
				"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
				"secret",
				"100",
				"100",
				"10000",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
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
				activateRegulator(s, val)
				addNewDidDoc(s, val, issuerDid, "issuer")
				createNewIssuerCredentials(s, val, issuerDid, "issuer_short_name", "issuer_long_name", "GBP", "1000000")
				createNewIssuer(s, val, issuerDid, "token3", "1")
				addNewDidDoc(s, val, userDid, "user")
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()

			cmd := vccli.GetCmdQueryVerifiableCredentials()
			args := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			addedResponse := &vctypes.QueryVerifiableCredentialsResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())
			vcs := addedResponse.GetVcs()
			size := len(vcs)

			cmd = cli.NewIssueUserVerifiableCredentialCmd()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			cmd = vccli.GetCmdQueryVerifiableCredentials()
			args = []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			addedResponse = &vctypes.QueryVerifiableCredentialsResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())
			vcs = addedResponse.GetVcs()
			s.Require().Equal(size+1, len(vcs))
			found := false
			for _, vc := range vcs {
				if vc.Issuer == "did:cosmos:net:"+clientCtx.ChainID+":"+issuerDid {
					found = true
					s.Require().Equal("did:cosmos:net:"+clientCtx.ChainID+":"+userDid, vc.GetSubjectDID().String())
					break
				}
			}
			s.Require().True(found)
		})
	}
}

func (s *IntegrationTestSuite) TestNewPauseTokenCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	issuerDid := "issuer4_did"

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			"PASS: pause token",
			[]string{
				"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
				"license-credential-" + issuerDid,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
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
				activateRegulator(s, val)
				addNewDidDoc(s, val, issuerDid, "issuer")
				createNewIssuerCredentials(s, val, issuerDid, "issuer_short_name", "issuer_long_name", "GBP", "1000000")
				createNewIssuer(s, val, issuerDid, "token4", "1")
			},
		},
		{
			"PASS: unpause token",
			[]string{
				"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
				"license-credential-" + issuerDid,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf(
					"--%s=%s",
					flags.FlagFees,
					sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
				),
			},
			&sdk.TxResponse{},
			func() {},
		},
	}

	var firstTest bool = true
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()

			cmd := cli.GetCmdQueryIssuers()
			args := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			response := &types.QueryIssuersResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			issuers := response.GetIssuers()
			found := false
			for _, issuer := range issuers {
				if issuer.IssuerDid == "did:cosmos:net:"+clientCtx.ChainID+":"+issuerDid {
					found = true
					if firstTest {
						s.Require().False(issuer.Paused)
					} else {
						s.Require().True(issuer.Paused)
					}
					break
				}
			}
			s.Require().True(found)
			cmd = cli.NewPauseTokenCmd()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			cmd = cli.GetCmdQueryIssuers()
			args = []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			response = &types.QueryIssuersResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), response), out.String())
			issuers = response.GetIssuers()
			found = false
			for _, issuer := range issuers {
				if issuer.IssuerDid == "did:cosmos:net:"+clientCtx.ChainID+":"+issuerDid {
					found = true
					if firstTest {
						s.Require().True(issuer.Paused)
					} else {
						s.Require().False(issuer.Paused)
					}
					break
				}
			}
			s.Require().True(found)
			firstTest = false
		})
	}
}

func (s *IntegrationTestSuite) TestNewMintTokenCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	issuerDid := "issuer5_did"

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			"PASS: mint token",
			[]string{
				"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
				"license-credential-" + issuerDid,
				"1000token5",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
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
				activateRegulator(s, val)
				addNewDidDoc(s, val, issuerDid, "issuer")
				createNewIssuerCredentials(s, val, issuerDid, "issuer_short_name", "issuer_long_name", "GBP", "1000000")
				createNewIssuer(s, val, issuerDid, "token5", "1")
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()
			found := false
			cmd := cli.NewMintTokenCmd()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			args := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			cmd = bankcli.GetCmdQueryTotalSupply()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			addedResponse := &banktypes.QueryTotalSupplyResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())
			supply := addedResponse.GetSupply()
			for _, supp := range supply {
				if supp.Denom == "token5" {
					found = true
					var expected int64 = 1000
					s.Require().Equal(expected, supp.Amount.Int64())
				}
			}
			s.Require().True(found)
		})
	}
}

func (s *IntegrationTestSuite) TestNewBurnTokenCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	issuerDid := "issuer6_did"

	testCases := []struct {
		name     string
		args     []string
		respType proto.Message
		malleate func()
	}{
		{
			"PASS: burn token",
			[]string{
				"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
				"license-credential-" + issuerDid,
				"900token6",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
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
				activateRegulator(s, val)
				addNewDidDoc(s, val, issuerDid, "issuer")
				createNewIssuerCredentials(s, val, issuerDid, "issuer_short_name", "issuer_long_name", "GBP", "1000000")
				createNewIssuer(s, val, issuerDid, "token6", "1")
				args := []string{
					"did:cosmos:net:" + clientCtx.ChainID + ":" + issuerDid,
					"license-credential-" + issuerDid,
					"1000token6",
					fmt.Sprintf("--%s=%s", flags.FlagFrom, "issuer"),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf(
						"--%s=%s",
						flags.FlagFees,
						sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String(),
					),
				}
				cmd := cli.NewMintTokenCmd()
				out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
				s.Require().NoError(err)
				// wait for blocks
				for i := 0; i < 2; i++ {
					netError := s.network.WaitForNextBlock()
					s.Require().NoError(netError)
				}
				addedResponse := &sdk.TxResponse{}
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())
			},
		},
	}

	found1 := false
	found2 := false
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.malleate()

			args := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			cmd := bankcli.GetCmdQueryTotalSupply()
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			addedResponse := &banktypes.QueryTotalSupplyResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())
			supply := addedResponse.GetSupply()
			for _, supp := range supply {
				if supp.Denom == "token6" {
					found1 = true
					var expected int64 = 1000
					s.Require().Equal(expected, supp.Amount.Int64())
				}
			}
			s.Require().True(found1)

			cmd = cli.NewBurnTokenCmd()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			// wait for blocks
			for i := 0; i < 2; i++ {
				netError := s.network.WaitForNextBlock()
				s.Require().NoError(netError)
			}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			args = []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			cmd = bankcli.GetCmdQueryTotalSupply()
			out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			addedResponse = &banktypes.QueryTotalSupplyResponse{}
			s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), addedResponse), out.String())
			supply = addedResponse.GetSupply()
			for _, supp := range supply {
				if supp.Denom == "token6" {
					found2 = true
					var expected int64 = 100
					s.Require().Equal(expected, supp.Amount.Int64())
				}
			}
			s.Require().True(found2)
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
