package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// identifier module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeIdentifierCreated                = "did_document_created"
	EventTypeIdentifierUpdated                = "did_document_updated"
	EventTypeVerificationMethodAdded          = "verification_method_added"
	EventTypeVerificationRevoked              = "verification_method_revoked"
	EventTypeVerificationRelationshipsUpdated = "verification_relationships_updated"
	EventTypeServiceAdded                     = "service_added"
	EventTypeServiceDeleted                   = "service_deleted"

	AttributeDID           = "did"
	AttributeKeyOwner      = "owner"
	AttributeKeyController = "verification_method_controller"
	AttributeKeyServiceID  = "service_id"
)

// NewIdentifierCreatedEvent constructs a new identifier_created sdk.Event
func NewIdentifierCreatedEvent(owner string) sdk.Event {
	return sdk.NewEvent(
		EventTypeIdentifierCreated,
		sdk.NewAttribute(AttributeKeyOwner, owner),
	)
}

// NewIdentifierUpdatedEvent constructs a new identifier_created sdk.Event
// XXX: does it make sense ? cc @paddy
func NewIdentifierUpdatedEvent(did string, controllers ...string) sdk.Event {
	e := sdk.NewEvent(
		EventTypeIdentifierUpdated,
		sdk.NewAttribute(AttributeDID, did),
	)
	for _, c := range controllers {
		e.AppendAttributes(sdk.NewAttribute(AttributeKeyOwner, c))
	}
	return e
}

// NewVerificationAddedEvent constructs a new authentication_added sdk.Event
func NewVerificationAddedEvent(owner string, controller string) sdk.Event {
	return sdk.NewEvent(
		EventTypeVerificationMethodAdded,
		sdk.NewAttribute(AttributeKeyOwner, owner),
		sdk.NewAttribute(AttributeKeyController, controller),
	)
}

// NewServiceAddedEvent constructs a new service_added sdk.Event
func NewServiceAddedEvent(owner string, serviceID string) sdk.Event {
	return sdk.NewEvent(
		EventTypeServiceAdded,
		sdk.NewAttribute(AttributeKeyOwner, owner),
		sdk.NewAttribute(AttributeKeyServiceID, serviceID),
	)
}

// NewVerificationRevokedEvent constructs a new authentication_deleted sdk.Event
func NewVerificationRevokedEvent(owner string, controller string) sdk.Event {
	return sdk.NewEvent(
		EventTypeVerificationRevoked,
		sdk.NewAttribute(AttributeKeyOwner, owner),
		sdk.NewAttribute(AttributeKeyController, controller),
	)
}

// NewServiceDeletedEvent constructs a new service_deleted sdk.Event
func NewServiceDeletedEvent(owner string, serviceID string) sdk.Event {
	return sdk.NewEvent(
		EventTypeServiceDeleted,
		sdk.NewAttribute(AttributeKeyOwner, owner),
		sdk.NewAttribute(AttributeKeyServiceID, serviceID),
	)
}

// NewVerificationRelationshipsUpdatedEvent constructs a new relationships updated sdk.Event
func NewVerificationRelationshipsUpdatedEvent(owner string, methodID string) sdk.Event {
	return sdk.NewEvent(
		EventTypeVerificationRelationshipsUpdated,
		sdk.NewAttribute(AttributeKeyOwner, owner),
		sdk.NewAttribute(AttributeKeyServiceID, methodID),
	)
}
