package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

var (
	_ sdk.Msg = &MsgActivate{}
	_ sdk.Msg = &MsgIssueLicenseCredential{}
	_ sdk.Msg = &MsgIssueRegistrationCredential{}
	_ sdk.Msg = &MsgRevokeCredential{}
)

func NewMsgActivate(credential vctypes.VerifiableCredential, signerAccount string) *MsgActivate {
	return &MsgActivate{
		Owner:      signerAccount,
		Credential: &credential,
	}
}

func (msg *MsgActivate) Route() string {
	return RouterKey
}

func (msg *MsgActivate) Type() string {
	return "Activate"
}

func (msg *MsgActivate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgActivate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgActivate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// NewMsgIssueLicenseCredential builds a new instance of a IssuerLicenceCredential message
func NewMsgIssueLicenseCredential(credential vctypes.VerifiableCredential, signerAccount string) *MsgIssueLicenseCredential {
	return &MsgIssueLicenseCredential{
		Credential: &credential,
		Owner:      signerAccount,
	}
}

// Route returns the module router key
func (msg *MsgIssueLicenseCredential) Route() string {
	return RouterKey
}

// Type returns the string name of the message
func (msg *MsgIssueLicenseCredential) Type() string {
	return "IssueLicense"
}

// GetSigners returns the account addresses singing the message
func (msg *MsgIssueLicenseCredential) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes returns the bytes of the signed message
func (msg *MsgIssueLicenseCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic validation of the message
func (msg *MsgIssueLicenseCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// NewMsgIssueRegistrationCredential builds a new instance of a the message
func NewMsgIssueRegistrationCredential(credential vctypes.VerifiableCredential, signerAccount string) *MsgIssueRegistrationCredential {
	return &MsgIssueRegistrationCredential{
		Credential: &credential,
		Owner:      signerAccount,
	}
}

// Route returns the module router key
func (msg *MsgIssueRegistrationCredential) Route() string {
	return RouterKey
}

// Type returns the string name of the message
func (msg *MsgIssueRegistrationCredential) Type() string {
	return "IssueRegistration"
}

// GetSigners returns the account addresses singing the message
func (msg *MsgIssueRegistrationCredential) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes returns the bytes of the signed message
func (msg *MsgIssueRegistrationCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic validation of the message
func (msg *MsgIssueRegistrationCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// NewMsgRevokeCredential builds a new instance of a the message
func NewMsgRevokeCredential(credentialID, signerAccount string) *MsgRevokeCredential {
	return &MsgRevokeCredential{
		CredentialId: credentialID,
		Owner:        signerAccount,
	}
}

// Route returns the module router key
func (msg *MsgRevokeCredential) Route() string {
	return RouterKey
}

// Type returns the string name of the message
func (msg *MsgRevokeCredential) Type() string {
	return "Revoke"
}

// GetSigners returns the account addresses singing the message
func (msg *MsgRevokeCredential) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes returns the bytes of the signed message
func (msg *MsgRevokeCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic validation of the message
func (msg *MsgRevokeCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
