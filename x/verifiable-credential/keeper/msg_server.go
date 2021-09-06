package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
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

// CreateVerifiableCredential creates a new verifiable credential
func (k msgServer) CreateVerifiableCredential(
	goCtx context.Context,
	msg *types.MsgCreateVerifiableCredential,
) (*types.MsgCreateVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	ctx.EventManager().EmitEvent(
		types.NewCredentialCreatedEvent(msg.Owner, msg.VerifiableCredential.Id),
	)

	return &types.MsgCreateVerifiableCredentialResponse{}, nil
}

// DeleteVerifiableCredential deletes a verifiable credential
func (k msgServer) DeleteVerifiableCredential(
	goCtx context.Context,
	msg *types.MsgDeleteVerifiableCredential,
) (*types.MsgDeleteVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	did, found := k.Keeper.didKeeper.GetDidDocument(ctx, []byte(msg.IssuerDid))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrDidDocumentDoesNotExist,
			"did does not exists",
		)
	}

	vc, found := k.Keeper.GetVerifiableCredential(ctx, []byte(msg.VerifiableCredentialId))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrVerifiableCredentialNotFound,
			"error deleting credential; credential not found",
		)
	}

	if vc.Issuer != did.Id {
		return nil, sdkerrors.Wrapf(
			types.ErrVerifiableCredentialIssuer,
			"provided vc and did issuer do not match",
		)
	}

	if !did.HasRelationship(msg.Owner, didtypes.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"signer is not in issuer did",
		)
	}

	k.Keeper.DeleteVerifiableCredentialFromStore(ctx, []byte(vc.Id))

	ctx.EventManager().EmitEvent(
		types.NewCredentialDeletedEvent(msg.Owner, msg.VerifiableCredentialId),
	)

	return &types.MsgDeleteVerifiableCredentialResponse{}, nil
}
