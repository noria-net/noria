package wasm

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/noria-net/noria/x/coinmaster/types"
)

func NewMessageEncoders() *wasmkeeper.MessageEncoders {
	parser := NewWasmMsgParser()
	return &wasmkeeper.MessageEncoders{
		Custom: parser.ParseCustom,
	}
}

var ErrInvalidMsg = sdkerrors.Register(types.ModuleName, 8, "invalid Msg from the contract")

// WasmMsgParser - wasm msg parser for coinmaster msgs
type WasmMsgParser struct{}

// NewWasmMsgParser returns coinmaster wasm msg parser
func NewWasmMsgParser() WasmMsgParser {
	return WasmMsgParser{}
}

// CosmosMsg only contains mint and burn msg
type CosmosMsg struct {
	Mint *types.MsgMint `json:"mint,omitempty"`
	Burn *types.MsgBurn `json:"burn,omitempty"`
}

// ParseCustom implements custom parser
func (WasmMsgParser) ParseCustom(contractAddr sdk.AccAddress, msg json.RawMessage) ([]sdk.Msg, error) {
	var sdkMsg CosmosMsg
	err := json.Unmarshal(msg, &sdkMsg)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to parse market custom msg")
	}

	if sdkMsg.Mint != nil {
		sdkMsg.Mint.Creator = contractAddr.String()
		return []sdk.Msg{sdkMsg.Mint}, sdkMsg.Mint.ValidateBasic()
	} else if sdkMsg.Burn != nil {
		sdkMsg.Burn.Creator = contractAddr.String()
		return []sdk.Msg{sdkMsg.Burn}, sdkMsg.Burn.ValidateBasic()
	}

	return nil, sdkerrors.Wrap(ErrInvalidMsg, "Unknown variant of Coinmaster")
}
