package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
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

	// Check to see if the provided did is in the store
	did, found := k.Keeper.didKeeper.GetDidDocument(ctx, []byte(msg.IssuerDid))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrDidDocumentDoesNotExist,
			"did does not exists",
		)
	}

	// Check to see if the msg signer has a verification relationship in the did document
	if !did.HasRelationship(msg.Owner, didtypes.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrIncorrectControllerOfDidDocument,
			"msg sender not in auth array in did document",
		)
	}

	// Check to see if the provided verifiable credential is in the store
	vc, found := k.Keeper.vcKeeper.GetVerifiableCredential(ctx, []byte(msg.LicenseCredId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"verifiable credential not found",
		)
	}

	// Validate the credential subject ID the same as the provided did document
	issuerCred := vc.GetLicenseCred()
	if issuerCred.Id != did.Id {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"issuer id not correct",
		)
	}

	// TODO: validate credential was issued by a regulator

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
	}

	k.Keeper.SetIssuer(ctx, issuer)

	ctx.EventManager().EmitEvent(
		types.NewIssuerCreatedEvent(msg.Owner, msg.Token),
	)

	return &types.MsgCreateIssuerResponse{}, nil
}

// CreateIssuer creates a new e-money token issuer
func (k msgServer) BurnToken(
	goCtx context.Context,
	msg *types.MsgBurnToken,
) (*types.MsgBurnTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	issuer, found := k.Keeper.GetIssuer(ctx, []byte(msg.Owner))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerNotFound,
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

func (k msgServer) MintToken(
	goCtx context.Context,
	msg *types.MsgMintToken,
) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	issuer, found := k.Keeper.GetIssuer(ctx, []byte(msg.Owner))
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
	}

	// the recipient is the issuer itself
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
