package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
	vcstypes "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
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

	identifier, _ := types.NewIdentifier(msg.Id, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewIdentifierCreatedEvent(msg.Id),
	)

	return &types.MsgCreateIdentifierResponse{}, nil
}

// AddAuthentication adds a public key and controller to an existing DID document
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

	// Only the first public key can add new public keys that control the DID document
	if identifier.Authentication[0].Controller != msg.Owner {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorized: AddAuthentication",
		)
	}

	// TODO: handle duplicates in the authentication slice
	msg.Authentication.Id = msg.Id + "#keys-" + fmt.Sprintf("%d", len(identifier.Authentication)+1)
	identifier.Authentication = append(identifier.Authentication, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewAuthenticationAddedEvent(msg.Id, msg.Authentication.Controller),
	)

	return &types.MsgAddAuthenticationResponse{}, nil
}

// AddService adds a service to an existing DID document
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

	for _, service := range identifier.Services {
		if service.Id == msg.ServiceData.Id {
			return nil, sdkerrors.Wrapf(
				types.ErrIdentifierNotFound,
				"service already exists: AddService",
			)
		}
	}

	if !vcstypes.IsValidCredentialType(msg.ServiceData.Type) {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"invalid service type: AddService",
		)
	}

	identifier.Services = append(identifier.Services, msg.ServiceData)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewServiceAddedEvent(msg.Id, msg.ServiceData.Id),
	)

	return &types.MsgAddServiceResponse{}, nil
}

// DeleteAuthentication removes a public key and controller from an existing DID document
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

	// Only the first public key can remove public keys that control the DID document
	if identifier.Authentication[0].Controller != msg.Owner {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorized: DeleteAuthentication",
		)
	}

	pubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.Key)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"pubkey not correct: DeleteAuthentication",
		)
	}
	address := sdk.AccAddress(pubKey.Address())

	// TODO: don't delete if only one auth
	auth := identifier.Authentication
	for i, key := range identifier.Authentication {
		if key.Controller == address.String() {
			// TODO: improve this logic
			// TODO: reorder auth ids as deleting and adding keys can lead to duplicated ids
			auth = append(
				identifier.Authentication[:i],
				identifier.Authentication[i+1:]...,
			)
		}
	}
	identifier.Authentication = auth

	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewAuthenticationDeletedEvent(msg.Id, address.String()),
	)

	return &types.MsgDeleteAuthenticationResponse{}, nil
}

// DeleteService removes a service from an existing DID document
func (k msgServer) DeleteService(
	goCtx context.Context,
	msg *types.MsgDeleteService,
) (*types.MsgDeleteServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"identifier not found: DeleteService",
		)
	}

	// Only the first public key can remove services from the DID document
	if identifier.Authentication[0].Controller != msg.Owner {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorized: DeleteService",
		)
	}

	// Only try to remove service if there are services
	if len(identifier.Services) == 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"no services found: DeleteService",
		)
	}

	services := identifier.Services
	for i, key := range identifier.Services {
		if key.Id == msg.ServiceId {
			// TODO: improve this logic
			services = append(
				identifier.Services[:i],
				identifier.Services[i+1:]...,
			)
		}
	}
	identifier.Services = services

	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewServiceDeletedEvent(msg.Id, msg.Id),
	)

	return &types.MsgDeleteServiceResponse{}, nil
}
