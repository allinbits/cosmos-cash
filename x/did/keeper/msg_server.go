package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/allinbits/cosmos-cash/x/did/types"
	vcstypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
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

// CreateDidDocument creates a new DID document
func (k msgServer) CreateDidDocument(
	goCtx context.Context,
	msg *types.MsgCreateDidDocument,
) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// setup a new did document (performs input validation)
	did, err := types.NewDidDocument(msg.Id,
		types.WithServices(msg.Services...),
		types.WithVerifications(msg.Verifications...),
	)
	if err != nil {
		return nil, err
	}
	// check that the did is not already taken
	_, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentFound, "a document with did %s already exists", msg.Id)
	}

	// persist the did document
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), did)

	k.Logger(ctx).Info("Created a DidDocument for", "did", msg.Id, "controller", msg.Signer)

	// emit the event
	ctx.EventManager().EmitEvent(
		types.NewDidDocumentCreatedEvent(msg.Id),
	)
	return &types.MsgCreateDidDocumentResponse{}, nil
}

// UpdateDidDocument update an existing DID document
func (k msgServer) UpdateDidDocument(
	goCtx context.Context,
	msg *types.MsgUpdateDidDocument,
) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the did document
	didDoc, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", msg.Id)
	}

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(msg.Signer, types.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer %s not authorized to update the target did document at %s",
			msg.Signer, msg.Id,
		)
	}
	// set the controllers
	err := didDoc.SetControllers(msg.Controller...)
	if err != nil {
		return nil, err
	}
	// write the did
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("Updated a DidDocument for", "did", msg.Id, "controller", msg.Signer)

	// NOTE: events are expected to change during client development
	ctx.EventManager().EmitEvent(
		types.NewDidDocumentUpdatedEvent(msg.Id, msg.Controller...),
	)
	return &types.MsgUpdateDidDocumentResponse{}, nil
}

// AddVerification adds a verification method and it's relationships to a DID Document
func (k msgServer) AddVerification(
	goCtx context.Context,
	msg *types.MsgAddVerification,
) (*types.MsgAddVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the did document
	didDoc, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", msg.Id)
	}

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(msg.Signer, types.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer account %s not authorized to add verification methods to the target did document at %s",
			msg.Signer, msg.Id,
		)
	}

	// add verifications (perform additional checks)
	if err := didDoc.AddVerifications(msg.Verification); err != nil {
		return nil, err
	}

	// write the did
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("Added a new verification method for", "did", msg.Id, "controller", msg.Signer)

	// NOTE: events are expected to change during client development
	ctx.EventManager().EmitEvent(
		types.NewVerificationAddedEvent(msg.Id, msg.Verification.Method.Controller),
	)
	return &types.MsgAddVerificationResponse{}, nil
}

// AddService adds a service to an existing DID document
func (k msgServer) AddService(
	goCtx context.Context,
	msg *types.MsgAddService,
) (*types.MsgAddServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// perform checks on the service
	if err := types.ValidateService(msg.ServiceData); err != nil {
		return nil, err
	}

	didDoc, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", msg.Id)
	}
	// verify that the service type is of type credential
	if !vcstypes.IsValidCredentialType(msg.ServiceData.Type) {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidInput,
			"invalid service type %s", msg.ServiceData.Type,
		)
	}
	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(msg.Signer, types.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer %s not authorized to add services to the target did document at %s",
			msg.Signer, msg.Id,
		)
	}
	// add the service to the document
	err := didDoc.AddServices(msg.ServiceData)
	if err != nil {
		return nil, err
	}
	// write to storage
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("Added a new service for", "did", msg.Id, "controller", msg.Signer)
	// NOTE: events are expected to change during client development
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
	// retrieve the did document
	didDoc, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", msg.Id)
	}
	// any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(msg.Signer, types.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer %s not authorized to revoke verification methods from the target did document at %s",
			msg.Signer, msg.Id,
		)
	}
	// revoke the verification method + relationships
	if err := didDoc.RevokeVerification(msg.MethodId); err != nil {
		return nil, err
	}

	// persist to storage
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("Revoked verification method from did document for", "did", msg.Id, "controller", msg.Signer)

	// emit event
	ctx.EventManager().EmitEvent(
		types.NewVerificationRevokedEvent(msg.Id, msg.Signer),
	)

	return &types.MsgRevokeVerificationResponse{}, nil
}

// DeleteService removes a service from an existing DID document
func (k msgServer) DeleteService(
	goCtx context.Context,
	msg *types.MsgDeleteService,
) (*types.MsgDeleteServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve the did document
	didDoc, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", msg.Id)
	}
	// any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(msg.Signer, types.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer %s not authorized to delete services from the target did document at %s",
			msg.Signer, msg.Id,
		)
	}
	// Only try to remove service if there are services
	if len(didDoc.Services) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have services associated")
	}
	// delete the service
	didDoc.DeleteService(msg.ServiceId)

	// persist the did document
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("Removed service from did document for", "did", msg.Id, "controller", msg.Signer)

	// NOTE: events are expected to change during client development
	ctx.EventManager().EmitEvent(
		types.NewServiceDeletedEvent(msg.Id, msg.ServiceId),
	)

	return &types.MsgDeleteServiceResponse{}, nil
}

// SetVerificationRelationships set the verification relationships for an existing DID document
func (k msgServer) SetVerificationRelationships(
	goCtx context.Context,
	msg *types.MsgSetVerificationRelationships,
) (*types.MsgSetVerificationRelationshipsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve the did document
	didDoc, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", msg.Id)
	}
	// any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(msg.Signer, types.Authentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer %s not authorized to set verification relationships on the target did document at %s",
			msg.Signer, msg.Id,
		)
	}

	// set the verification relationships
	err := didDoc.SetVerificationRelationships(msg.MethodId, msg.Relationships...)
	if err != nil {
		return nil, err
	}

	// persist the did document
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("Set verification relationship from did document for", "did", msg.Id, "controller", msg.Signer)

	// NOTE: events are expected to change during client development
	ctx.EventManager().EmitEvent(
		types.NewVerificationRelationshipsUpdatedEvent(msg.Id, msg.MethodId),
	)

	return &types.MsgSetVerificationRelationshipsResponse{}, nil
}
