package keeper

import (
	"context"

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

func (k msgServer) UpdateIdentifier(goCtx context.Context, msg *types.MsgUpdateIdentifier) (*types.MsgUpdateIdentifierResponse, error) {
	return nil, nil
}

func (k msgServer) SetVerificationRelationships(goCtx context.Context, msg *types.MsgSetVerificationRelationships) (*types.MsgSetVerificationRelationshipsResponse, error) {
	return nil, nil
}

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

	identifier, _ := types.NewIdentifier(msg.Id, msg.Services, msg.Verifications)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewIdentifierCreatedEvent(msg.Id),
	)

	return &types.MsgCreateIdentifierResponse{}, nil
}

// AddAuthentication adds a public key and controller to an existing DID document
func (k msgServer) AddVerification(
	goCtx context.Context,
	msg *types.MsgAddVerification,
) (*types.MsgAddVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get the did document
	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"identifier not found: AddAuthentication",
		)
	}

	// Any verification method in the authentication relationship can update the DID document
	if !identifier.ControllerInRelationships(msg.Owner, types.VerificationRelationship_authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorized: AddVerification",
		)
	}

	// TODO: handle duplicates in the authentication slice
	if err := identifier.AddVerifications(msg.Verification); err != nil {
		return nil, err
	}

	// msg.Verification.Method.Id = msg.Id + "#keys-" + fmt.Sprintf("%d", len(identifier.Authentication)+1)
	// identifier.VerificationMethods = append(identifier.VerificationMethods, msg.Authentication)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewAuthenticationAddedEvent(msg.Id, msg.Verification.Method.Controller),
	)

	return &types.MsgAddVerificationResponse{}, nil
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

// RevokeVerification removes a public key and controller from an existing DID document
func (k msgServer) RevokeVerification(
	goCtx context.Context,
	msg *types.MsgRevokeVerification,
) (*types.MsgRevokeVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifier, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"identifier not found: RevokeVerification",
		)
	}

	// Only the first public key can remove public keys that control the DID document
	if !identifier.ControllerInRelationships(msg.Owner, types.VerificationRelationship_authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"msg sender not authorized: RevokeVerification",
		)
	}

	// XXX: there is something wrong here
	pubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.MethodId)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIdentifierNotFound,
			"pubkey not correct: RevokeVerification",
		)
	}
	address := sdk.AccAddress(pubKey.Address())

	// TODO: don't delete if only one auth
	// if len(identifier.VerificationMethods) == 1 {

	// }

	identifier.RevokeVerification(msg.MethodId)

	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)

	ctx.EventManager().EmitEvent(
		types.NewAuthenticationDeletedEvent(msg.Id, address.String()),
	)

	return &types.MsgRevokeVerificationResponse{}, nil
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

	// TODO add check on the controller (did level)
	// Only the first public key can remove services from the DID document
	if !identifier.ControllerInRelationships(msg.Owner, types.VerificationRelationship_authentication) {
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
