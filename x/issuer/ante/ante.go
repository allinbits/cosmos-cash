package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/ante"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential/keeper"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// CheckUserCredentialsDecorator checks the users has a UserCredential in a preprocessing hook
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
// it will validate the User credential
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

			// if the issuer cannot be found or is not associated with the token; continue
			if !found || msg.Amount[0].Denom != issuer.Token {
				return next(ctx, tx, simulate)

			}

			// if the issuer has paused the token block the transaction
			if issuer.Paused {
				return ctx, sdkerrors.Wrapf(
					types.ErrBankSendDisabled,
					"the token being send has been blocked",
				)
			}

			vcs := cicd.vcsk.GetAllVerifiableCredentialsWithCondition(ctx, []byte(issuer.IssuerDid), func(vc vctypes.VerifiableCredential) bool {
				if vc.Issuer == issuer.IssuerDid {
					return true
				}
				return false
			})

			if found {
				// validate that kyc credentials have been issued to the `FromAddress`
				err = cicd.validateUserCredential(ctx, msg.FromAddress, vcs)
				if err != nil {
					return ctx, err
				}

				// validate that kyc credentials have been issued to the `ToAddress`
				err = cicd.validateUserCredential(ctx, msg.ToAddress, vcs)
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

// validateUserCredential validates a users User credential when they try to send a token
// to another user, this is called on every bank send message
func (cicd CheckUserCredentialsDecorator) validateUserCredential(
	ctx sdk.Context,
	address string,
	vcs []vctypes.VerifiableCredential,
) error {
	hasUserCredential := false

	a, _ := sdk.AccAddressFromBech32(address)

	account := cicd.accountk.GetAccount(ctx, a)
	if account == nil {
		return sdkerrors.Wrapf(
			types.ErrPublicKeyNotFound,
			"user has not created a did and has no public key associated with their account",
		)
	}

	pubkey := account.GetPubKey()
	dids := cicd.ik.GetDidDocumentsByPubKey(ctx, pubkey)

	// TODO: this is brute force, find a better way
	if len(dids) > 0 {
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
	}

	// TODO: is this check needed
	// hasUserCredential = vc.Validate(pubkey)

	if !hasUserCredential {
		return sdkerrors.Wrapf(
			types.ErrIncorrectUserCredential,
			"did document does not have a User credential to send e-money tokens",
		)
	}

	return nil
}
