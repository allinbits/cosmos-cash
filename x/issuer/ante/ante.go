package ante

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/ante"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	identifierkeeper "github.com/allinbits/cosmos-cash/x/identifier/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/keeper"
	vcstypes "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
)

// CheckIssuerCredentialsDecorator checks the issuer has a EMILicense in a preprocessing hook
type CheckIssuerCredentialsDecorator struct {
	issuerk keeper.Keeper
	ik      identifierkeeper.Keeper
	vcsk    vcskeeper.Keeper
}

func NewCheckIssuerCredentialsDecorator(
	issuerk keeper.Keeper,
	ik identifierkeeper.Keeper,
	vcsk vcskeeper.Keeper,
) CheckIssuerCredentialsDecorator {
	return CheckIssuerCredentialsDecorator{
		issuerk: issuerk,
		ik:      ik,
		vcsk:    vcsk,
	}
}

func (cicd CheckIssuerCredentialsDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	// TODO: improve logic here
	for _, msg := range tx.GetMsgs() {
		if msg.Type() == "create-issuer" {
			imsg := msg.(*types.MsgCreateIssuer)
			didURI := "did:cash:" + imsg.Owner

			// TODO: pass in the did URI as an arg {msg.Id}
			// TODO: ensure this keeper can only read from store
			did, found := cicd.ik.GetIdentifier(ctx, []byte(didURI))
			if !found {
				return ctx, sdkerrors.Wrapf(
					types.ErrIssuerFound,
					"identifier does not exists",
				)
			}

			// TODO: optimize here
			foundKey := false
			for _, auth := range did.Authentication {
				if auth.Controller == imsg.Owner {
					foundKey = true
				}
			}
			if !foundKey {
				return ctx, sdkerrors.Wrapf(
					types.ErrIssuerFound,
					"msg sender not in auth array in did document",
				)
			}

			// TODO: optimize here
			// check if the did document has the issuer credential
			hasIssuerCredential := false
			for _, service := range did.Services {
				// TODO use enum here
				if service.Type == "IssuerCredential" {
					// TODO: ensure this keeper can only read from store
					vc, found := cicd.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
					if !found {
						return ctx, sdkerrors.Wrapf(
							types.ErrIssuerFound,
							"credential not found",
						)
					}

					issuerCred := vc.GetUserCred()
					if issuerCred.Id != didURI {
						return ctx, sdkerrors.Wrapf(
							types.ErrIssuerFound,
							"issuer id not correct",
						)
					}

					//	if issuerCred.IsVerified == false {
					//		return ctx, sdkerrors.Wrapf(
					//			types.ErrIssuerFound,
					//			"issuer is not verified",
					//		)
					//	}

					hasIssuerCredential = true
					// TODO: validate credential here has been issued by regulator
				}
			}
			if !hasIssuerCredential {
				return ctx, sdkerrors.Wrapf(
					types.ErrIssuerFound,
					"did document doesnt have a credential to create issuers",
				)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// CheckUserCredentialsDecorator checks the users has a KYCCredential in a preprocessing hook
type CheckUserCredentialsDecorator struct {
	accountk accountKeeper.AccountKeeper
	issuerk  keeper.Keeper
	ik       identifierkeeper.Keeper
	vcsk     vcskeeper.Keeper
}

func NewCheckUserCredentialsDecorator(
	accountk accountKeeper.AccountKeeper,
	issuerk keeper.Keeper,
	ik identifierkeeper.Keeper,
	vcsk vcskeeper.Keeper,
) CheckUserCredentialsDecorator {
	return CheckUserCredentialsDecorator{
		accountk: accountk,
		issuerk:  issuerk,
		ik:       ik,
		vcsk:     vcsk,
	}
}

func (cicd CheckUserCredentialsDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	// TODO: ensure this keepers can only read from store
	// TODO: improve logic here
	for _, msg := range tx.GetMsgs() {
		if msg.Type() == "send" {
			imsg := msg.(*bank.MsgSend)
			// FIXME: iterate over tokens and check the multi-send
			issuer, found := cicd.issuerk.GetIssuerByToken(ctx, []byte(imsg.Amount[0].Denom))

			if found {
				err := cicd.validateKYCCredential(ctx, imsg.ToAddress, issuer.Address)
				if err != nil {
					return ctx, err
				}

				err = cicd.validateKYCCredential(ctx, imsg.FromAddress, issuer.Address)
				if err != nil {
					return ctx, err
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}

func (cicd CheckUserCredentialsDecorator) validateKYCCredential(
	ctx sdk.Context,
	address string,
	issuerAddress string,
) error {
	didURI := "did:cash:" + address

	// TODO: tidy this functionality into the keeper,
	// GetIdentifierWithCondition, GetIdentifierWithService, GetIdentifierWithAuth
	did, found := cicd.ik.GetIdentifier(ctx, []byte(didURI))
	if !found {
		return sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"identifier does not exists",
		)
	}

	foundKey := false
	for _, auth := range did.Authentication {
		if auth.Controller == address {
			foundKey = true
			break
		}
	}
	if !foundKey {
		return sdkerrors.Wrapf(
			types.ErrUserFound,
			"msg sender not in auth slice in did document",
		)
	}

	// check if the did document has the issuer credential
	hasUserCredential := false
	for _, service := range did.Services {
		// TODO use enum here
		if service.Type != "KYCCredential" {
			continue
		}

		// TODO: ensure this keeper can only read from store
		vc, found := cicd.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
		if !found {
			continue
		}

		userCred := vc.GetUserCred()
		if userCred.Id != didURI {
			continue
		}

		if !userCred.IsVerified {
			continue
		}

		if vc.Issuer != issuerAddress {
			continue
		}

		address, err := sdk.AccAddressFromBech32(issuerAddress)
		if err != nil {
			continue
		}

		account := cicd.accountk.GetAccount(ctx, address)
		pubkey := account.GetPubKey()

		s, err := base64.StdEncoding.DecodeString(vc.Proof.Signature)
		if err != nil {
			continue
		}
		emptyProof := vcstypes.NewProof("", "", "", "", "")
		vc.Proof = &emptyProof

		// TODO: this is an expesive operation, could lead to DDOS
		// TODO: we can hash this and make this less expensive
		hasUserCredential = pubkey.VerifySignature(
			vc.GetBytes(),
			s,
		)

		break
	}
	if !hasUserCredential {
		return sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"did document does not have a credential to send e-money tokens",
		)
	}

	return nil
}
