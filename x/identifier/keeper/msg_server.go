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

// CreateIdentifier creates a new DID document
func (k msgServer) CreateIdentifier(
	goCtx context.Context,
	msg *types.MsgCreateIdentifier,
) (*types.MsgCreateIdentifierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// setup a new did document (performs input validation)
	identifier, err := types.NewIdentifier(msg.Id,
		types.WithServices(msg.Services...),
		types.WithVerifications(msg.Verifications...),
	)
	if err != nil {
		return nil, err
	}
	// check that the identifier is not already taken
	_, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierFound, "a document with did %s already exists", msg.Id)
	}

	// persist the did document
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifier)
	// emit the event
	ctx.EventManager().EmitEvent(
		types.NewIdentifierCreatedEvent(msg.Id),
	)
	return &types.MsgCreateIdentifierResponse{}, nil
}

// UpdateIdentifier update an existing DID document
func (k msgServer) UpdateIdentifier(goCtx context.Context, msg *types.MsgUpdateIdentifier) (*types.MsgUpdateIdentifierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// XXX: does it make sense to do it twice to avoid opening the db
	// validate input before accessing the database
	if !types.IsValidDID(msg.Id) {
		return nil, sdkerrors.Wrap(types.ErrInvalidDIDFormat, "invalid identifier")
	}

	// get the did document
	didDoc, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "did document at %s not found", msg.Id)
	}
	// compute the signer did
	signerDID := types.DID(msg.Signer)

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(signerDID, types.RelationshipAuthentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer did %s not authorized to update the target did document at %s",
			signerDID, msg.Id,
		)
	}
	// set the controllers
	err := didDoc.SetControllers(msg.Controller...)
	if err != nil {
		return nil, err
	}
	// write the identifier
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), didDoc)

	// XXX: @paddy, wdyt? rubbish?
	ctx.EventManager().EmitEvent(
		types.NewIdentifierUpdatedEvent(msg.Id, msg.Controller...),
	)
	return &types.MsgUpdateIdentifierResponse{}, nil
}

// AddVerification adds a verification method and it's relationships to a DID Document
func (k msgServer) AddVerification(
	goCtx context.Context,
	msg *types.MsgAddVerification,
) (*types.MsgAddVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// XXX: does it make sense to do it twice to avoid opening the db
	// validate input before accessing the database
	if !types.IsValidDID(msg.Id) {
		return nil, sdkerrors.Wrap(types.ErrInvalidDIDFormat, "invalid identifier")
	}
	if err := types.ValidateVerification(msg.Verification); err != nil {
		return nil, err
	}

	// get the did document
	didDoc, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "did document at %s not found", msg.Id)
	}
	// compute the signer did
	signerDID := types.DID(msg.Signer)

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(signerDID, types.RelationshipAuthentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer did %s not authorized to add verification methods to the target did document at %s",
			signerDID, msg.Id,
		)
	}

	// add verifications (perform additional checks)
	if err := didDoc.AddVerifications(msg.Verification); err != nil {
		return nil, err
	}

	// write the identifier
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), didDoc)

	// TODO: may be useful to emit also the methodId
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

	// XXX: does it make sense to do it twice to avoid opening the db
	// validate input before accessing the database
	if !types.IsValidDID(msg.Id) {
		return nil, sdkerrors.Wrap(types.ErrInvalidDIDFormat, "invalid identifier")
	}
	if err := types.ValidateService(msg.ServiceData); err != nil {
		return nil, err
	}

	didDoc, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "did document at %s not found", msg.Id)
	}
	// add the service to the document
	err := didDoc.AddServices(msg.ServiceData)
	if err != nil {
		return nil, err
	}
	// verify that the service type is of type credential
	if !vcstypes.IsValidCredentialType(msg.ServiceData.Type) {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidInput,
			"invalid service type %s", msg.ServiceData.Type,
		)
	}
	// write to storage
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), didDoc)
	// emit events
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
	didDoc, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "did document at %s not found", msg.Id)
	}
	// compute the did based on the signer address
	signerDID := types.DID(msg.Signer)
	// any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(signerDID, types.RelationshipAuthentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer did %s not authorized to revoke verification methods from the target did document at %s",
			signerDID, msg.Id,
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

	// revoke the verification method + relationships
	didDoc.RevokeVerification(msg.MethodId)

	// persist to storage
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), didDoc)

	// emit event
	ctx.EventManager().EmitEvent(
		types.NewVerificationRevokedEvent(msg.Id, address.String()),
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
	didDoc, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "did document at %s not found", msg.Id)
	}
	// compute the did based on the signer address
	signerDID := types.DID(msg.Signer)

	// any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(signerDID, types.RelationshipAuthentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer did %s not authorized to delete services from the target did document at %s",
			signerDID, msg.Id,
		)
	}
	// Only try to remove service if there are services
	if len(didDoc.Services) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have services associated")
	}
	// delete the service
	didDoc.DeleteService(msg.ServiceId)

	// persist the did document
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), didDoc)

	// emit the event
	ctx.EventManager().EmitEvent(
		types.NewServiceDeletedEvent(msg.Id, msg.ServiceId),
	)

	return &types.MsgDeleteServiceResponse{}, nil
}

// SetVerificationRelationships set the verification relationships for an existing DID document
func (k msgServer) SetVerificationRelationships(goCtx context.Context, msg *types.MsgSetVerificationRelationships) (*types.MsgSetVerificationRelationshipsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve the did document
	didDoc, found := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrIdentifierNotFound, "did document at %s not found", msg.Id)
	}
	// compute the did based on the signer address
	signerDID := types.DID(msg.Signer)

	// any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(signerDID, types.RelationshipAuthentication) {
		return nil, sdkerrors.Wrapf(
			types.ErrUnauthorized,
			"signer did %s not authorized to set verification relationships on the target did document at %s",
			signerDID, msg.Id,
		)
	}

	// set the verification relationships
	err := didDoc.SetVerificationRelationships(msg.MethodId, msg.Relationships...)
	if err != nil {
		return nil, err
	}

	// persist the did document
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), didDoc)

	// emit the event
	// TODO: rubbish
	ctx.EventManager().EmitEvent(
		types.NewVerificationRelationshipsUpdatedEvent(msg.Id, msg.MethodId),
	)

	return &types.MsgSetVerificationRelationshipsResponse{}, nil
}
