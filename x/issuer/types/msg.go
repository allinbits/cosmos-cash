package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// msg types
const (
	TypeMsgCreateIssuer = "transfer"
)

var _ sdk.Msg = &MsgCreateIssuer{}

// NewMsgCreateIdentifier creates a new MsgCreateIdentifier instance
func NewMsgCreateIssuer(name string, token string, fee uint32, state string, address string) *MsgCreateIssuer {
	return &MsgCreateIssuer{
		Name:    name,
		Token:   token,
		Fee:     0,
		State:   state,
		Address: address,
	}
}

// Route implements sdk.Msg
func (MsgCreateIssuer) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateIssuer) Type() string {
	return TypeMsgCreateIssuer
}

// ValidateBasic performs a basic check of the MsgCreateIdentifier fields.
func (msg MsgCreateIssuer) ValidateBasic() error {
	return nil
}

func (msg MsgCreateIssuer) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIssuer) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
