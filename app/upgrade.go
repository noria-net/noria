package app

import (
	tokenfactorytypes "github.com/CosmWasm/token-factory/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// UpgradeHandler h for software upgrade proposal
type UpgradeHandler struct {
	*WasmApp
}

// NewUpgradeHandler return new instance of UpgradeHandler
func NewUpgradeHandler(app *WasmApp) UpgradeHandler {
	return UpgradeHandler{app}
}

func (h UpgradeHandler) CreateUpgradeHandler(app *WasmApp) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// Run migrations
		versionMap, err := app.mm.RunMigrations(ctx, h.configurator, vm)

		// TokenFactory
		newTokenFactoryParams := tokenfactorytypes.Params{
			DenomCreationFee: sdk.NewCoins(sdk.NewCoin("ucrd", sdk.NewInt(1000000))),
		}
		app.TokenFactoryKeeper.SetParams(ctx, newTokenFactoryParams)
		logger.Info("set tokenfactory params")

		return versionMap, err
	}
}
