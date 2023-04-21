package bindings

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	coinmasterkeeper "github.com/noria-net/noria/x/coinmaster/keeper"
	coinmastertypes "github.com/noria-net/noria/x/coinmaster/types"
)

type QueryPlugin struct {
	coinmasterKeeper *coinmasterkeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(k *coinmasterkeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		coinmasterKeeper: k,
	}
}

func (qp QueryPlugin) GetParams(ctx sdk.Context) (*coinmastertypes.ParamsResponse, error) {
	params := qp.coinmasterKeeper.GetParams(ctx)
	return &coinmastertypes.ParamsResponse{
		Params: coinmastertypes.Params{
			Minters: params.Minters,
			Denoms:  params.Denoms,
		},
	}, nil
}
