package regulator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/regulator/keeper"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
)

// DefaultGenesis returns the default issuer genesis state
func DefaultGenesis(addresses ...string) *types.GenesisState {
	return &types.GenesisState{
		Regulators: &types.Regulators{
			Addresses: addresses,
		},
	}
}

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// get the regulator parameters and store them on chain
	rp := genState.GetRegulators()
	k.SetRegulators(ctx, rp)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	genesis := types.DefaultGenesis(k.GetRegulatorsAddresses(ctx)...)

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
