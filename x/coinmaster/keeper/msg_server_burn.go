package keeper

import (
	"context"

	"github.com/noria-net/noria/x/coinmaster/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgCoinmasterBurn) (*types.MsgCoinmasterBurnResponse, error) {
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

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx,
		addr,
		types.ModuleName,
		coins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgCoinmasterBurnResponse{}, nil
}
