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

// CheckIssuerCredentialsDecorator checks the issuer has a EMILicense in a preprocessing hook
type CheckIssuerCredentialsDecorator struct {
	issuerk keeper.Keeper
	ik      didkeeper.Keeper
	vcsk    vcskeeper.Keeper
}

func NewCheckIssuerCredentialsDecorator(
	issuerk keeper.Keeper,
	ik didkeeper.Keeper,
	vcsk vcskeeper.Keeper,
) CheckIssuerCredentialsDecorator {
	return CheckIssuerCredentialsDecorator{
		issuerk: issuerk,
		ik:      ik,
		vcsk:    vcsk,
	}
}

// AnteHandle check issuer credentials
// FIXME: fix this file
func (cicd CheckIssuerCredentialsDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	// TODO: improve logic here
	//	for _, msg := range tx.GetMsgs() {
	//		if msg.String() == "create-issuer" {
	//			imsg := msg.(*types.MsgCreateIssuer)
	//
	//			signerDID := didtypes.DID(ctx.ChainID(), imsg.Owner)
	//
	//			// TODO: pass in the did URI as an arg {msg.Id}
	//			// TODO: ensure this keeper can only read from store
	//			did, found := cicd.ik.GetDidDocument(ctx, []byte(signerDID))
	//			if !found {
	//				return ctx, sdkerrors.Wrapf(
	//					types.ErrDidDocumentDoesNotExist,
	//					"did does not exists",
	//				)
	//			}
	//
	//			// verification authorization
	//			if !did.HasRelationship(signerDID, didtypes.Authentication) {
	//				return ctx, sdkerrors.Wrapf(
	//					types.ErrIncorrectControllerOfDidDocument,
	//					"msg sender not in auth array in did document",
	//				)
	//			}
	//
	//			// check if the did document has the issuer credential
	//			hasIssuerCredential := false
	//			for _, service := range did.Services {
	//				// TODO use enum here
	//				if service.Type == "IssuerCredential" {
	//					// TODO: ensure this keeper can only read from store
	//					vc, found := cicd.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
	//					if !found {
	//						return ctx, sdkerrors.Wrapf(
	//							types.ErrIssuerFound,
	//							"verifiable credential not found",
	//						)
	//					}
	//
	//					issuerCred := vc.GetUserCred()
	//					if issuerCred.Id != signerDID {
	//						return ctx, sdkerrors.Wrapf(
	//							types.ErrIssuerFound,
	//							"issuer id not correct",
	//						)
	//					}
	//
	//					//	if issuerCred.IsVerified == false {
	//					//		return ctx, sdkerrors.Wrapf(
	//					//			types.ErrIssuerFound,
	//					//			"issuer is not verified",
	//					//		)
	//					//	}
	//
	//					hasIssuerCredential = true
	//					// TODO: validate credential here has been issued by regulator
	//				}
	//			}
	//			if !hasIssuerCredential {
	//				return ctx, sdkerrors.Wrapf(
	//					types.ErrIssuerFound,
	//					"did document doesnt have a credential to create issuers",
	//				)
	//			}
	//		}
	//	}
	return next(ctx, tx, simulate)
}

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
//		s, err := base64.StdEncoding.DecodeString(vc.Proof.Signature)
//		if err != nil {
//			continue
//		}
//		emptyProof := vcstypes.NewProof("", "", "", "", "")
//		vc.Proof = &emptyProof
//
//		// TODO: this is an expesive operation, could lead to DDOS
//		// TODO: we can hash this and make this less expensive
//		hasUserCredential = pubkey.VerifySignature(
//			vc.GetBytes(),
//			s,
//		)
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
