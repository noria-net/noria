package app

import (
	tokenfactorytypes "github.com/CosmWasm/token-factory/x/tokenfactory/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	// UpgradeName gov proposal name
	UpgradeName = "v1.1.0"
)

func GetStoreUpgrades() *store.StoreUpgrades {
	return &store.StoreUpgrades{
		Added: []string{
			tokenfactorytypes.ModuleName,
		},
	}
}
