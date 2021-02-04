package keeper

import (
	"context"
	"fmt"
	"github.com/allinbits/cosmos-cash/x/identifier/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// CreateIdentifier creates a new DID document
func (k msgServer) CreateIdentifier(
	goCtx context.Context,
	msg *types.MsgCreateIdentifier,
) (*types.MsgCreateIdentifierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierFound, "identifier already exists")

	}

	identifer, _ := types.NewIdentifier(msg.Id, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifer)

	return &types.MsgCreateIdentifierResponse{}, nil
}

// AddAuthentication adds a public key nad controller to am existing DID document
func (k msgServer) AddAuthentication(
	goCtx context.Context,
	msg *types.MsgAddAuthentication,
) (*types.MsgAddAuthenticationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "identifier not found")
	}

	// TODO: handle duplicates in the authentication slice
	msg.Authentication.Id = msg.Id + "#keys-" + fmt.Sprintf("%d", len(identifier.Authentication)+1)
	identifier.Authentication = append(identifier.Authentication, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	return &types.MsgAddAuthenticationResponse{}, nil
}
