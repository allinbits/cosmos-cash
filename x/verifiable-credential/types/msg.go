package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
)

// msg types
const (
	TypeMsgDeleteVerifiableCredential = "delete-verifiable-credential"
)

var (
	_ sdk.Msg = &MsgDeleteVerifiableCredential{}
)

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
	if !didtypes.IsValidDID(msg.IssuerDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDIDFormat, msg.IssuerDid)
	}

	return nil
}

func (msg MsgDeleteVerifiableCredential) GetSignBytes() []byte {
	panic("VerifiableCredential messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteVerifiableCredential) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
