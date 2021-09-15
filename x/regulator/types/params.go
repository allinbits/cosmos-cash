package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamStoreKeyRegulators key to store the parameters on
var (
	ParamStoreKeyRegulators = []byte("regulators_keys")
)

// ParamKeyTable - Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyRegulators, Regulators{}, validateRegulators),
	)
}

// NewRegulators creates a new Regulators object
func NewRegulators(addresses ...string) Regulators {
	return Regulators{
		Addresses: addresses,
	}
}

// DefaultRegulators default parameters for deposits
func DefaultRegulators() Regulators {
	return NewRegulators()
}

// // String implements stringer interface
// func (dp Regulators) String() string {
// 	out, _ := yaml.Marshal(dp)
// 	return string(out)
// }

// Equal checks equality of Regulators
func (m Regulators) Equal(dp2 Regulators) bool {
	// TODO: do the equal right
	return len(dp2.Addresses) == len(m.Addresses)
}

func validateRegulators(i interface{}) error {
	v, ok := i.(Regulators)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(v.Addresses) == 0 {
		return fmt.Errorf("at least one regulator address is required: %T", i)
	}
	// TODO check that every address is formally valid
	return nil
}

// Params returns all the regulators params
type Params struct {
	Regulators Regulators `json:"regulators_params" yaml:"regulators_params"`
}

// NewParams creates a new gov Params instance
func NewParams(m Regulators) Params {
	return Params{
		Regulators: m,
	}
}

// DefaultParams default governance params
func DefaultParams() Params {
	return NewParams(DefaultRegulators())
}
