package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
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

// RevokeCredential revoke a credential
func (k msgServer) RevokeCredential(goCtx context.Context, msg *types.MsgRevokeCredential) (*types.MsgRevokeCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Info("revoke credential request", "credential", msg.CredentialId, "address", msg.Owner)

	if vcErr := k.DeleteVerifiableCredentialFromStore(ctx, []byte(msg.CredentialId), msg.Owner); vcErr != nil {
		err := sdkerrors.Wrapf(vcErr, "credential proof cannot be verified")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("revoke license request successful", "credential", msg.CredentialId, "address", msg.Owner)

	ctx.EventManager().EmitEvent(
		types.NewCredentialDeletedEvent(msg.Owner, msg.CredentialId),
	)

	return &types.MsgRevokeCredentialResponse{}, nil
}
