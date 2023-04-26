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

func Test_Mint_BaseCase(t *testing.T) {

	bankKeeperMock := mocks.NewBankKeeper(t)
	k, ctx := keepertest.NewCoinmasterKeeperWithBankKeeper(t, bankKeeperMock)
	msgServer := keeper.NewMsgServerImpl(*k)
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	denom := "unoria"
	wCtx := sdk.WrapSDKContext(ctx)

	mintAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", denom))

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")
	creatorAccAddress, _ := sdk.AccAddressFromBech32(creator)

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(creator, denom))

	// BankKeeper expectations
	bankKeeperMock.EXPECT().MintCoins(ctx, "coinmaster", sdk.NewCoins(mintAmount)).Return(nil).Once()
	bankKeeperMock.EXPECT().SendCoinsFromModuleToAccount(ctx, "coinmaster", creatorAccAddress, sdk.NewCoins(mintAmount)).Return(nil).Once()

	// Base case
	resp, err := msgServer.Mint(wCtx, &types.MsgCoinmasterMint{
		Creator: creator,
		Amount:  mintAmount,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Mint_MultiplePermissableDenoms(t *testing.T) {

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
		mintAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", d))

		// BankKeeper expectation
		bankKeeperMock.EXPECT().MintCoins(ctx, "coinmaster", sdk.NewCoins(mintAmount)).Return(nil).Once()
		bankKeeperMock.EXPECT().SendCoinsFromModuleToAccount(ctx, "coinmaster", creatorAccAddress, sdk.NewCoins(mintAmount)).Return(nil).Once()

		// Base case
		resp, err := msgServer.Mint(wCtx, &types.MsgCoinmasterMint{
			Creator: creator,
			Amount:  mintAmount,
		})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	}
}

func Test_Mint_DenomNotAllowed(t *testing.T) {

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

	_, err := msgServer.Mint(wCtx, &types.MsgCoinmasterMint{
		Creator: creator,
		Amount:  mintAmount,
	})
	assert.ErrorContains(t, err, keeper.Error_unauthorized_denom)
}

func Test_Mint_MinterNotAllowed(t *testing.T) {

	k, ctx := keepertest.NewCoinmasterKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	whitelistedCreator := "noria19r3350dnszl6r7r9mtlneccr9p9hpwe6fscgkz"
	denom := "unoria"
	wCtx := sdk.WrapSDKContext(ctx)

	mintAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", denom))

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(whitelistedCreator, "unoria"))

	_, err := msgServer.Mint(wCtx, &types.MsgCoinmasterMint{
		Creator: creator,
		Amount:  mintAmount,
	})
	assert.ErrorContains(t, err, keeper.Error_unauthorized_account)
}
