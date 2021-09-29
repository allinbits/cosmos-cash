package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
)

// msg types
const (
	TypeMsgCreateIssuer = "create-issuer"
	TypeMsgBurnToken    = "burn-token"
	TypeMsgMintToken    = "mint-token"
	TypeMsgPauseToken   = "pause-token"
)

var _ sdk.Msg = &MsgCreateIssuer{}

// NewMsgCreateIssuer creates a new MsgCreateIssuer instance
func NewMsgCreateIssuer(
	issuerDid string,
	licenseCredID string,
	token string,
	fee int32,
	owner string,
) *MsgCreateIssuer {
	return &MsgCreateIssuer{
		IssuerDid:     issuerDid,
		LicenseCredId: licenseCredID,
		Token:         token,
		Fee:           fee,
		Owner:         owner,
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
	if !didtypes.IsValidDID(msg.IssuerDid) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid did ID")
	}

	if msg.LicenseCredId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "licenseCredID is empty")
	}

	if msg.Token == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token is required")
	}

	if msg.Fee == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "please enter a fee greater than 0")
	}

	return nil

}

func (msg MsgCreateIssuer) GetSignBytes() []byte {
	panic("Issuer messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIssuer) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// Token burn
var _ sdk.Msg = &MsgBurnToken{}

// NewMsgBurnToken creates a new MsgBurnToken instance
func NewMsgBurnToken(
	issuerDid string,
	licenseCredID string,
	amount sdk.Coin,
	owner string,
) *MsgBurnToken {
	return &MsgBurnToken{
		IssuerDid:     issuerDid,
		LicenseCredId: licenseCredID,
		Amount:        amount.String(),
		Owner:         owner,
	}
}

// Route implements sdk.Msg
func (MsgBurnToken) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgBurnToken) Type() string {
	return TypeMsgBurnToken
}

// ValidateBasic performs a basic check of the MsgBurnToken fields.
func (msg MsgBurnToken) ValidateBasic() error {
	if !didtypes.IsValidDID(msg.IssuerDid) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid did ID")
	}

	if msg.LicenseCredId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "licenseCredID is empty")
	}

	amount, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return err
	}
	if amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "please enter an amount greater than 0")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgBurnToken) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func (msg MsgBurnToken) GetSignBytes() []byte {
	panic("Issuer messages do not support amino")
}

// Mint token
var _ sdk.Msg = &MsgMintToken{}

// NewMsgMintToken creates a new MsgMintToken instance
func NewMsgMintToken(
	issuerDid string,
	licenseCredID string,
	amount sdk.Coin,
	owner string,
) *MsgMintToken {
	return &MsgMintToken{
		IssuerDid:     issuerDid,
		LicenseCredId: licenseCredID,
		Amount:        amount.String(),
		Owner:         owner,
	}
}

// Route implements sdk.Msg
func (MsgMintToken) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgMintToken) Type() string {
	return TypeMsgMintToken
}

// ValidateBasic performs a basic check of the MsgMintToken fields.
func (msg MsgMintToken) ValidateBasic() error {
	if !didtypes.IsValidDID(msg.IssuerDid) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid did ID")
	}

	if msg.LicenseCredId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "licenseCredID is empty")
	}

	amount, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return err
	}
	if amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "please enter an amount greater than 0")
	}
	return nil
}

func (msg MsgMintToken) GetSignBytes() []byte {
	panic("Issuer messages do not support amino")
}

func (msg MsgMintToken) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// Pause Token
var _ sdk.Msg = &MsgPauseToken{}

// NewMsgPauseToken creates a new MsgPauseToken instance
func NewMsgPauseToken(
	issuerDid string,
	licenseCredID string,
	owner string,
) *MsgPauseToken {
	return &MsgPauseToken{
		IssuerDid:     issuerDid,
		LicenseCredId: licenseCredID,
		Owner:         owner,
	}
}

// Route implements sdk.Msg
func (MsgPauseToken) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgPauseToken) Type() string {
	return TypeMsgPauseToken
}

// ValidateBasic performs a basic check of the MsgPauseToken fields.
func (msg MsgPauseToken) ValidateBasic() error {
	if !didtypes.IsValidDID(msg.IssuerDid) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid did ID")
	}

	if msg.LicenseCredId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "licenseCredID is empty")
	}

	return nil

}

func (msg MsgPauseToken) GetSignBytes() []byte {
	panic("Issuer messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgPauseToken) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
