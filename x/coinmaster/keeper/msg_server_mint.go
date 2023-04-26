package keeper

import (
	"context"

	"github.com/noria-net/noria/x/coinmaster/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Mint coins in a specified denomination
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgCoinmasterMint) (*types.MsgCoinmasterMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.ValidateMinter(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	coins := sdk.NewCoins(msg.Amount)
	err = k.ValidateDenoms(ctx, coins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName,
		addr,
		coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgCoinmasterMintResponse{}, nil
}
