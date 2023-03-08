package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinters            = []byte("Minters")
	KeyDenoms             = []byte("Denoms")
	DefaultMinters string = ""
	DefaultDenoms  string = ""
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minters string,
	denoms string,
) Params {
	return Params{
		Minters: minters,
		Denoms:  denoms,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinters,
		DefaultDenoms,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinters, &p.Minters, validateMinters),
		paramtypes.NewParamSetPair(KeyDenoms, &p.Denoms, validateDenoms),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinters(p.Minters); err != nil {
		return err
	}
	if err := validateDenoms(p.Denoms); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMinters validates the Minters param
func validateMinters(v interface{}) error {
	minters, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minters

	return nil
}

// validateDenoms validates the Denoms param
func validateDenoms(v interface{}) error {
	denoms, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = denoms

	return nil
}
