package bindings

import (
	"encoding/json"

	sdkerrors "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	coinmasterkeeper "github.com/noria-net/noria/x/coinmaster/keeper"
	coinmastertypes "github.com/noria-net/noria/x/coinmaster/types"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(coinmaster *coinmasterkeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:    old,
			coinmaster: coinmaster,
		}
	}
}

type CustomMessenger struct {
	wrapped    wasmkeeper.Messenger
	coinmaster *coinmasterkeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really creating / minting / swapping ...
		// leave everything else for the wrapped version
		var contractMsg coinmastertypes.CoinmasterMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterMsg, "requires 'coinmaster' field")
		}
		if contractMsg.Coinmaster == nil {
			return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
		}
		coinmasterMsg := contractMsg.Coinmaster

		if coinmasterMsg.Mint != nil {
			return m.mint(ctx, contractAddr, coinmasterMsg.Mint)
		}
		if coinmasterMsg.Burn != nil {
			return m.burn(ctx, contractAddr, coinmasterMsg.Burn)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// mint
func (m *CustomMessenger) mint(ctx sdk.Context, contractAddr sdk.AccAddress, msg *coinmastertypes.MsgCustomCoinmasterMint) ([]sdk.Event, [][]byte, error) {
	bz, err := PerformMint(m.coinmaster, ctx, contractAddr, msg)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterMintMsg, err.Error())
	}
	return nil, [][]byte{bz}, nil
}

func PerformMint(f *coinmasterkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, msg *coinmastertypes.MsgCustomCoinmasterMint) ([]byte, error) {
	if msg == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "nil mint message"}
	}

	msgServer := coinmasterkeeper.NewMsgServerImpl(*f)

	coins, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(coinmastertypes.ErrCoinmasterMintMsg, "unable to parse coin from 'amount' field: %s", err.Error())
	}

	msgMint := coinmastertypes.NewMsgCoinmasterMint(contractAddr.String(), coins)

	if err := msgMint.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrapf(coinmastertypes.ErrCoinmasterMintMsg, "failed validating MsgCoinmasterMint: %s", err.Error())
	}

	resp, err := msgServer.Mint(
		sdk.WrapSDKContext(ctx),
		msgMint,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterMintMsg, err.Error())
	}

	return resp.Marshal()
}

// burn tokens
func (m *CustomMessenger) burn(ctx sdk.Context, contractAddr sdk.AccAddress, msg *coinmastertypes.MsgCustomCoinmasterBurn) ([]sdk.Event, [][]byte, error) {
	bz, err := PerformBurn(m.coinmaster, ctx, contractAddr, msg)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterBurnMsg, err.Error())
	}
	return nil, [][]byte{bz}, nil
}

func PerformBurn(f *coinmasterkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, msg *coinmastertypes.MsgCustomCoinmasterBurn) ([]byte, error) {
	if msg == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "nil burn message"}
	}

	msgServer := coinmasterkeeper.NewMsgServerImpl(*f)

	coins, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(coinmastertypes.ErrCoinmasterBurnMsg, "unable to parse coin from 'amount' field: %s", err.Error())
	}

	msgBurn := coinmastertypes.NewMsgCoinmasterBurn(contractAddr.String(), coins)

	if err := msgBurn.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrapf(coinmastertypes.ErrCoinmasterBurnMsg, "failed validating MsgCoinmasterMint: %s", err.Error())
	}

	// Create denom
	resp, err := msgServer.Burn(
		sdk.WrapSDKContext(ctx),
		msgBurn,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(coinmastertypes.ErrCoinmasterBurnMsg, err.Error())
	}

	return resp.Marshal()
}
