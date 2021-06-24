package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// identifier module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeIdentifierCreated     = "identifier_created"
	EventTypeAuthenticationAdded   = "authentication_added"
	EventTypeServiceAdded          = "service_added"
	EventTypeAuthenticationDeleted = "authentication_deleted"
	EventTypeServiceDeleted        = "service_deleted"

	AttributeKeyOwner      = "owner"
	AttributeKeyController = "authentication_controller"
	AttributeKeyServiceID  = "service_id"
)

// NewIdentifierCreatedEvent constructs a new identifier_created sdk.Event
func NewIdentifierCreatedEvent(owner string) sdk.Event {
	return sdk.NewEvent(
		EventTypeIdentifierCreated,
		sdk.NewAttribute(AttributeKeyOwner, owner),
	)
}

// NewAuthenticationAddedEvent constructs a new authentication_added sdk.Event
func NewAuthenticationAddedEvent(owner string, controller string) sdk.Event {
	return sdk.NewEvent(
		EventTypeAuthenticationAdded,
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

// NewAuthenticationDeletedEvent constructs a new authentication_deleted sdk.Event
func NewAuthenticationDeletedEvent(owner string, controller string) sdk.Event {
	return sdk.NewEvent(
		EventTypeAuthenticationDeleted,
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
