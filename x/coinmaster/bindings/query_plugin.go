package bindings

import (
	"encoding/json"
	"fmt"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	coinmastertypes "github.com/noria-net/noria/x/coinmaster/types"
)

// CustomQuerier dispatches custom CosmWasm bindings queries.
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery coinmastertypes.CoinmasterQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "coinmaster query")
		}
		if contractQuery.Coinmaster == nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "nil coinmaster field")
		}
		coinmasterMsg := contractQuery.Coinmaster

		switch {
		case coinmasterMsg.Params != nil:
			res, err := qp.GetParams(ctx)
			if err != nil {
				return nil, err
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, fmt.Errorf("failed to JSON marshal ParamsResponse: %w", err)
			}

			return bz, nil

		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown coinmaster query variant"}
		}
	}
}
