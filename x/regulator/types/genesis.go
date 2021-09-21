package types

import "fmt"

// this line is used by starport scaffolding # genesis/types/import
// this line is used by starport scaffolding # ibc/genesistype/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis(addresses ...string) *GenesisState {
	return &GenesisState{
		Regulators: &Regulators{
			Addresses: addresses,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if gs.Regulators == nil {
		return fmt.Errorf("invalid regulator params")
	}

	if len(gs.Regulators.GetAddresses()) == 0 {
		return fmt.Errorf("invalid regulator params")
	}

	return nil
}
