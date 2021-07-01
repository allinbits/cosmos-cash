package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// --------------------------
// CREATE IDENTIFIER
// --------------------------

// ValidateBasic performs a basic check of the MsgCreateIdentifier fields.
func (msg MsgCreateIdentifier) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if msg.Verifications == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verifications are required")
	}

	return nil

}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------

// ValidateBasic performs a basic check of the MsgUpdateIdentifier fields.
func (msg MsgUpdateIdentifier) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	// if controller is set must be compliant
	if !IsEmpty(msg.Controller) && !IsValidDID(msg.Controller) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, "controller validation error")
	}

	// TODO: this
	return nil
}

// --------------------------
// ADD VERIFICATION
// --------------------------

// ValidateBasic performs a basic check of the MsgAddVerification fields.
func (msg MsgAddVerification) ValidateBasic() error {
	if msg.Id == "" {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if msg.Verification == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verification data is required")
	}

	// TODO: add more verification stuff

	return nil

}

// --------------------------
// REVOKE VERIFICATION
// --------------------------

// ValidateBasic performs a basic check of the MsgRevokeVerification fields.
func (msg MsgRevokeVerification) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidDIDURL(msg.MethodId) {
		return sdkerrors.Wrap(ErrInvalidDIDURLFormat, "verification method id validation error")
	}
	return nil
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------

// ValidateBasic performs a basic check of the MsgSetVerificationRelationships fields.
func (msg MsgSetVerificationRelationships) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	// if controller is set must be compliant
	if !IsValidDID(msg.MethodId) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, "controller validation error")
	}

	// there should be more then one relationship
	if len(msg.Relationships) == 0 {
		return sdkerrors.Wrap(ErrEmptyRelationships, "controller validation error")
	}

	// TODO: there should be at least one authentication for the did subject

	return nil
}

// --------------------------
// ADD SERVICE
// --------------------------

// ValidateBasic performs a basic check of the MsgAddService fields.
func (msg MsgAddService) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if msg.ServiceData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "service is required")
	}

	if !IsValidRFC3986Uri(msg.ServiceData.Id) {
		return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	}

	if IsEmpty(msg.ServiceData.Type) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "service type required")
	}

	// XXX: compliance with the spec breaks the issuer module/flow
	if !IsValidRFC3986Uri(msg.ServiceData.ServiceEndpoint) {
		return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service endpoint validation error")
	}

	return nil
}

// --------------------------
// DELETE SERVICE
// --------------------------

// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteService) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidRFC3986Uri(msg.ServiceId) {
		return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	}
	return nil
}
