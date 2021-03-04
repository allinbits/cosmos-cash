package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// msg types
const (
	TypeMsgCreateIssuer = "create-identifier"
)

var _ sdk.Msg = &MsgCreateIssuer{}

// NewMsgCreateIssuer creates a new MsgCreateIssuer instance
func NewMsgCreateIssuer(
	token string,
	fee int32,
	owner string,
) *MsgCreateIssuer {
	return &MsgCreateIssuer{
		Token: token,
		Fee:   fee,
		Owner: owner,
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

// ValidateBasic performs a basic check of the MsgCreateIssuer fields.
func (msg MsgCreateIssuer) ValidateBasic() error {
	if msg.Token == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token is required")
	}

	if msg.Fee == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "please enter a fee greater than 0")
	}

	return nil

}

func (msg MsgCreateIssuer) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIssuer) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
