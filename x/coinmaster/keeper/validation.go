package keeper

import (
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noria-net/noria/x/coinmaster/types"
	"golang.org/x/exp/slices"
)

func (k Keeper) ValidateDenoms(ctx sdk.Context, coins sdk.Coins) error {
	denoms := strings.Split(k.Denoms(ctx), ",")
	for _, coin := range coins {
		if !IsDenomWhiteListed(denoms, coin.Denom) {
			return errors.New(Error_unauthorized_denom)
		}
	}
	return nil
}

func (k Keeper) ValidateMinter(ctx sdk.Context, minter string) error {
	minters := k.Minters(ctx)
	if minters != types.DefaultMinters {
		minters := strings.Split(minters, ",")
		if slices.Contains(minters, minter) {
			return nil
		}
		return errors.New(Error_unauthorized_account)
	}
	return nil
}

// Determine whether the requested denom is allowed
func IsDenomWhiteListed(denoms []string, denom string) bool {

	if len(denoms) == 0 || len(denoms) == 1 && denoms[0] == "" {
		// Denom is allowed IF
		//  - there are no denoms in the whitelist
		//  - OR there is ONE denom and its value is empty
		return true
	} else {
		// ... otherwise the denom must exist in the whitelist as an EXACT match
		return slices.Contains(denoms, denom)
	}
}
