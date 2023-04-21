package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCoinmasterBurn = "burn"

var _ sdk.Msg = &MsgCoinmasterBurn{}

func NewMsgCoinmasterBurn(creator string, amount sdk.Coin) *MsgCoinmasterBurn {
	return &MsgCoinmasterBurn{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgCoinmasterBurn) Route() string {
	return RouterKey
}

func (msg *MsgCoinmasterBurn) Type() string {
	return TypeMsgCoinmasterBurn
}

func (msg *MsgCoinmasterBurn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCoinmasterBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCoinmasterBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
