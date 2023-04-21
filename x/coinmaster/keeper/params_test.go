package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/noria-net/noria/testutil/keeper"
	"github.com/noria-net/noria/x/coinmaster/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewCoinmasterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Minters, k.Minters(ctx))
}

func TestSetParams(t *testing.T) {
	k, ctx := testkeeper.NewCoinmasterKeeper(t)
	invalidMinterMsg := "invalid minter"
	invalidDenomMsg := "invalid denomination"
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")
	addr := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"

	// Test invalid Minters parameter
	params := types.Params{Minters: "invalid"}
	err := k.SetParams(ctx, params)
	assert.EqualErrorf(t, err, invalidMinterMsg, "Error should be: %v, got: %v", invalidMinterMsg, err)

	// Test invalid Denoms parameter
	params = types.Params{Denoms: "no"}
	err = k.SetParams(ctx, params)
	assert.EqualErrorf(t, err, invalidDenomMsg, "Error should be: %v, got: %v", invalidDenomMsg, err)

	// Test valid parameters
	params = types.Params{Minters: addr, Denoms: "btc,eth"}
	err = k.SetParams(ctx, params)
	require.NoError(t, err)

	// Test Minters parameter contains invalid address
	params = types.Params{Minters: fmt.Sprintf("%v,invalid", addr)}
	err = k.SetParams(ctx, params)
	assert.EqualErrorf(t, err, invalidMinterMsg, "Error should be: %v, got: %v", invalidMinterMsg, err)

	// Test Denoms parameter contains invalid denomination
	params = types.Params{Denoms: "btc,n,,"}
	err = k.SetParams(ctx, params)
	assert.EqualErrorf(t, err, invalidDenomMsg, "Error should be: %v, got: %v", invalidDenomMsg, err)

	// Test Minters parameter contains multiple addresses
	params = types.Params{Minters: fmt.Sprintf("%v,%v", addr, addr)}
	err = k.SetParams(ctx, params)
	require.NoError(t, err)

	// Test Denoms parameter contains multiple denominations
	params = types.Params{Denoms: "btc,eth,usdt"}
	err = k.SetParams(ctx, params)
	require.NoError(t, err)
}
