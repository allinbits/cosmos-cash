package ibcidentifier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/ibc-identifier/keeper"
	"github.com/allinbits/cosmos-cash/x/ibc-identifier/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetPort(ctx, genState.PortId)

	_ = k.BindPort(ctx, genState.PortId)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesis()
}
