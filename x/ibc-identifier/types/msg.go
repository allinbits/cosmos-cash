package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
)

// msg types
const (
	TypeMsgTransferIdentifierIBC = "ibcidentifier"
)

var _ sdk.Msg = &MsgTransferIdentifierIBC{}

// NewMsgTransferIdentifierIBC creates a new MsgTransferIdentifierIBC instance
func NewMsgTransferIdentifierIBC(id, port, channel string, height clienttypes.Height, timestamp uint64) *MsgTransferIdentifierIBC {
	return &MsgTransferIdentifierIBC{
		Id:               id,
		SourcePort:       port,
		SourceChannel:    channel,
		TimeoutHeight:    height,
		TimeoutTimestamp: timestamp,
	}
}

// Route implements sdk.Msg
func (MsgTransferIdentifierIBC) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgTransferIdentifierIBC) Type() string {
	return TypeMsgTransferIdentifierIBC
}

// ValidateBasic performs a basic check of the MsgTransferIdentifierIBC fields.
func (msg MsgTransferIdentifierIBC) ValidateBasic() error {
	return nil
}

func (msg MsgTransferIdentifierIBC) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgTransferIdentifierIBC) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Id)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
