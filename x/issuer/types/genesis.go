package types

import (
	fmt "fmt"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis(chainName, controllerAddress string) *GenesisState {
	did := fmt.Sprint(didtypes.DID(chainName, "regulator"))
	didDoc, err := didtypes.NewDidDocument(did,
		didtypes.WithControllers(didtypes.DIDKey(controllerAddress)),
	)
	if err != nil {
		panic("error creating genesis file")
	}
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		RegulatorDidDocument: &didDoc,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if gs.RegulatorDidDocument == nil {
		return fmt.Errorf("the regulator did cannot be nil")
	}
	if len(gs.RegulatorDidDocument.Controller) == 0 {
		return fmt.Errorf("there must be one or more controller for the regulator did document")
	}
	return nil
}
