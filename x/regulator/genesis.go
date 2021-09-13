package regulator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
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
	// get the regulator parameters
	rp := genState.GetRegulators()
	// store the list of regulators
	k.SetRegulators(ctx, rp)
	// for each address generate a did and store it
	for _, ra := range rp.GetAddresses() {
		// build the did document and metadata
		didDoc, _ := didtypes.NewDidDocument(
			didtypes.NewChainDID(ctx.ChainID(), ra),
			didtypes.WithControllers(didtypes.NewKeyDID(ra).String()),
		)
		didMeta := didtypes.NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
		// store the document
		k.SetDidDocumentWithMeta(ctx, didDoc, didMeta)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
