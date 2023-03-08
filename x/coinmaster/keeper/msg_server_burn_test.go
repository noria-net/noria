package keeper_test

import (
	"fmt"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/noria-net/noria/mocks"
	keepertest "github.com/noria-net/noria/testutil/keeper"
	"github.com/noria-net/noria/x/coinmaster/keeper"
	"github.com/noria-net/noria/x/coinmaster/types"
	"github.com/stretchr/testify/assert"
)

func Test_Burn_BaseCase(t *testing.T) {

	bankKeeperMock := mocks.NewBankKeeper(t)
	k, ctx := keepertest.NewCoinmasterKeeperWithBankKeeper(t, bankKeeperMock)
	msgServer := keeper.NewMsgServerImpl(*k)
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	denom := "unoria"
	wCtx := sdk.WrapSDKContext(ctx)

	burnAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", denom))

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")
	creatorAccAddress, _ := sdk.AccAddressFromBech32(creator)

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(creator, denom))

	// BankKeeper expectations
	bankKeeperMock.EXPECT().SendCoinsFromAccountToModule(ctx, creatorAccAddress, types.ModuleName, sdk.NewCoins(burnAmount)).Return(nil).Once()
	bankKeeperMock.EXPECT().BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnAmount)).Return(nil).Once()

	// Base case
	resp, err := msgServer.Burn(wCtx, &types.MsgBurn{
		Creator: creator,
		Amount:  burnAmount,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Burn_MultiplePermissableDenoms(t *testing.T) {

	bankKeeperMock := mocks.NewBankKeeper(t)
	k, ctx := keepertest.NewCoinmasterKeeperWithBankKeeper(t, bankKeeperMock)
	msgServer := keeper.NewMsgServerImpl(*k)
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	permittedDenoms := []string{"unoria", "vnoria", "wnoria"}
	combinedPermittedDenoms := strings.Join([]string{"unoria", "vnoria", "wnoria"}, ",")
	wCtx := sdk.WrapSDKContext(ctx)

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")
	creatorAccAddress, _ := sdk.AccAddressFromBech32(creator)

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(creator, combinedPermittedDenoms))

	for _, d := range permittedDenoms {
		burnAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", d))

		// BankKeeper expectation
		bankKeeperMock.EXPECT().SendCoinsFromAccountToModule(ctx, creatorAccAddress, types.ModuleName, sdk.NewCoins(burnAmount)).Return(nil).Once()
		bankKeeperMock.EXPECT().BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnAmount)).Return(nil).Once()

		// Base case
		resp, err := msgServer.Burn(wCtx, &types.MsgBurn{
			Creator: creator,
			Amount:  burnAmount,
		})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	}
}

func Test_Burn_DenomNotAllowed(t *testing.T) {

	k, ctx := keepertest.NewCoinmasterKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	denom := "unoria"
	wCtx := sdk.WrapSDKContext(ctx)

	mintAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", denom))

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(creator, "Xnoria"))

	_, err := msgServer.Burn(wCtx, &types.MsgBurn{
		Creator: creator,
		Amount:  mintAmount,
	})
	assert.ErrorContains(t, err, keeper.Error_unauthorized_denom)
}

func Test_Burn_BurnerNotAllowed(t *testing.T) {

	k, ctx := keepertest.NewCoinmasterKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	whitelistedCreator := "noriaWhitelistedCreator"
	denom := "unoria"
	wCtx := sdk.WrapSDKContext(ctx)

	burnAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", denom))

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(whitelistedCreator, "unoria"))

	_, err := msgServer.Burn(wCtx, &types.MsgBurn{
		Creator: creator,
		Amount:  burnAmount,
	})
	assert.ErrorContains(t, err, keeper.Error_unauthorized_account)
}
