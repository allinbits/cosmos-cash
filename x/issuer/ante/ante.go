package ante

import (

	//	"encoding/base64"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/ante"

	//bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"

	//	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/issuer/keeper"

	//	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential/keeper"
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
	// TODO: ensure this keepers can only read from store
	// TODO: improve logic here
	//	for _, msg := range tx.GetMsgs() {
	//		// FIXME: this will fail with the next cosmos-sdk update as Type() gone from interface
	//		// use protoMessage type e.g bank/v1beta1/send
	//		if msg.String() == "send" {
	//			imsg := msg.(*bank.MsgSend)
	//			// FIXME: iterate over tokens and check the multi-send
	//			issuer, found := cicd.issuerk.GetIssuerByToken(ctx, []byte(imsg.Amount[0].Denom))
	//
	//			if found {
	//				err := cicd.validateKYCCredential(ctx, imsg.ToAddress, issuer.Address)
	//				if err != nil {
	//					return ctx, err
	//				}
	//
	//				err = cicd.validateKYCCredential(ctx, imsg.FromAddress, issuer.Address)
	//				if err != nil {
	//					return ctx, err
	//				}
	//			}
	//		}
	//	}

	return next(ctx, tx, simulate)
}

// validateKYCCredential validates a users KYC credential when they try to send a token
// to another user, this is called on every bank send message
//func (cicd CheckUserCredentialsDecorator) validateKYCCredential(
//	ctx sdk.Context,
//	address string,
//	issuerAddress string,
//) error {
//	issuerDID := didtypes.DID(ctx.ChainID(), address)
//
//	// TODO: tidy this functionality into the keeper,
//	// GetDidDocumentWithCondition, GetDidDocumentWithService, GetDidDocumentWithAuth
//	did, found := cicd.ik.GetDidDocument(ctx, []byte(issuerDID))
//	if !found {
//		return sdkerrors.Wrapf(
//			types.ErrDidDocumentDoesNotExist,
//			"did does not exists when validating the KYC credential",
//		)
//	}
//
//	if !did.HasRelationship(issuerDID, didtypes.Authentication) {
//		return sdkerrors.Wrapf(
//			types.ErrIncorrectControllerOfDidDocument,
//			"msg sender not in auth slice in did document when validating the KYC credential",
//		)
//	}
//
//	// check if the did document has the issuer credential
//	hasUserCredential := false
//	for _, service := range did.Service {
//		// TODO use enum here
//		if service.Type != "KYCCredential" {
//			continue
//		}
//
//		// TODO: ensure this keeper can only read from store
//		vc, found := cicd.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
//		if !found {
//			continue
//		}
//
//		userCred := vc.GetUserCred()
//		if userCred.Id != issuerDID {
//			continue
//		}
//
//		if !userCred.IsVerified {
//			continue
//		}
//
//		if vc.Issuer != issuerAddress {
//			continue
//		}
//
//		address, err := sdk.AccAddressFromBech32(issuerAddress)
//		if err != nil {
//			continue
//		}
//
//		account := cicd.accountk.GetAccount(ctx, address)
//		pubkey := account.GetPubKey()
//
//		hasUserCredential = vc.Validate(pubkey)
//
//		break
//	}
//	if !hasUserCredential {
//		return sdkerrors.Wrapf(
//			types.ErrIncorrectUserCredential,
//			"did document does not have a KYC credential to send e-money tokens",
//		)
//	}
//
//	return nil
//}
