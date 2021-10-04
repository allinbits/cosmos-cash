package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateIssuer creates a new e-money token issuer
func (k msgServer) CreateIssuer(
	goCtx context.Context,
	msg *types.MsgCreateIssuer,
) (*types.MsgCreateIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check to see if the provided verifiable credential is in the store
	vc, found := k.Keeper.vcKeeper.GetVerifiableCredential(ctx, []byte(msg.LicenseCredId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrLicenseCredentialNotFound,
			"verifiable credential not found",
		)
	}

	// Validate the provided issuer credential
	err := k.validateIssuerCredential(ctx, msg.IssuerDid, vc, msg.Owner)
	if err != nil {
		return nil, err
	}

	_, found = k.Keeper.GetIssuer(ctx, []byte(msg.IssuerDid))
	if found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"issuer already exists",
		)
	}

	_, found = k.Keeper.GetIssuerByToken(ctx, []byte(msg.Token))
	if found {
		return nil, sdkerrors.Wrapf(
			types.ErrTokenAlreadyExists,
			"token denom already exists",
		)
	}

	issuer := types.Issuer{
		Token:     msg.Token,
		Fee:       msg.Fee,
		IssuerDid: msg.IssuerDid,
		Paused:    false,
	}

	k.Keeper.SetIssuer(ctx, issuer)

	ctx.EventManager().EmitEvent(
		types.NewIssuerCreatedEvent(msg.Owner, msg.Token),
	)

	return &types.MsgCreateIssuerResponse{}, nil
}

// BurnToken burns a token for an e-money issuer
func (k msgServer) BurnToken(
	goCtx context.Context,
	msg *types.MsgBurnToken,
) (*types.MsgBurnTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check to see if the provided verifiable credential is in the store
	vc, found := k.Keeper.vcKeeper.GetVerifiableCredential(ctx, []byte(msg.LicenseCredId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrLicenseCredentialNotFound,
			"verifiable credential not found",
		)
	}

	// Validate the provided issuer credential
	err := k.validateIssuerCredential(ctx, msg.IssuerDid, vc, msg.Owner)
	if err != nil {
		return nil, err
	}

	issuer, found := k.Keeper.GetIssuer(ctx, []byte(msg.IssuerDid))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"issuer does not exists",
		)
	}

	// parse the token amount and verify that the amount requested is for the issuer token
	amounts, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdk.ErrInvalidDecimalStr,
			"coin string format not recognized",
		)
	}
	for _, a := range amounts {
		if a.GetDenom() != issuer.Token {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidIssuerDenom,
				"issuer can only burn tokens of %s (requested %s)",
				issuer.Token, a.GetDenom(),
			)
		}
	}

	// sender is the issuer
	sender, _ := sdk.AccAddressFromBech32(msg.Owner)

	if err := k.bk.SendCoinsFromAccountToModule(
		ctx, sender, types.ModuleName, amounts,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrBurningTokens,
			"cannot send tokens from issuer account to module",
		)
	}

	if err := k.bk.BurnCoins(
		ctx, types.ModuleName, amounts,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrBurningTokens,
			"cannot burn coins",
		)
	}

	ctx.EventManager().EmitEvent(
		types.NewTokenBurnedEvent(msg.Owner, issuer.Token, string(msg.Amount)),
	)

	return &types.MsgBurnTokenResponse{}, nil
}

// MintToken mints a token for an e-money issuer
func (k msgServer) MintToken(
	goCtx context.Context,
	msg *types.MsgMintToken,
) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check to see if the provided verifiable credential is in the store
	vc, found := k.Keeper.vcKeeper.GetVerifiableCredential(ctx, []byte(msg.LicenseCredId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrLicenseCredentialNotFound,
			"verifiable credential not found",
		)
	}

	// Validate the provided issuer credential
	err := k.validateIssuerCredential(ctx, msg.IssuerDid, vc, msg.Owner)
	if err != nil {
		return nil, err
	}

	issuer, found := k.Keeper.GetIssuer(ctx, []byte(msg.IssuerDid))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"issuer does not exists",
		)
	}

	// parse the token amount and verify that the amount requested is for the issuer token
	amounts, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdk.ErrInvalidDecimalStr,
			"coin string format not recognized",
		)
	}
	for _, a := range amounts {
		if a.GetDenom() != issuer.Token {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidIssuerDenom,
				"issuer can only issue tokens of %s (requested %s)",
				issuer.Token, a.GetDenom(),
			)
		}

		// Validate the minting amount is in range
		err := k.validateMintingAmount(ctx, vc, a)
		if err != nil {
			return nil, err
		}
	}

	// the recipient is the signer of the message
	recipient, _ := sdk.AccAddressFromBech32(msg.Owner)

	if err := k.bk.MintCoins(
		ctx, types.ModuleName, amounts,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrMintingTokens,
			"cannot mint coins",
		)
	}

	if err := k.bk.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, recipient, amounts,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrMintingTokens,
			"failed sending tokens from module to issuer",
		)
	}

	ctx.EventManager().EmitEvent(
		types.NewTokenMintedEvent(msg.Owner, issuer.Token, string(msg.Amount)),
	)

	return &types.MsgMintTokenResponse{}, nil
}

