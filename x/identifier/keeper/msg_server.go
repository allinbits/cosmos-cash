package keeper

import (
	"context"
	"github.com/allinbits/cosmos-cash/x/identifier/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateIdentifier(
	goCtx context.Context,
	msg *types.MsgCreateIdentifier,
) (*types.MsgCreateIdentifierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: check if identifier exists

	identifer, _ := types.NewIdentifier(msg.Id, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifer)

	return &types.MsgCreateIdentifierResponse{}, nil
}
