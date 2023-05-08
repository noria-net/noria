package bindings

import (
	"encoding/json"

	sdkerrors "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	coinmasterkeeper "github.com/noria-net/noria/x/coinmaster/keeper"
	coinmastertypes "github.com/noria-net/noria/x/coinmaster/types"
)

type CustomQueryHandler struct {
	wrapped    wasmkeeper.WasmVMQueryHandler
	coinmaster *coinmasterkeeper.Keeper
}

func (qh CustomQueryHandler) GetParams(ctx sdk.Context) (*coinmastertypes.ParamsResponse, error) {
	params := qh.coinmaster.GetParams(ctx)
	return &coinmastertypes.ParamsResponse{
		Params: coinmastertypes.ParamsProperties{
			Minters: params.Minters,
			Denoms:  params.Denoms,
		},
	}, nil
}

func (m *CustomQueryHandler) HandleQuery(ctx sdk.Context, caller sdk.AccAddress, request wasmvmtypes.QueryRequest) ([]byte, error) {
	if request.Custom == nil {
		return m.wrapped.HandleQuery(ctx, caller, request)
	}
	customQuery := request.Custom

	var coinmasterQuery coinmastertypes.CoinmasterQuery
	if err := json.Unmarshal(customQuery, &coinmasterQuery); err != nil {
		return nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterMsg, "requires 'coinmaster' field")
	}
	if coinmasterQuery.Coinmaster == nil {
		return m.wrapped.HandleQuery(ctx, caller, request)
	}

	switch {
	case coinmasterQuery.Coinmaster.Params != nil:
		params, err := m.GetParams(ctx)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(params)
		if err != nil {
			return nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterQuery, "failed to marshall coinmaster ParamsResponse")
		}

		return bz, nil
	default:
		return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown coinmaster query variant"}
	}
}

func CustomQueryDecorator(coinmaster *coinmasterkeeper.Keeper) func(wasmkeeper.WasmVMQueryHandler) wasmkeeper.WasmVMQueryHandler {
	return func(old wasmkeeper.WasmVMQueryHandler) wasmkeeper.WasmVMQueryHandler {
		return &CustomQueryHandler{
			wrapped:    old,
			coinmaster: coinmaster,
		}
	}
}
