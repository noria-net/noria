package app

import (
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	alliancemodule "github.com/noria-net/alliance/x/alliance"
	alliancemoduletypes "github.com/noria-net/alliance/x/alliance/types"
	ibchookstypes "github.com/noria-net/ibc-hooks/x/ibc-hooks/types"
)

const UpgradeName = "v1.3.0"

func (app WasmApp) RegisterUpgradeHandlers() {

	app.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

			// Set new params for the alliance module
			// Make the reward delay time 1 second for testnet, instead of a week
			newAllianceParams := alliancemoduletypes.NewParams()
			newAllianceParams.RewardDelayTime = time.Second * 1

			newAllianceGenesis := alliancemoduletypes.GenesisState{
				Params: newAllianceParams,
			}
			encoded, err := app.appCodec.MarshalJSON(&newAllianceGenesis)
			if err != nil {
				return nil, err
			}

			fromVM[alliancemoduletypes.ModuleName] = alliancemodule.AppModule{}.ConsensusVersion()
			module := app.ModuleManager.Modules[alliancemoduletypes.ModuleName].(alliancemodule.AppModule)
			module.InitGenesis(ctx, app.appCodec, encoded)

			return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{
				alliancemoduletypes.ModuleName,
				ibchookstypes.StoreKey,
			},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
