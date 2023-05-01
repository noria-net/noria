package keeper

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/noria-net/noria/x/coinmaster/keeper"
	"github.com/noria-net/noria/x/coinmaster/types"
	"github.com/stretchr/testify/require"
)

// NewCoinmasterKeeper: Creates a new test keeper
func NewCoinmasterKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	return NewCoinmasterKeeperWithBankKeeper(t, nil)
}

// NewCoinmasterKeeper: Creates a new test keeper with the specified bank keeper
func NewCoinmasterKeeperWithBankKeeper(t testing.TB, bankKeeper types.BankKeeper) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	// stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	// stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CoinmasterParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	err := k.SetParams(ctx, types.DefaultParams())
	if err != nil {
		panic(err)
	}

	return k, ctx
}
