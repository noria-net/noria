package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	util "github.com/noria-net/noria/util"
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
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	var err error
	if params.Minters != "" {
		// validate that params.Minters is a comma separated list of addresses
		_, err = util.SplitStringIntoAddresses(params.Minters)
		if err != nil {
			return errors.New(Error_invalid_minter)
		}
	}

	if params.Denoms != "" {
		// validate that params.Denoms is a comma separated list of denoms
		_, err = util.SplitStringIntoDenoms(params.Denoms)
		if err != nil {
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
