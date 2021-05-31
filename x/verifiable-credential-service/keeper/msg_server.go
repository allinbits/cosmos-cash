package keeper

import (
	"context"
	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
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

// CreateVerifiableCredential creates a new verifiable credential
func (k msgServer) CreateVerifiableCredential(
	goCtx context.Context,
	msg *types.MsgCreateVerifiableCredential,
) (*types.MsgCreateVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: only issuers can create user verifiable creds
	_, found := k.Keeper.GetVerifiableCredential(ctx, []byte(msg.VerifiableCredential.Id))
	if found {
		return nil, sdkerrors.Wrapf(
			types.ErrVerifiableCredentialFound,
			"vc already exists",
		)

	}

	k.Keeper.SetVerifiableCredential(
		ctx,
		[]byte(msg.VerifiableCredential.Id),
		*msg.VerifiableCredential,
	)

	return &types.MsgCreateVerifiableCredentialResponse{}, nil
}
