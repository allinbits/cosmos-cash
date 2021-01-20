package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/allinbits/cosmos-cash/x/ibc-identifier/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	//capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	//channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	//paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler

	channelKeeper    types.ChannelKeeper
	portKeeper       types.PortKeeper
	authKeeper       types.AccountKeeper
	bankKeeper       types.BankKeeper
	scopedKeeper     capabilitykeeper.ScopedKeeper
	identifierKeeper types.IdentifierKeeper
}

func NewKeeper(
	cdc codec.BinaryMarshaler, key sdk.StoreKey,
	channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	identifierKeeper types.IdentifierKeeper,
) Keeper {

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		channelKeeper:    channelKeeper,
		portKeeper:       portKeeper,
		authKeeper:       authKeeper,
		bankKeeper:       bankKeeper,
		scopedKeeper:     scopedKeeper,
		identifierKeeper: identifierKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.scopedKeeper.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the transfer module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the transfer module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}
