package ante

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	ct "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	server "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	issuerkeeper "github.com/allinbits/cosmos-cash/x/issuer/keeper"
	issuertypes "github.com/allinbits/cosmos-cash/x/issuer/types"
	vckeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential/keeper"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// TestAccount represents an account used in the tests in x/auth/ante.
type TestAccount struct {
	acc  authtypes.AccountI
	priv cryptotypes.PrivKey
}

func (suite AnteTestSuite) GetEMTiAddress() sdk.Address {
	return suite.GetKeyAddress("emti")
}

func (suite AnteTestSuite) GetAliceAddress() sdk.Address {
	return suite.GetKeyAddress("alice")
}

func (suite AnteTestSuite) GetBobAddress() sdk.Address {
	return suite.GetKeyAddress("bob")
}

func (suite AnteTestSuite) GetRegulatorAddress() sdk.Address {
	return suite.GetKeyAddress("regulator")
}

func (suite AnteTestSuite) GetRegulatorUnknownAddress() sdk.Address {
	return suite.GetKeyAddress("regulator-anon")
}

func (suite AnteTestSuite) GetKeyAddress(uid string) sdk.Address {
	i, _ := suite.keyring.Key(uid)
	return i.GetAddress()
}

// Keeper test suit enables the keeper package to be tested
type AnteTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	cucd          CheckUserCredentialsDecorator
	vckeeper      vckeeper.Keeper
	didkeeper     didkeeper.Keeper
	issuerkeeper  issuerkeeper.Keeper
	accountkeeper authkeeper.AccountKeeper
	bankkeeper    bankkeeper.Keeper
	txBuilder     client.TxBuilder
	clientCtx     client.Context
	keyring       keyring.Keyring
}

// SetupTest creates a test suite to test the issuer
func (suite *AnteTestSuite) SetupTest() {
	keyIssuer := sdk.NewKVStoreKey(issuertypes.StoreKey)
	memKeyIssuer := sdk.NewKVStoreKey(issuertypes.MemStoreKey)
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramtypes.StoreKey)
	memKeyParams := sdk.NewKVStoreKey(paramtypes.TStoreKey)
	keyIdentifier := sdk.NewKVStoreKey(didtypes.StoreKey)
	memKeyIdentifier := sdk.NewKVStoreKey(didtypes.MemStoreKey)
	keyVcs := sdk.NewKVStoreKey(vctypes.StoreKey)
	memKeyVcs := sdk.NewKVStoreKey(vctypes.MemStoreKey)

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

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "foochainid"}, true, server.ZeroLogWrapper{log.Logger})

	interfaceRegistry := ct.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
		issuertypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
		banktypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
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

	DidKeeper := didkeeper.NewKeeper(
		marshaler,
		keyIdentifier,
		memKeyIdentifier,
	)

	VcsKeeper := vckeeper.NewKeeper(
		marshaler,
		keyVcs,
		memKeyVcs,
		DidKeeper,
		AccountKeeper,
	)

	IssuerKeeper := issuerkeeper.NewKeeper(
		marshaler,
		keyIssuer,
		memKeyIssuer,
		BankKeeper,
		DidKeeper,
		VcsKeeper,
	)

	var authanteAccountKeeper authante.AccountKeeper
	authanteAccountKeeper = AccountKeeper

	cucd := NewCheckUserCredentialsDecorator(
		authanteAccountKeeper,
		*IssuerKeeper,
		*DidKeeper,
		*VcsKeeper,
	)

	suite.keyring = keyring.NewInMemory()

	// helper func to register accounts in the keychain and the account keeper
	registerAccount := func(uid string, withPubKey bool) {
		i, _, _ := suite.keyring.NewMnemonic(uid, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
		a := AccountKeeper.NewAccountWithAddress(ctx, i.GetAddress())
		if withPubKey {
			a.SetPubKey(i.GetPubKey())
		}
		AccountKeeper.SetAccount(ctx, AccountKeeper.NewAccount(ctx, a))
	}

	registerAccount("regulator", true)
	registerAccount("alice", true)
	registerAccount("regulator-anon", false)
	registerAccount("bob", true)
	registerAccount("emti", true) // e-money token issuer

	suite.clientCtx = client.Context{}.
		WithTxConfig(simapp.MakeTestEncodingConfig().TxConfig)
	suite.ctx, suite.cucd = ctx, cucd
	suite.accountkeeper, suite.bankkeeper = AccountKeeper, BankKeeper
	suite.vckeeper, suite.didkeeper, suite.issuerkeeper = *VcsKeeper, *DidKeeper, *IssuerKeeper
}

