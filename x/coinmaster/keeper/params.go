package keeper

import (
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
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
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
