package app

import (
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

func (h UpgradeHandler) CreateUpgradeHandler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		return h.mm.RunMigrations(ctx, h.configurator, vm)
	}
}
