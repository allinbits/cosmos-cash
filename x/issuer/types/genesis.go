package types

import "fmt"

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default issuer genesis state
func DefaultGenesis(addresses ...string) *GenesisState {
	return &GenesisState{
		RegulatorsParams: &RegulatorsParams{
			Addresses: addresses,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (m GenesisState) Validate() error {
	if m.RegulatorsParams == nil {
		return fmt.Errorf("invalid regulator params")
	}

	if len(m.GetRegulatorsParams().Addresses) == 0 {
		return fmt.Errorf("invalid regulator params")
	}

	return nil
}
