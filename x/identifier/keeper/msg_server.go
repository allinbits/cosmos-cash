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
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierFound,
			"identifier already exists",
		)

	}

	identifer, _ := types.NewIdentifier(msg.Id, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifer)

	return &types.MsgCreateIdentifierResponse{}, nil
}

// AddAuthentication adds a public key and controller to am existing DID document
func (k msgServer) AddAuthentication(
	goCtx context.Context,
	msg *types.MsgAddAuthentication,
) (*types.MsgAddAuthenticationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"identifier not found: AddAuthentication",
		)
	}

	// Only the first public key can add new public keys that controls the did document
	if identifier.Authentication[0].Controller != msg.Owner {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorised",
		)
	}

	// TODO: handle duplicates in the authentication slice
	msg.Authentication.Id = msg.Id + "#keys-" + fmt.Sprintf("%d", len(identifier.Authentication)+1)
	identifier.Authentication = append(identifier.Authentication, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	return &types.MsgAddAuthenticationResponse{}, nil
}

// AddService adds a serivce to am existing DID document
func (k msgServer) AddService(
	goCtx context.Context,
	msg *types.MsgAddService,
) (*types.MsgAddServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"identifier not found: AddService",
		)
	}

	identifier.Services = append(identifier.Services, msg.ServiceData)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	return &types.MsgAddServiceResponse{}, nil
}

// AddAuthentication adds a public key and controller to am existing DID document
func (k msgServer) DeleteAuthentication(
	goCtx context.Context,
	msg *types.MsgDeleteAuthentication,
) (*types.MsgDeleteAuthenticationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"identifier not found: DeleteAuthentication",
		)
	}

	// Only the first public key can add new public keys that controls the did document
	if identifier.Authentication[0].Controller != msg.Owner {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorised",
		)
	}

	pubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.Key)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"pubkey not correct",
		)
	}
	address := sdk.AccAddress(pubKey.Address())

	// TODO: don't delete if only one auth
	auth := identifier.Authentication
	for i, key := range identifier.Authentication {
		if key.Controller == address.String() {
			auth = append(
				identifier.Authentication[:i],
				identifier.Authentication[i+1:]...,
			)
		}
	}
	identifier.Authentication = auth

	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	return &types.MsgDeleteAuthenticationResponse{}, nil
}
