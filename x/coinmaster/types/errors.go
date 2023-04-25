package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/coinmaster module sentinel errors
var (
	ErrCoinmasterMsg     = sdkerrors.Register(ModuleName, 1100, "invalid coinmaster message")
	ErrCoinmasterMintMsg = sdkerrors.Register(ModuleName, 1101, "error with coinmaster Mint message")
	ErrCoinmasterBurnMsg = sdkerrors.Register(ModuleName, 1102, "error with coinmaster Burn message")
	ErrCoinmasterQuery   = sdkerrors.Register(ModuleName, 1103, "error with coinmaster query")
)
