package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Message types types
const (
	TypeMsgDeleteVerifiableCredential = "delete-verifiable-credential"
)

var (
	_ sdk.Msg = &MsgRevokeCredential{}
)

// NewMsgRevokeVerifiableCredential creates a new MsgDeleteVerifiableCredential instance
func NewMsgRevokeVerifiableCredential(
	id string,
	owner string,
) *MsgRevokeCredential {
	return &MsgRevokeCredential{
		CredentialId: id,
		Owner:        owner,
	}
}

// Route implements sdk.Msg
func (m MsgRevokeCredential) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m MsgRevokeCredential) Type() string {
	return TypeMsgDeleteVerifiableCredential
}

// ValidateBasic performs a basic check of the MsgDeleteVerifiableCredential fields.
func (m MsgRevokeCredential) ValidateBasic() error {
	if m.CredentialId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty verifiable cred")
	}
	return nil
}

func (m MsgRevokeCredential) GetSignBytes() []byte {
	panic("VerifiableCredential messages do not support amino")
}

// GetSigners implements sdk.Msg
func (m MsgRevokeCredential) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