// CreateTestAccounts creates `numAccs` accounts, and return all relevant
// information about them including their private keys.
func (suite *AnteTestSuite) CreateTestAccounts(numAccs int) []TestAccount {
	var accounts []TestAccount

	for i := 0; i < numAccs; i++ {
		priv, pub, addr := KeyTestPubAddr()
		acc := suite.accountkeeper.NewAccountWithAddress(suite.ctx, addr)
		err := acc.SetAccountNumber(uint64(i))
		acc.SetPubKey(pub)
		suite.Require().NoError(err)
		suite.accountkeeper.SetAccount(suite.ctx, acc)

		accounts = append(accounts, TestAccount{acc, priv})
	}

	return accounts
}

func (suite *AnteTestSuite) CreateTestCredentials(testaccount TestAccount, kr keyring.Keyring, credID, issuerDID string) {
	address := testaccount.acc.GetAddress()
	did := didtypes.NewKeyDID(address.String())
	vc := vctypes.NewUserVerifiableCredential(
		credID,
		issuerDID,
		time.Now(),
		vctypes.NewUserCredentialSubject(
			did.String(),
			"root",
			true,
		),
	)
	vc, _ = vc.Sign(nil,
		address,
		didtypes.NewVerificationMethodIDFromAddress(address.String()),
	)
	suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
}

// KeyTestPubAddr generates a new secp256k1 keypair.
func KeyTestPubAddr() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(AnteTestSuite))
}

