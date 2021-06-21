package ante

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	identifierKeeper "github.com/allinbits/cosmos-cash/x/identifier/keeper"
	identifierTypes "github.com/allinbits/cosmos-cash/x/identifier/types"
	issuerKeeper "github.com/allinbits/cosmos-cash/x/issuer/keeper"
	issuerTypes "github.com/allinbits/cosmos-cash/x/issuer/types"
	vcsKeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/keeper"
	vcsTypes "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	dbm "github.com/tendermint/tm-db"
)

// Keeper test suit enables the keeper package to be tested
type KeeperTestSuite struct {
	suite.Suite

	ctx  sdk.Context
	cucd CheckUserCredentialsDecorator
}

// SetupTest creates a test suite to test the issuer
func (suite *KeeperTestSuite) SetupTest() {
	keyIssuer := sdk.NewKVStoreKey(issuerTypes.StoreKey)
	memKeyIssuer := sdk.NewKVStoreKey(issuerTypes.MemStoreKey)
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramtypes.StoreKey)
	memKeyParams := sdk.NewKVStoreKey(paramtypes.TStoreKey)
	keyIdentifier := sdk.NewKVStoreKey(identifierTypes.StoreKey)
	memKeyIdentifier := sdk.NewKVStoreKey(identifierTypes.MemStoreKey)
	keyVcs := sdk.NewKVStoreKey(vcsTypes.StoreKey)
	memKeyVcs := sdk.NewKVStoreKey(vcsTypes.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIssuer, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyIssuer, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyIdentifier, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyIdentifier, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyVcs, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyVcs, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, nil)

	interfaceRegistry := ct.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
		issuerTypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
	}

	allowedReceivingModAcc := map[string]bool{}

	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	paramsKeeper := paramskeeper.NewKeeper(marshaler, nil, keyParams, memKeyParams)

	AccountKeeper := authkeeper.NewAccountKeeper(
		marshaler,
		keyAcc,
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	BankKeeper := bankkeeper.NewBaseKeeper(
		marshaler,
		keyBank,
		AccountKeeper,
		paramsKeeper.Subspace(banktypes.ModuleName),
		blockedAddrs,
	)

	IdentifierKeeper := identifierKeeper.NewKeeper(
		marshaler,
		keyIdentifier,
		memKeyIdentifier,
	)

	VcsKeeper := vcsKeeper.NewKeeper(
		marshaler,
		keyVcs,
		memKeyVcs,
	)

	issuerKeeper := issuerKeeper.NewKeeper(
		marshaler,
		keyIssuer,
		memKeyIssuer,
		BankKeeper,
	)

	var authanteAccountKeeper authante.AccountKeeper
	authanteAccountKeeper = AccountKeeper

	cucd := NewCheckUserCredentialsDecorator(
		authanteAccountKeeper,
		*issuerKeeper,
		*IdentifierKeeper,
		*VcsKeeper,
	)

	suite.ctx, suite.cucd = ctx, cucd
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestCheckUserCredentialDecorator() {
	var tx sdk.Tx
	var simulate bool
	var ante sdk.AnteHandler

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"FAIL: transaction is nil",
			func() {
				tx = nil
				simulate = false
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				suite.cucd.AnteHandle(
					suite.ctx,
					tx,
					simulate,
					ante,
				)
			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}
