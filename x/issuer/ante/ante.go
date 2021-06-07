package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	identifierkeeper "github.com/allinbits/cosmos-cash/x/identifier/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/keeper"
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
	issuerk keeper.Keeper
	ik      identifierkeeper.Keeper
	vcsk    vcskeeper.Keeper
}

func NewCheckUserCredentialsDecorator(
	issuerk keeper.Keeper,
	ik identifierkeeper.Keeper,
	vcsk vcskeeper.Keeper,
) CheckUserCredentialsDecorator {
	return CheckUserCredentialsDecorator{
		issuerk: issuerk,
		ik:      ik,
		vcsk:    vcsk,
	}
}

func (cicd CheckUserCredentialsDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	// TODO: improve logic here
	for _, msg := range tx.GetMsgs() {
		if msg.Type() == "send" {
			imsg := msg.(*bank.MsgSend)
			didURI := "did:cash:" + imsg.ToAddress
			_ = "did:cash:" + imsg.FromAddress
			_, found := cicd.issuerk.GetIssuerByToken(ctx, []byte(imsg.Amount[0].Denom))

			if found {
				// TODO: pass in the did URI as an arg {msg.Id}
				// TODO: tidy this functionality into the keeper,
				// GetIdentifierWithCondition, GetIdentifierWithService, GetIdentifierWithAuth
				// TODO: ensure this keeper can only read from store
				// TODO: esure to check the both from and to addresses
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
					if auth.Controller == imsg.ToAddress {
						foundKey = true
					}
				}
				if !foundKey {
					return ctx, sdkerrors.Wrapf(
						types.ErrUserFound,
						"msg sender not in auth array in did document",
					)
				}

				// TODO: optimize here
				// check if the did document has the issuer credential
				hasUserCredential := false
				for _, service := range did.Services {
					// TODO use enum here
					if service.Type == "KYCCredential" {
						// TODO: ensure this keeper can only read from store
						vc, found := cicd.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
						if !found {
							return ctx, sdkerrors.Wrapf(
								types.ErrUserFound,
								"credential not found",
							)
						}

						userCred := vc.GetUserCred()
						if userCred.Id != didURI {
							return ctx, sdkerrors.Wrapf(
								types.ErrUserFound,
								"user id not correct",
							)
						}

						if userCred.IsVerified == false {
							return ctx, sdkerrors.Wrapf(
								types.ErrUserFound,
								"user is not verified",
							)
						}

						//if vc.Issuer != issuer.Address {
						//	return ctx, sdkerrors.Wrapf(
						//		types.ErrUserFound,
						//		"user is not verified",
						//	)
						//}

						hasUserCredential = true
						// TODO: validate credential here has been issued by issuer
					}
				}
				if !hasUserCredential {
					return ctx, sdkerrors.Wrapf(
						types.ErrIssuerFound,
						"did document doesnt have a credential to send e-money tokens",
					)
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
