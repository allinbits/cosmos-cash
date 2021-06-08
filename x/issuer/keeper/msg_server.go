package keeper

import (
	"context"

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

	// TODO: should the did URI be the issuer address
	identifier := types.Issuer{
		Token:   msg.Token,
		Fee:     msg.Fee,
		Address: msg.Owner,
	}

	k.Keeper.SetIssuer(ctx, identifier)

	// TODO: this needs to be refactored
	circulatingSupply := 1000000000000

	issuerToken := sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(circulatingSupply)))

	// mint new tokens for the issuer
	if err := k.bk.MintCoins(
		ctx, types.ModuleName, issuerToken,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"cannot mint coins",
		)
	}

	receipent, _ := sdk.AccAddressFromBech32(msg.Owner)

	// send tokens from module to issuer
	if err := k.bk.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receipent, issuerToken,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"cannot send tokens from module to issuer account",
		)
	}

	return &types.MsgCreateIssuerResponse{}, nil
}
