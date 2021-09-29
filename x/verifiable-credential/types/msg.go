package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Message types types
const (
	TypeMsgDeleteVerifiableCredential = "delete-verifiable-credential"
	TypeMsgIssueVerifiableCredential  = "issue-verifiable-credential"
)

var (
	_ sdk.Msg = &MsgRevokeCredential{}
	_ sdk.Msg = &MsgIssueCredential{}
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

// GetSignBytes legacy amino
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

// NewMsgIssueCredential build a new message to issue credentials
func NewMsgIssueCredential(credential VerifiableCredential, signerAccount string) *MsgIssueCredential {
	return &MsgIssueCredential{
		Owner:      signerAccount,
		Credential: &credential,
	}
}

// Route implements sdk.Msg
func (m *MsgIssueCredential) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m *MsgIssueCredential) Type() string {
	return TypeMsgIssueVerifiableCredential
}

// GetSigners implements sdk.Msg
func (m *MsgIssueCredential) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes bytes of json serialization
func (m *MsgIssueCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validate a credential
func (m *MsgIssueCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
