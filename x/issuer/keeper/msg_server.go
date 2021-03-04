package keeper

import (
	"context"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	// TODO: is state needed. should be handler in DID
	identifer := types.Issuer{
		Token:   msg.Token,
		Fee:     msg.Fee,
		State:   "CREATED",
		Address: msg.Owner,
	}

	k.Keeper.SetIssuer(ctx, []byte(identifer.Address), identifer)

	return &types.MsgCreateIssuerResponse{}, nil
}