func (suite *AnteTestSuite) TestCheckUserCredentialDecorator() {
	var tx sdk.Tx
	var simulate bool
	var msgs []sdk.Msg
	var errExp error

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"PASS: issuer is not associated with the token",
			func() {

				accounts := suite.CreateTestAccounts(2)
				coins, _ := sdk.ParseCoinsNormalized("1000sEUR")
				msg := banktypes.NewMsgSend(accounts[0].acc.GetAddress(), accounts[1].acc.GetAddress(), coins)

				msgs = []sdk.Msg{msg}
				suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

				suite.txBuilder.SetMsgs(msgs...)
				tx = suite.txBuilder.GetTx()
				simulate = false
				errExp = nil
			},
			true,
		},
		{
			"PASS: two kyc'd users can exchange emoney tokens",
			func() {
				var vc vctypes.VerifiableCredential
				var err error
				// require issuer credential
				vc = vctypes.NewLicenseVerifiableCredential(
					"license-credential-for-emti",
					didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
					time.Now(),
					vctypes.NewLicenseCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"MICAEMI",
						"IRL",
						"Another Financial Services Body (AFFB)",
						sdk.NewCoin("sEUR", sdk.NewInt(100000000)),
					),
				)
				vc, err = vc.Sign(
					suite.keyring, suite.GetEMTiAddress(),
					didtypes.NewVerificationMethodIDFromAddress(suite.GetEMTiAddress().String()),
				)
				suite.Require().NoError(err)
				err = suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.Require().NoError(err)
				// require emti user credential
				vc = vctypes.NewUserVerifiableCredential(
					"emti-user-cred",
					didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
					time.Now(),
					vctypes.NewUserCredentialSubject(
						didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
						"root",
						true,
					),
				)
				vc, err = vc.Sign(suite.keyring,
					suite.GetEMTiAddress(),
					didtypes.NewVerificationMethodIDFromAddress(suite.GetEMTiAddress().String()),
				)
				suite.Require().NoError(err)
				err = suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.Require().NoError(err)
				// require bob credential
				vc = vctypes.NewUserVerifiableCredential(
					"bob-user-cred",
					didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
					time.Now(),
					vctypes.NewUserCredentialSubject(
						didtypes.NewKeyDID(suite.GetBobAddress().String()).String(),
						"root",
						true,
					),
				)
				vc, err = vc.Sign(suite.keyring,
					suite.GetEMTiAddress(),
					didtypes.NewVerificationMethodIDFromAddress(suite.GetEMTiAddress().String()),
				)
				suite.Require().NoError(err)
				err = suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
				suite.Require().NoError(err)
				// create the issuer
				issuer := issuertypes.Issuer{
					Token:     "sEUR",
					Fee:       1,
					IssuerDid: didtypes.NewKeyDID(suite.GetEMTiAddress().String()).String(),
					Paused:    false,
				}
				suite.issuerkeeper.SetIssuer(suite.ctx, issuer)
				// mint coins
				coins, _ := sdk.ParseCoinsNormalized("10000sEUR")
				suite.bankkeeper.MintCoins(suite.ctx, banktypes.ModuleName, coins)

				emti, _ := sdk.AccAddressFromBech32(suite.GetEMTiAddress().String())
				suite.bankkeeper.SendCoinsFromModuleToAccount(suite.ctx, banktypes.ModuleName, emti, coins)

				bob, _ := sdk.AccAddressFromBech32(suite.GetBobAddress().String())

				msg := banktypes.NewMsgSend(emti, bob, coins)

				msgs = []sdk.Msg{msg}
				suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

				suite.txBuilder.SetMsgs(msgs...)
				tx = suite.txBuilder.GetTx()
				simulate = false
				errExp = nil
			},
			true,
		},
		//{
		//	"FAIL: user has paused emoney token",
		//	func() {
		//		did := "did:cosmos:cash:subject"
		//		vcID := "did:cosmos:cash:issuercred"
		//		issuerAddress, _ := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
		//		didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
		//			didtypes.NewVerification(
		//				didtypes.NewVerificationMethod(
		//					"did:cosmos:cash:subject#key-1",
		//					"did:cosmos:cash:subject",
		//					didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
		//						didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
		//				),
		//				[]string{didtypes.Authentication},
		//				nil,
		//			),
		//		))
		//		circulationLimit, _ := sdk.NewIntFromString("1000")
		//		coin := sdk.NewCoin("seuro", circulationLimit)
		//		cs := vctypes.NewLicenseCredentialSubject(
		//			didDoc.Id,
		//			"MICAEMI",
		//			"IRL",
		//			"Another Financial Services Body (AFFB)",
		//			coin,
		//		)
		//
		//		vc := vctypes.NewLicenseVerifiableCredential(
		//			vcID,
		//			didDoc.Id,
		//			time.Now(),
		//			cs,
		//		)
		//		suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
		//		suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
		//
		//		accounts := suite.CreateTestAccounts(2)
		//		suite.CreateTestCredentials(accounts[0], "user1", "kyccred1", didDoc.Id)
		//		suite.CreateTestCredentials(accounts[1], "user2", "kyccred2", didDoc.Id)
		//
		//		issuer := issuertypes.Issuer{
		//			Token:     "sEUR",
		//			Fee:       1,
		//			IssuerDid: didDoc.Id,
		//			Paused:    true,
		//		}
		//
		//		suite.issuerkeeper.SetIssuer(suite.ctx, issuer)
		//
		//		coins, _ := sdk.ParseCoinsNormalized("10000sEUR")
		//		suite.bankkeeper.MintCoins(suite.ctx, banktypes.ModuleName, coins)
		//		suite.bankkeeper.SendCoinsFromModuleToAccount(suite.ctx, banktypes.ModuleName, issuerAddress, coins)
		//
		//		sendingCoins, _ := sdk.ParseCoinsNormalized("10sEUR")
		//		msg := banktypes.NewMsgSend(accounts[0].acc.GetAddress(), accounts[1].acc.GetAddress(), sendingCoins)
		//
		//		msgs = []sdk.Msg{msg}
		//		suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
		//
		//		suite.txBuilder.SetMsgs(msgs...)
		//		tx = suite.txBuilder.GetTx()
		//		simulate = false
		//		errExp = sdkerrors.Wrapf(
		//			issuertypes.ErrBankSendDisabled,
		//			"the token being send has been blocked",
		//		)
		//
		//	},
		//	false,
		//},
		//{
		//	"FAIL: from address does not have required did and credential",
		//	func() {
		//		did := "did:cosmos:cash:subject"
		//		vcID := "did:cosmos:cash:issuercred"
		//		issuerAddress, _ := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
		//		didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
		//			didtypes.NewVerification(
		//				didtypes.NewVerificationMethod(
		//					"did:cosmos:cash:subject#key-1",
		//					"did:cosmos:cash:subject",
		//					didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
		//						didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
		//				),
		//				[]string{didtypes.Authentication},
		//				nil,
		//			),
		//		))
		//		circulationLimit, _ := sdk.NewIntFromString("1000")
		//		coin := sdk.NewCoin("seuro", circulationLimit)
		//		cs := vctypes.NewLicenseCredentialSubject(
		//			didDoc.Id,
		//			"MICAEMI",
		//			"IRL",
		//			"Another Financial Services Body (AFFB)",
		//			coin,
		//		)
		//
		//		vc := vctypes.NewLicenseVerifiableCredential(
		//			vcID,
		//			didDoc.Id,
		//			time.Now(),
		//			cs,
		//		)
		//		suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
		//		suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
		//
		//		accounts := suite.CreateTestAccounts(2)
		//		suite.CreateTestCredentials(accounts[1], "user2", "kyccred2", didDoc.Id)
		//
		//		issuer := issuertypes.Issuer{
		//			Token:     "sEUR",
		//			Fee:       1,
		//			IssuerDid: didDoc.Id,
		//			Paused:    false,
		//		}
		//
		//		suite.issuerkeeper.SetIssuer(suite.ctx, issuer)
		//
		//		coins, _ := sdk.ParseCoinsNormalized("10000sEUR")
		//		suite.bankkeeper.MintCoins(suite.ctx, banktypes.ModuleName, coins)
		//		suite.bankkeeper.SendCoinsFromModuleToAccount(suite.ctx, banktypes.ModuleName, issuerAddress, coins)
		//
		//		sendingCoins, _ := sdk.ParseCoinsNormalized("10sEUR")
		//		msg := banktypes.NewMsgSend(accounts[0].acc.GetAddress(), accounts[1].acc.GetAddress(), sendingCoins)
		//
		//		msgs = []sdk.Msg{msg}
		//		suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
		//
		//		suite.txBuilder.SetMsgs(msgs...)
		//		tx = suite.txBuilder.GetTx()
		//		simulate = false
		//		errExp = sdkerrors.Wrapf(
		//			issuertypes.ErrIncorrectUserCredential,
		//			"did document does not have a User credential to send e-money tokens",
		//		)
		//
		//	},
		//	false,
		//},
		//{
		//	"FAIL: to address does not have required did and credential",
		//	func() {
		//		did := "did:cosmos:cash:subject"
		//		vcID := "did:cosmos:cash:issuercred"
		//		issuerAddress, _ := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
		//		didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
		//			didtypes.NewVerification(
		//				didtypes.NewVerificationMethod(
		//					"did:cosmos:cash:subject#key-1",
		//					"did:cosmos:cash:subject",
		//					didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
		//						didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
		//				),
		//				[]string{didtypes.Authentication},
		//				nil,
		//			),
		//		))
		//		circulationLimit, _ := sdk.NewIntFromString("1000")
		//		coin := sdk.NewCoin("seuro", circulationLimit)
		//		cs := vctypes.NewLicenseCredentialSubject(
		//			didDoc.Id,
		//			"MICAEMI",
		//			"IRL",
		//			"Another Financial Services Body (AFFB)",
		//			coin,
		//		)
		//
		//		vc := vctypes.NewLicenseVerifiableCredential(
		//			vcID,
		//			didDoc.Id,
		//			time.Now(),
		//			cs,
		//		)
		//		suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
		//		suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
		//
		//		accounts := suite.CreateTestAccounts(2)
		//		suite.CreateTestCredentials(accounts[0], "user1", "kyccred1", didDoc.Id)
		//
		//		issuer := issuertypes.Issuer{
		//			Token:     "sEUR",
		//			Fee:       1,
		//			IssuerDid: didDoc.Id,
		//			Paused:    false,
		//		}
		//
		//		suite.issuerkeeper.SetIssuer(suite.ctx, issuer)
		//
		//		coins, _ := sdk.ParseCoinsNormalized("10000sEUR")
		//		suite.bankkeeper.MintCoins(suite.ctx, banktypes.ModuleName, coins)
		//		suite.bankkeeper.SendCoinsFromModuleToAccount(suite.ctx, banktypes.ModuleName, issuerAddress, coins)
		//
		//		sendingCoins, _ := sdk.ParseCoinsNormalized("10sEUR")
		//		msg := banktypes.NewMsgSend(accounts[0].acc.GetAddress(), accounts[1].acc.GetAddress(), sendingCoins)
		//
		//		msgs = []sdk.Msg{msg}
		//		suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
		//
		//		suite.txBuilder.SetMsgs(msgs...)
		//		tx = suite.txBuilder.GetTx()
		//		simulate = false
		//		errExp = sdkerrors.Wrapf(
		//			issuertypes.ErrIncorrectUserCredential,
		//			"did document does not have a User credential to send e-money tokens",
		//		)
		//
		//	},
		//	false,
		//},
		//{
		//	"FAIL: to address does not have a public key in the account store",
		//	func() {
		//		did := "did:cosmos:cash:subject"
		//		vcID := "did:cosmos:cash:issuercred"
		//		issuerAddress, _ := sdk.AccAddressFromBech32("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
		//		didDoc, _ := didtypes.NewDidDocument(did, didtypes.WithVerifications(
		//			didtypes.NewVerification(
		//				didtypes.NewVerificationMethod(
		//					"did:cosmos:cash:subject#key-1",
		//					"did:cosmos:cash:subject",
		//					didtypes.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
		//						didtypes.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019),
		//				),
		//				[]string{didtypes.Authentication},
		//				nil,
		//			),
		//		))
		//		circulationLimit, _ := sdk.NewIntFromString("1000")
		//		coin := sdk.NewCoin("seuro", circulationLimit)
		//		cs := vctypes.NewLicenseCredentialSubject(
		//			didDoc.Id,
		//			"MICAEMI",
		//			"IRL",
		//			"Another Financial Services Body (AFFB)",
		//			coin,
		//		)
		//
		//		vc := vctypes.NewLicenseVerifiableCredential(
		//			vcID,
		//			didDoc.Id,
		//			time.Now(),
		//			cs,
		//		)
		//		suite.vckeeper.SetVerifiableCredential(suite.ctx, []byte(vc.Id), vc)
		//		suite.didkeeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
		//
		//		accounts := suite.CreateTestAccounts(1)
		//		suite.CreateTestCredentials(accounts[0], "user1", "kyccred1", didDoc.Id)
		//
		//		issuer := issuertypes.Issuer{
		//			Token:     "sEUR",
		//			Fee:       1,
		//			IssuerDid: didDoc.Id,
		//			Paused:    false,
		//		}
		//
		//		suite.issuerkeeper.SetIssuer(suite.ctx, issuer)
		//
		//		coins, _ := sdk.ParseCoinsNormalized("10000sEUR")
		//		suite.bankkeeper.MintCoins(suite.ctx, banktypes.ModuleName, coins)
		//		suite.bankkeeper.SendCoinsFromModuleToAccount(suite.ctx, banktypes.ModuleName, issuerAddress, coins)
		//
		//		sendingCoins, _ := sdk.ParseCoinsNormalized("10sEUR")
		//		acc, _ := sdk.AccAddressFromBech32("cosmos1c3dmkzyjyj2gs7zcp5qjq40js963a0q7sxrtxj")
		//		msg := banktypes.NewMsgSend(accounts[0].acc.GetAddress(), acc, sendingCoins)
		//
		//		msgs = []sdk.Msg{msg}
		//		suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
		//
		//		suite.txBuilder.SetMsgs(msgs...)
		//		tx = suite.txBuilder.GetTx()
		//		simulate = false
		//		errExp = sdkerrors.Wrapf(
		//			issuertypes.ErrPublicKeyNotFound,
		//			"user has not created a did and has no public key associated with their account",
		//		)
		//
		//	},
		//	false,
		//},
	}
	for _, tc := range testCases {
		tc.malleate()
		antehandler := sdk.ChainAnteDecorators(suite.cucd)

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			_, err := antehandler(
				suite.ctx,
				tx,
				simulate,
			)

			if tc.expPass {
				suite.NoError(err)
				suite.NoError(errExp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(err.Error(), errExp.Error())
				suite.Require().False(tc.expPass)
			}
		})
	}
}
