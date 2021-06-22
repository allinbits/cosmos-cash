package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// credential module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeCredentialCreated = "credential_created"

	AttributeKeyOwner        = "owner"
	AttributeKeyCredentialID = "credential_id"
)

// NewCredentialCreatedEvent constructs a new credential_created sdk.Event
func NewCredentialCreatedEvent(owner string, credentialID string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCredentialCreated,
		sdk.NewAttribute(AttributeKeyOwner, owner),
		sdk.NewAttribute(AttributeKeyCredentialID, credentialID),
	)
}
