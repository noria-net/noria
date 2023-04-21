package bindings

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	coinmasterkeeper "github.com/noria-net/noria/x/coinmaster/keeper"
)

func RegisterCustomPlugins(
	coinmasterKeeper *coinmasterkeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(coinmasterKeeper)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(coinmasterKeeper),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
