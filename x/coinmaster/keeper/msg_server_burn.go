package keeper

import (
	"context"
	"errors"
	"strings"

	"github.com/noria-net/noria/x/coinmaster/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minters := k.Minters(ctx)
	if minters != types.DefaultMinters {
		if msg.Creator != minters {
			return nil, errors.New(Error_unauthorized_account)
		}
	}

	coins := sdk.NewCoins(msg.Amount)

	denoms := strings.Split(k.Denoms(ctx), ",")
	if !IsDenomWhiteListed(denoms, coins[0].Denom) {
		return nil, errors.New(Error_unauthorized_denom)
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
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

	return &types.MsgBurnResponse{}, nil
}
