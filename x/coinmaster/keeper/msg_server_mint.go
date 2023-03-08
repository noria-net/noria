package keeper

import (
	"context"
	"errors"
	"strings"

	"github.com/noria-net/noria/x/coinmaster/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Mint coins in a specified denomination
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
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

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName,
		addr,
		coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
