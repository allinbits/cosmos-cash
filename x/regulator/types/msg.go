package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

var (
	_ sdk.Msg = &MsgIssueCredential{}
)

func NewMsgIssueCredential(credential vctypes.VerifiableCredential, signerAccount string) *MsgIssueCredential {
	return &MsgIssueCredential{
		Owner:      signerAccount,
		Credential: &credential,
	}
}

func (m *MsgIssueCredential) Route() string {
	return RouterKey
}

func (m *MsgIssueCredential) Type() string {
	return "Activate"
}

func (m *MsgIssueCredential) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgIssueCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgIssueCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
