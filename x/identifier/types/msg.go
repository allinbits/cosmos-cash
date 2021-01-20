package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// msg types
const (
	TypeMsgCreateIdentifier = "transfer"
)

var _ sdk.Msg = &MsgCreateIdentifier{}

// NewMsgCreateIdentifier creates a new MsgCreateIdentifier instance
func NewMsgCreateIdentifier(id string) *MsgCreateIdentifier {
	return &MsgCreateIdentifier{
		Context: "context",
		Id:      id,
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
	return nil
}

func (msg MsgCreateIdentifier) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIdentifier) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Id)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
