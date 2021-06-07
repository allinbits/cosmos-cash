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
	authentication []*Authentication,
	owner string,
) *MsgCreateIdentifier {
	return &MsgCreateIdentifier{
		Context:        "https://www.w3.org/ns/did/v1",
		Id:             id,
		Authentication: authentication,
		Owner:          owner,
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

	if msg.Authentication == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "authentication is required")
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
	TypeMsgAddAuthentication = "add-authentication"
)

var _ sdk.Msg = &MsgAddAuthentication{}

// NewMsgAddAuthentication creates a new MsgAddAuthentication instance
func NewMsgAddAuthentication(
	id string,
	authentication *Authentication,
	owner string,
) *MsgAddAuthentication {
	return &MsgAddAuthentication{
		Id:             id,
		Authentication: authentication,
		Owner:          owner,
	}
}

// Route implements sdk.Msg
func (MsgAddAuthentication) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddAuthentication) Type() string {
	return TypeMsgAddAuthentication
}

// ValidateBasic performs a basic check of the MsgAddAuthentication fields.
func (msg MsgAddAuthentication) ValidateBasic() error {
	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}

	if msg.Authentication == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "authentication is required")
	}

	return nil

}

func (msg MsgAddAuthentication) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddAuthentication) GetSigners() []sdk.AccAddress {
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
	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}

	if msg.ServiceData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "service is required")
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
	TypeMsgDeleteAuthentication = "delete-authentication"
)

var _ sdk.Msg = &MsgDeleteAuthentication{}

// NewMsgDeleteAuthentication creates a new MsgDeleteAuthentication instance
func NewMsgDeleteAuthentication(
	id string,
	key string,
	owner string,
) *MsgDeleteAuthentication {
	return &MsgDeleteAuthentication{
		Id:    id,
		Key:   key,
		Owner: owner,
	}
}

// Route implements sdk.Msg
func (MsgDeleteAuthentication) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteAuthentication) Type() string {
	return TypeMsgDeleteAuthentication
}

// ValidateBasic performs a basic check of the MsgDeleteAuthentication fields.
func (msg MsgDeleteAuthentication) ValidateBasic() error {
	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}

	if msg.Key == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "authentication is required")
	}

	return nil

}

func (msg MsgDeleteAuthentication) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteAuthentication) GetSigners() []sdk.AccAddress {
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
	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}

	if msg.ServiceId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "authentication is required")
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
