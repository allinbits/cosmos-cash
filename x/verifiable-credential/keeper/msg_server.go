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
	vc, found := k.Keeper.GetVerifiableCredential(ctx, []byte(msg.VerifiableCredentialId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrVerifiableCredentialNotFound,
			"error deleting credential; credential not found",
		)
	}
	// TODO: the validate proof also accepts validation methods that are not authentication
	if err := ValidateProof(ctx, k.Keeper, vc); err != nil {
		return nil, sdkerrors.Wrapf(
			err, "verifiable credential validation failed",
		)
	}

	if err := k.Keeper.DeleteVerifiableCredentialFromStore(ctx, []byte(vc.Id)); err != nil {
		return nil, sdkerrors.Wrapf(
			err, "verifiable credential validation failed",
		)
	}

	ctx.EventManager().EmitEvent(
		types.NewCredentialDeletedEvent(msg.Owner, msg.VerifiableCredentialId),
	)

	return &types.MsgDeleteVerifiableCredentialResponse{}, nil
}
