package ibcidentifier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	//	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

var (
	_ porttypes.IBCModule = AppModule{}
)

//____________________________________________________________________________

// OnChanOpenInit implements the IBCModule interface
func (am AppModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	return am.keeper.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

// OnChanOpenTry implements the IBCModule interface
func (am AppModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version,
	counterpartyVersion string,
) error {
	return am.keeper.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version, counterpartyVersion)
}

// OnChanOpenAck implements the IBCModule interface
func (am AppModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyVersion string,
) error {
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (am AppModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (am AppModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for transfer channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (am AppModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (am AppModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
) (*sdk.Result, []byte, error) {
	acknowledgement := channeltypes.NewResultAcknowledgement([]byte{byte(0)})

	if err := am.keeper.OnRecvPacket(ctx, packet); err != nil {
		return nil, nil, err
	}

	//ctx.EventManager().EmitEvent(
	//	sdk.NewEvent(
	//		types.EventTypePacket,
	//		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	//		sdk.NewAttribute(types.AttributeKeyReceiver, data.Receiver),
	//		sdk.NewAttribute(types.AttributeKeyDenom, data.Denom),
	//		sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", data.Amount)),
	//		sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
	//	),
	//)

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, acknowledgement.GetBytes(), nil
}

// OnAcknowledgementPacket implements the IBCModule interface
func (am AppModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) (*sdk.Result, error) {
	if err := am.keeper.OnAcknowledgementPacket(ctx, packet, acknowledgement); err != nil {
		return nil, err
	}
	//	ctx.EventManager().EmitEvent(
	//		sdk.NewEvent(
	//			types.EventTypePacket,
	//			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	//			sdk.NewAttribute(types.AttributeKeyReceiver, data.Receiver),
	//			sdk.NewAttribute(types.AttributeKeyDenom, data.Denom),
	//			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", data.Amount)),
	//			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
	//		),
	//	)
	//
	//	switch resp := ack.Response.(type) {
	//	case *channeltypes.Acknowledgement_Result:
	//		ctx.EventManager().EmitEvent(
	//			sdk.NewEvent(
	//				types.EventTypePacket,
	//				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
	//			),
	//		)
	//	case *channeltypes.Acknowledgement_Error:
	//		ctx.EventManager().EmitEvent(
	//			sdk.NewEvent(
	//				types.EventTypePacket,
	//				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
	//			),
	//		)
	//	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// OnTimeoutPacket implements the IBCModule interface
func (am AppModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
) (*sdk.Result, error) {
	if err := am.keeper.OnTimeoutPacket(ctx, packet); err != nil {
		return nil, err
	}

	//	ctx.EventManager().EmitEvent(
	//		sdk.NewEvent(
	//			types.EventTypeTimeout,
	//			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	//			sdk.NewAttribute(types.AttributeKeyRefundReceiver, data.Sender),
	//			sdk.NewAttribute(types.AttributeKeyRefundDenom, data.Denom),
	//			sdk.NewAttribute(types.AttributeKeyRefundAmount, fmt.Sprintf("%d", data.Amount)),
	//		),
	//	)

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
