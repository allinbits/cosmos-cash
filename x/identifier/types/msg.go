package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	signer string,
) *MsgCreateIdentifier {
	return &MsgCreateIdentifier{
		Id:            id,
		Verifications: verifications,
		Services:      services,
		Signer:        signer,
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

func (msg MsgCreateIdentifier) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIdentifier) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------

// msg types
const (
	TypeMsgUpdateIdentifier = "update-identifier"
)

func NewMsgUpdateIdentifier(
	id string,
	controller string,
	signer string,
) *MsgUpdateIdentifier {
	return &MsgUpdateIdentifier{
		Id:         id,
		Controller: controller,
		Signer:     signer,
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

func (msg MsgUpdateIdentifier) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateIdentifier) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD VERIFICATION
// --------------------------
// msg types
const (
	TypeMsgAddVerification = "add-verification"
)

var _ sdk.Msg = &MsgAddVerification{}

// NewMsgAddVerification creates a new MsgAddVerification instance
func NewMsgAddVerification(
	id string,
	verification *Verification,
	signer string,
) *MsgAddVerification {
	return &MsgAddVerification{
		Id:           id,
		Verification: verification,
		Signer:       signer,
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

func (msg MsgAddVerification) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// REVOKE VERIFICATION
// --------------------------

// msg types
const (
	TypeMsgRevokeVerification = "revoke-verification"
)

var _ sdk.Msg = &MsgRevokeVerification{}

// NewMsgRevokeVerification creates a new MsgRevokeVerification instance
func NewMsgRevokeVerification(
	id string,
	methodID string,
	signer string,
) *MsgRevokeVerification {
	return &MsgRevokeVerification{
		Id:       id,
		MethodId: methodID,
		Signer:   signer,
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

func (msg MsgRevokeVerification) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------
// msg types
const (
	TypeMsgSetVerificationRelationships = "set-verification-relationships"
)

func NewMsgSetVerificationRelationships(
	id string,
	methodID string,
	relationships []string,
	signer string,
) *MsgSetVerificationRelationships {
	return &MsgSetVerificationRelationships{
		Id:            id,
		MethodId:      methodID,
		Relationships: relationships,
		Signer:        signer,
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

func (msg MsgSetVerificationRelationships) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgSetVerificationRelationships) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD SERVICE
// --------------------------

// msg types
const (
	TypeMsgAddService = "add-service"
)

var _ sdk.Msg = &MsgAddService{}

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddService(
	id string,
	service *Service,
	signer string,
) *MsgAddService {
	return &MsgAddService{
		Id:          id,
		ServiceData: service,
		Signer:      signer,
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

func (msg MsgAddService) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE SERVICE
// --------------------------

// msg types
const (
	TypeMsgDeleteService = "delete-service"
)

func NewMsgDeleteService(
	id string,
	serviceID string,
	signer string,
) *MsgDeleteService {
	return &MsgDeleteService{
		Id:        id,
		ServiceId: serviceID,
		Signer:    signer,
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

func (msg MsgDeleteService) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
