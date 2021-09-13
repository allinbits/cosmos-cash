package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamStoreKeyRegulatorsParams key to store the parameters on
var (
	ParamStoreKeyRegulatorsParams = []byte("regulators_keys")
)

// ParamKeyTable - Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyRegulatorsParams, RegulatorsParams{}, validateRegulatorsParams),
	)
}

// NewRegulatorsParams creates a new RegulatorsParams object
func NewRegulatorsParams(addresses ...string) RegulatorsParams {
	return RegulatorsParams{
		Addresses: addresses,
	}
}

// DefaultRegulatorsParams default parameters for deposits
func DefaultRegulatorsParams() RegulatorsParams {
	return NewRegulatorsParams()
}

// // String implements stringer interface
// func (dp RegulatorsParams) String() string {
// 	out, _ := yaml.Marshal(dp)
// 	return string(out)
// }

// Equal checks equality of RegulatorsParams
func (m RegulatorsParams) Equal(dp2 RegulatorsParams) bool {
	// TODO: do the equal right
	return len(dp2.Addresses) == len(m.Addresses)
}

func validateRegulatorsParams(i interface{}) error {
	v, ok := i.(RegulatorsParams)
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
	RegulatorsParams RegulatorsParams `json:"regulators_params" yaml:"regulators_params"`
}

// NewParams creates a new gov Params instance
func NewParams(m RegulatorsParams) Params {
	return Params{
		RegulatorsParams: m,
	}
}

// DefaultParams default governance params
func DefaultParams() Params {
	return NewParams(DefaultRegulatorsParams())
}
