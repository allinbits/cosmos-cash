package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// msg types
const (
	TypeMsgCreateVerifiableCredential = "create-verifiable-credential"
	TypeMsgDeleteVerifiableCredential = "delete-verifiable-credential"
)

var (
	_ sdk.Msg = &MsgCreateVerifiableCredential{}
	_ sdk.Msg = &MsgDeleteVerifiableCredential{}
)

// NewMsgCreateVerifiableCredential creates a new MsgCreateVerifiableCredential instance
func NewMsgCreateVerifiableCredential(
	vc VerifiableCredential,
	owner string,
) *MsgCreateVerifiableCredential {
	return &MsgCreateVerifiableCredential{
		VerifiableCredential: &vc,
		Owner:                owner,
	}
}

// Route implements sdk.Msg
func (MsgCreateVerifiableCredential) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateVerifiableCredential) Type() string {
	return TypeMsgCreateVerifiableCredential
}

// ValidateBasic performs a basic check of the MsgCreateVerifiableCredential fields.
func (msg MsgCreateVerifiableCredential) ValidateBasic() error {
	if msg.VerifiableCredential == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty verifiable cred")
	}

	return nil

}

func (msg MsgCreateVerifiableCredential) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateVerifiableCredential) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// NewMsgDeleteVerifiableCredential creates a new MsgDeleteVerifiableCredential instance
func NewMsgDeleteVerifiableCredential(
	id string,
	issuerDid string,
	owner string,
) *MsgDeleteVerifiableCredential {
	return &MsgDeleteVerifiableCredential{
		VerifiableCredentialId: id,
		IssuerDid:              issuerDid,
		Owner:                  owner,
	}
}

// Route implements sdk.Msg
func (msg MsgDeleteVerifiableCredential) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgDeleteVerifiableCredential) Type() string {
	return TypeMsgDeleteVerifiableCredential
}

// ValidateBasic performs a basic check of the MsgDeleteVerifiableCredential fields.
func (msg MsgDeleteVerifiableCredential) ValidateBasic() error {
	if msg.VerifiableCredentialId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty verifiable cred")
	}

	return nil
}

func (msg MsgDeleteVerifiableCredential) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteVerifiableCredential) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
