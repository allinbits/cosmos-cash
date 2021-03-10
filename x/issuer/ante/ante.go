package ante

import (
	"fmt"
	identifierkeeper "github.com/allinbits/cosmos-cash/x/identifier/keeper"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/keeper"
	//identifiertypes "github.com/allinbits/cosmos-cash/x/identifier/types"
	"github.com/allinbits/cosmos-cash/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//bank "github.com/cosmos/cosmos-sdk/x/bank"
)

// CheckIssuerCredentialsDecorator deducts fees from the every send transaction
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

func (cicd CheckIssuerCredentialsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		if msg.Type() == "create-issuer" {
			imsg := msg.(*types.MsgCreateIssuer)

			// TODO: pass in the did URI as an arg {msg.Id}
			// TODO: ensure this keeper can only read from store
			did, found := cicd.ik.GetIdentifier(ctx, []byte("did:cash:"+imsg.Owner))
			if !found {
				return ctx, sdkerrors.Wrapf(
					types.ErrIssuerFound,
					"identifer does not exists",
				)
			}

			// TODO: optimise here
			foundKey := false
			for _, auth := range did.Authentication {
				if auth.Controller == imsg.Owner {
					fmt.Println("found key")
					foundKey = true
				}
			}
			if !foundKey {
				return ctx, sdkerrors.Wrapf(
					types.ErrIssuerFound,
					"msg sender not in auth array in did document",
				)
			}

			// TODO: optimise here
			// check if the did document has the issuer credential
			hasIssuerCredential := false
			for _, service := range did.Services {
				// TODO use enum here
				if service.Type == "KYCCredential" {
					// TODO: ensure this keeper can only read from store
					vc, found := cicd.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
					if !found {
						return ctx, sdkerrors.Wrapf(
							types.ErrIssuerFound,
							"credential not found",
						)
					}
					hasIssuerCredential = vc.CredentialSubject.HasKyc
					// TODO: validate credential here
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
