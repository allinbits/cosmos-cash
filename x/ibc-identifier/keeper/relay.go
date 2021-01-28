package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	identifiertypes "github.com/allinbits/cosmos-cash/x/identifier/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"

	//	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	//	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/allinbits/cosmos-cash/x/ibc-identifier/types"
)

func (k Keeper) OnSendPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	return nil
}

func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	var data identifiertypes.DidDocument
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal identifier transfer packet data: %s", err.Error())
	}

	k.identifierKeeper.SetIdentifier(ctx, []byte(data.Id), data)

	id, _ := k.identifierKeeper.GetIdentifier(ctx, []byte(data.Id))
	fmt.Println(id)

	return nil
}

func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	return nil
}

func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	return nil
}
