package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/exp/slices"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/noria-net/noria/x/coinmaster/types"
)

// The Keeper instance
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramstore paramtypes.Subspace

	bankKeeper types.BankKeeper
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

// Creates a new Coinmaster keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