// PauseToken pauses a token for an issuer
func (k msgServer) PauseToken(
	goCtx context.Context,
	msg *types.MsgPauseToken,
) (*types.MsgPauseTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check to see if the provided verifiable credential is in the store
	vc, found := k.Keeper.vcKeeper.GetVerifiableCredential(ctx, []byte(msg.LicenseCredId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrLicenseCredentialNotFound,
			"verifiable credential not found",
		)
	}

	// Validate the provided issuer credential
	err := k.validateIssuerCredential(ctx, msg.IssuerDid, vc, msg.Owner)
	if err != nil {
		return nil, err
	}

	issuer, found := k.Keeper.GetIssuer(ctx, []byte(msg.IssuerDid))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"issuer does not exists",
		)
	}

	issuer.Paused = !issuer.Paused

	k.Keeper.SetIssuer(ctx, issuer)

	ctx.EventManager().EmitEvent(
		types.NewTokenPausedEvent(msg.Owner, issuer.Token),
	)

	return &types.MsgPauseTokenResponse{}, nil
}

// validateIssuerCredential validate the signer of the message is part of the issuer did and the provided credential
func (k msgServer) validateIssuerCredential(
	ctx sdk.Context,
	issuerDid string,
	licenseCred vctypes.VerifiableCredential,
	signer string,
) error {
	// Check to see if the provided did is in the store
	did, _, err := k.Keeper.didKeeper.ResolveDid(ctx, didtypes.DID(issuerDid))
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrDidDocumentDoesNotExist,
			"did does not exists",
		)
	}

	// Check to see if the msg signer has a verification relationship in the did document
	if !did.HasRelationship(didtypes.NewBlockchainAccountID(ctx.ChainID(), signer), didtypes.Authentication) {
		return sdkerrors.Wrapf(
			types.ErrIncorrectControllerOfDidDocument,
			"msg sender not in auth array in did document",
		)
	}

	// Validate the credential subject ID the same as the provided did document
	// TODO: dig deeper into the license type
	issuerCred := licenseCred.GetLicenseCred()
	if issuerCred.Id != did.Id {
		return sdkerrors.Wrapf(
			types.ErrIncorrectLicenseCredential,
			"issuer id not correct",
		)
	}

	return nil
}

// validateMintingAmount validates the amount being minted is correct
func (k msgServer) validateMintingAmount(
	ctx sdk.Context,
	licenseCred vctypes.VerifiableCredential,
	mintingAmount sdk.Coin,
) error {
	// Get the supply from the bank keeper
	supply := k.bk.GetSupply(ctx, mintingAmount.Denom)

	// Calculate the reserve by subtraction the supply from the issuer credential circulation limit
	reserve := licenseCred.GetLicenseCred().CirculationLimit.Amount.Sub(supply.Amount)

	// Validate the reserve is greater that the amount being minted
	if reserve.LT(mintingAmount.Amount) {
		return sdkerrors.Wrapf(
			types.ErrMintingTokens,
			"issuer cannot mint more than the circulation limit defined in their credential",
		)
	}

	return nil
}

// IssueUserCredential activates a regulator
func (k msgServer) IssueUserCredential(
	goCtx context.Context,
	msg *types.MsgIssueUserCredential,
) (*types.MsgIssueUserCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("issue user credential request", "credential", msg.Credential, "address", msg.Owner)

	// check that the issuer is a holder of LicenseCredential
	// TODO: need to go a bit deeper about the type of the license
	vcs := k.vcKeeper.GetVerifiableCredentialWithType(ctx, msg.Credential.GetIssuer(), vctypes.LicenseCredential)
	if len(vcs) != 1 { // there must be exactly one
		err := sdkerrors.Wrapf(types.ErrLicenseCredentialNotFound, "credential issuer is not a licensed e-money issuer")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// store the credentials
	if vcErr := k.vcKeeper.SetVerifiableCredential(ctx, []byte(msg.Credential.Id), *msg.Credential); vcErr != nil {
		err := sdkerrors.Wrapf(vcErr, "credential proof cannot be verified")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("issue user credential request successful", "credentialID", msg.Credential.Id)

	ctx.EventManager().EmitEvent(
		vctypes.NewCredentialCreatedEvent(msg.Owner, msg.Credential.Id),
	)

	return &types.MsgIssueUserCredentialResponse{}, nil
}
