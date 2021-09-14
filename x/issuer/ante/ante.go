package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/ante"

	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/allinbits/cosmos-cash/x/issuer/keeper"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential/keeper"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// CheckUserCredentialsDecorator checks the users has a KYCCredential in a preprocessing hook
type CheckUserCredentialsDecorator struct {
	accountk accountKeeper.AccountKeeper
	issuerk  keeper.Keeper
	ik       didkeeper.Keeper
	vcsk     vcskeeper.Keeper
}

func NewCheckUserCredentialsDecorator(
	accountk accountKeeper.AccountKeeper,
	issuerk keeper.Keeper,
	ik didkeeper.Keeper,
	vcsk vcskeeper.Keeper,
) CheckUserCredentialsDecorator {
	return CheckUserCredentialsDecorator{
		accountk: accountk,
		issuerk:  issuerk,
		ik:       ik,
		vcsk:     vcsk,
	}
}

// AnteHandle CheckUserCredentialsDecorator is used as a hook to intercept the bank send message then
// it will validate the KYC credential
func (cicd CheckUserCredentialsDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *bank.MsgSend:
			issuer, found := cicd.issuerk.GetIssuerByToken(ctx, []byte(msg.Amount[0].Denom))
			if !found {
				return ctx, sdkerrors.Wrapf(
					types.ErrDidDocumentDoesNotExist,
					"did does not exists when validating the KYC credential",
				)
			}

			if msg.Amount[0].Denom != issuer.Token {
				return next(ctx, tx, simulate)
			}

			vcs := cicd.vcsk.GetAllVerifiableCredentialsByIssuer(ctx, issuer.IssuerDid)

			if found {
				err := cicd.validateKYCCredential(ctx, msg.ToAddress, vcs)
				if err != nil {
					return ctx, err
				}

				err = cicd.validateKYCCredential(ctx, msg.FromAddress, vcs)
				if err != nil {
					return ctx, err
				}
			}
		case *bank.MsgMultiSend:
			// TODO: implement multi send checks
		}
	}

	return next(ctx, tx, simulate)
}

// validateKYCCredential validates a users KYC credential when they try to send a token
// to another user, this is called on every bank send message
func (cicd CheckUserCredentialsDecorator) validateKYCCredential(
	ctx sdk.Context,
	address string,
	vcs []vctypes.VerifiableCredential,
) error {
	hasUserCredential := false

	a, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	// NOTE: test on account with no pubkey
	account := cicd.accountk.GetAccount(ctx, a)
	pubkey := account.GetPubKey()

	dids := cicd.ik.GetDidDocumentsByPubKey(ctx, pubkey)

	// TODO: this is brute force, find a better way
	for _, vc := range vcs {
		for _, did := range dids {
			switch key := vc.CredentialSubject.(type) {
			case *vctypes.VerifiableCredential_UserCred:
				if key.UserCred.Id == did.Id {
					hasUserCredential = true
				}
			}
		}
	}

	// TODO: is this check needed
	// hasUserCredential = vc.Validate(pubkey)

	if !hasUserCredential {
		return sdkerrors.Wrapf(
			types.ErrIncorrectUserCredential,
			"did document does not have a KYC credential to send e-money tokens",
		)
	}

	return nil
}
