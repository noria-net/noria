package coinmaster_test

import (
	"testing"

	keepertest "github.com/noria-net/noria/testutil/keeper"
	"github.com/noria-net/noria/testutil/nullify"
	"github.com/noria-net/noria/x/coinmaster"
	"github.com/noria-net/noria/x/coinmaster/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NewCoinmasterKeeper(t)
	coinmaster.InitGenesis(ctx, *k, genesisState)
	got := coinmaster.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
