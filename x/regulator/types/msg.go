package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	vctypes "github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
)

// define message types
const (
	TypeMsgIssueRegulatorCredential    = "issuer-regulator-credential"
	TypeMsgIssueRegistrationCredential = "issuer-registration-credential"
	TypeMsgIssueLicenseCredential      = "issuer-license-credential"
)

var (
	_ sdk.Msg = &MsgIssueRegulatorCredential{}
	_ sdk.Msg = &MsgIssueLicenseCredential{}
	_ sdk.Msg = &MsgIssueRegistrationCredential{}
)

func NewMsgIssueRegulatorCredential(credential vctypes.VerifiableCredential, signerAccount string) *MsgIssueRegulatorCredential {
	return &MsgIssueRegulatorCredential{
		Owner:      signerAccount,
		Credential: &credential,
	}
}

func (m *MsgIssueRegulatorCredential) Route() string {
	return RouterKey
}

func (m *MsgIssueRegulatorCredential) Type() string {
	return TypeMsgIssueRegulatorCredential
}

func (m *MsgIssueRegulatorCredential) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgIssueRegulatorCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgIssueRegulatorCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
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
func (m *MsgIssueLicenseCredential) Route() string {
	return RouterKey
}

// Type returns the string name of the message
func (m *MsgIssueLicenseCredential) Type() string {
	return TypeMsgIssueLicenseCredential
}

// GetSigners returns the account addresses singing the message
func (m *MsgIssueLicenseCredential) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes returns the bytes of the signed message
func (m *MsgIssueLicenseCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic validation of the message
func (m *MsgIssueLicenseCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
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
func (m *MsgIssueRegistrationCredential) Route() string {
	return RouterKey
}

// Type returns the string name of the message
func (m *MsgIssueRegistrationCredential) Type() string {
	return TypeMsgIssueRegistrationCredential
}

// GetSigners returns the account addresses singing the message
func (m *MsgIssueRegistrationCredential) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes returns the bytes of the signed message
func (m *MsgIssueRegistrationCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic validation of the message
func (m *MsgIssueRegistrationCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
