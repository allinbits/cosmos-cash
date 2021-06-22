package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// issuer module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeIssuerCreated = "issuer_created"
	EventTypeTokenMinted   = "issuer_minted_token"
	EventTypeTokenBurned   = "issuer_burned_token"

	AttributeKeyIssuer = "issuer"
	AttributeKeyDenom  = "denom"
	AttributeKeyAmount = "amount"
)

// NewIssuerCreatedEvent constructs a new issuer_created sdk.Event
func NewIssuerCreatedEvent(issuer string, denom string, amount string) sdk.Event {
	return sdk.NewEvent(
		EventTypeIssuerCreated,
		sdk.NewAttribute(AttributeKeyIssuer, issuer),
		sdk.NewAttribute(AttributeKeyDenom, denom),
		sdk.NewAttribute(AttributeKeyAmount, amount),
	)
}

// NewTokenMintedvent constructs a new token_minted sdk.Event
func NewTokenMintedEvent(issuer string, denom string, amount string) sdk.Event {
	return sdk.NewEvent(
		EventTypeTokenMinted,
		sdk.NewAttribute(AttributeKeyIssuer, issuer),
		sdk.NewAttribute(AttributeKeyDenom, denom),
		sdk.NewAttribute(AttributeKeyAmount, amount),
	)
}

// NewTokenBurnedEvent constructs a new token_burned sdk.Event
func NewTokenBurnedEvent(issuer string, denom string, amount string) sdk.Event {
	return sdk.NewEvent(
		EventTypeTokenBurned,
		sdk.NewAttribute(AttributeKeyIssuer, issuer),
		sdk.NewAttribute(AttributeKeyDenom, denom),
		sdk.NewAttribute(AttributeKeyAmount, amount),
	)
}
