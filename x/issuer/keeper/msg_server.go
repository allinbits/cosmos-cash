package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

	_, found := k.Keeper.GetIssuer(ctx, []byte(msg.Owner))
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

	// TODO: should the did URI be the issuer address
	issuer := types.Issuer{
		Token:   msg.Token,
		Fee:     msg.Fee,
		Address: msg.Owner,
	}

	k.Keeper.SetIssuer(ctx, issuer)

	// TODO: remove the next 24 lines in favor of minting tokens using the mint command
	circulatingSupply := 1000000000000

	issuerToken := sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(circulatingSupply)))

	if err := k.bk.MintCoins(
		ctx, types.ModuleName, issuerToken,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrMintingTokens,
			"cannot mint coins",
		)
	}

	recipient, _ := sdk.AccAddressFromBech32(msg.Owner)

	if err := k.bk.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, recipient, issuerToken,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrMintingTokens,
			"cannot send tokens from module to issuer account",
		)
	}

	ctx.EventManager().EmitEvent(
		types.NewIssuerCreatedEvent(msg.Owner, msg.Token, fmt.Sprint(circulatingSupply)),
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
	sender, _ := sdk.AccAddressFromBech32(issuer.Address)

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
	recipient, _ := sdk.AccAddressFromBech32(issuer.Address)

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
