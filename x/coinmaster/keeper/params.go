package keeper

import (
	"errors"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noria-net/noria/x/coinmaster/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Minters(ctx),
		k.Denoms(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) (err error) {
	if params.Minters != "" {
		// validate that params.Minters is a comma separated list of addresses
		mintersRegex := (`^([a-z0-9]{40,80},)*[a-z0-9]{40,80}$`)

		if match, err := regexp.MatchString(mintersRegex, params.Minters); !match || err != nil {
			return errors.New(Error_invalid_minter)
		}

		// validate that each minter is a valid address
		for _, minter := range strings.Split(params.Minters, ",") {
			_, err := sdk.AccAddressFromBech32(minter)
			if err != nil {
				return errors.New(Error_invalid_minter)
			}
		}
	}

	if params.Denoms != "" {
		// validate that params.Denoms is a comma separated list of denoms
		denomsRegex := `^([a-zA-Z0-9]{3,64},)*[a-zA-Z0-9]{3,64}$`
		if match, err := regexp.MatchString(denomsRegex, params.Denoms); !match || err != nil {
			return errors.New(Error_invalid_denom)
		}
	}

	k.paramstore.SetParamSet(ctx, &params)
	return nil
}

// Minters returns the Minters param
func (k Keeper) Minters(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMinters, &res)
	return
}

// Denoms returns the Denoms param
func (k Keeper) Denoms(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyDenoms, &res)
	return
}
