package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
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

// DeleteVerifiableCredential deletes a verifiable credential
func (k msgServer) DeleteVerifiableCredential(
	goCtx context.Context,
	msg *types.MsgDeleteVerifiableCredential,
) (*types.MsgDeleteVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: there is something missing here

	if err := k.Keeper.DeleteVerifiableCredentialFromStore(ctx, []byte(msg.VerifiableCredentialId), msg.Owner); err != nil {
		return nil, sdkerrors.Wrapf(
			err, "verifiable credential validation failed",
		)
	}

	ctx.EventManager().EmitEvent(
		types.NewCredentialDeletedEvent(msg.Owner, msg.VerifiableCredentialId),
	)

	return &types.MsgDeleteVerifiableCredentialResponse{}, nil
}
