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
	adminmodule "github.com/noria-net/module-admin/x/admin"
	adminmoduletypes "github.com/noria-net/module-admin/x/admin/types"
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
			encodedAllianceGenesis, err := app.appCodec.MarshalJSON(&newAllianceGenesis)
			if err != nil {
				return nil, err
			}

			fromVM[alliancemoduletypes.ModuleName] = alliancemodule.AppModule{}.ConsensusVersion()
			allianceModule := app.ModuleManager.Modules[alliancemoduletypes.ModuleName].(alliancemodule.AppModule)
			allianceModule.InitGenesis(ctx, app.appCodec, encodedAllianceGenesis)

			// Set new params for the admin module
			// Set the initial admin address as the Noria Dev wallet
			newAdminParams := adminmoduletypes.NewParams("noria1heyrp6y4q8p8z706nv4payx0f8xpqytulqaf6q")

			newAdminGenesis := adminmoduletypes.GenesisState{
				Params: newAdminParams,
			}
			encodedAdminGenesis, err := app.appCodec.MarshalJSON(&newAdminGenesis)
			if err != nil {
				return nil, err
			}

			fromVM[adminmoduletypes.ModuleName] = adminmodule.AppModule{}.ConsensusVersion()
			adminModule := app.ModuleManager.Modules[adminmoduletypes.ModuleName].(adminmodule.AppModule)
			adminModule.InitGenesis(ctx, app.appCodec, encodedAdminGenesis)

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
				adminmoduletypes.ModuleName,
			},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
