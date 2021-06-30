package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// msg types
const (
	TypeMsgCreateIdentifier = "create-identifier"
)

var _ sdk.Msg = &MsgCreateIdentifier{}

// NewMsgCreateIdentifier creates a new MsgCreateIdentifier instance
func NewMsgCreateIdentifier(
	id string,
	verifications []*Verification,
	services []*Service,
	owner string,
) *MsgCreateIdentifier {
	return &MsgCreateIdentifier{
		Id:            id,
		Verifications: verifications,
		Services:      services,
		Owner:         owner,
	}
}

// Route implements sdk.Msg
func (MsgCreateIdentifier) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateIdentifier) Type() string {
	return TypeMsgCreateIdentifier
}

// ValidateBasic performs a basic check of the MsgCreateIdentifier fields.
func (msg MsgCreateIdentifier) ValidateBasic() error {
	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}

	if msg.Verifications == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verifications are required")
	}

	return nil

}

func (msg MsgCreateIdentifier) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIdentifier) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// msg types
const (
	TypeMsgAddVerification = "add-verification"
)

var _ sdk.Msg = &MsgAddVerification{}

// NewMsgAddVerification creates a new MsgAddVerification instance
func NewMsgAddVerification(
	id string,
	verification *Verification,
	owner string,
) *MsgAddVerification {
	return &MsgAddVerification{
		Id:           id,
		Verification: verification,
		Owner:        owner,
	}
}

// Route implements sdk.Msg
func (MsgAddVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddVerification) Type() string {
	return TypeMsgAddVerification
}

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

func (msg MsgAddVerification) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// msg types
const (
	TypeMsgAddService = "add-service"
)

var _ sdk.Msg = &MsgAddService{}

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddService(
	id string,
	service *Service,
	owner string,
) *MsgAddService {
	return &MsgAddService{
		Id:          id,
		ServiceData: service,
		Owner:       owner,
	}
}

// Route implements sdk.Msg
func (MsgAddService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddService) Type() string {
	return TypeMsgAddService
}

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

func (msg MsgAddService) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// msg types
const (
	TypeMsgRevokeVerification = "revoke-verification"
)

var _ sdk.Msg = &MsgRevokeVerification{}

// NewMsgRevokeVerification creates a new MsgRevokeVerification instance
func NewMsgRevokeVerification(
	id string,
	methodID string,
	owner string,
) *MsgRevokeVerification {
	return &MsgRevokeVerification{
		Id:       id,
		MethodId: methodID,
		Owner:    owner,
	}
}

// Route implements sdk.Msg
func (MsgRevokeVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgRevokeVerification) Type() string {
	return TypeMsgRevokeVerification
}

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

func (msg MsgRevokeVerification) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// msg types
const (
	TypeMsgDeleteService = "delete-service"
)

func NewMsgDeleteService(
	id string,
	serviceID string,
	owner string,
) *MsgDeleteService {
	return &MsgDeleteService{
		Id:        id,
		ServiceId: serviceID,
		Owner:     owner,
	}
}

// Route implements sdk.Msg
func (MsgDeleteService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteService) Type() string {
	return TypeMsgDeleteService
}

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

func (msg MsgDeleteService) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// -----------------------------------------

// msg types
const (
	TypeMsgUpdateIdentifier = "update-identifier"
)

func NewMsgUpdateIdentifier(
	id string,
	controller string,
	owner string,
) *MsgUpdateIdentifier {
	return &MsgUpdateIdentifier{
		Id:         id,
		Controller: controller,
		Owner:      owner,
	}
}

// Route implements sdk.Msg
func (MsgUpdateIdentifier) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgUpdateIdentifier) Type() string {
	return TypeMsgUpdateIdentifier
}

// ValidateBasic performs a basic check of the MsgUpdateIdentifier fields.
func (msg MsgUpdateIdentifier) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	// if controller is set must be compliant
	if !IsEmpty(msg.Controller) && !IsValidDID(msg.Controller) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, "controller validation error")
	}
	return nil
}

func (msg MsgUpdateIdentifier) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateIdentifier) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// -----------------------------------------

// msg types
const (
	TypeMsgSetVerificationRelationships = "update-identifier"
)

func NewMsgSetVerificationRelationships(
	id string,
	methodID string,
	relationships []VerificationRelationship,
	owner string,
) *MsgSetVerificationRelationships {
	return &MsgSetVerificationRelationships{
		Id:            id,
		MethodId:      methodID,
		Relationships: relationships,
		Owner:         owner,
	}
}

// Route implements sdk.Msg
func (MsgSetVerificationRelationships) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgSetVerificationRelationships) Type() string {
	return TypeMsgSetVerificationRelationships
}

// ValidateBasic performs a basic check of the MsgSetVerificationRelationships fields.
func (msg MsgSetVerificationRelationships) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	// if controller is set must be compliant
	if !IsValidDID(msg.MethodId) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, "controller validation error")
	}

	return nil
}

func (msg MsgSetVerificationRelationships) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgSetVerificationRelationships) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
