package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCoinmasterMint = "mint"

var _ sdk.Msg = &MsgCoinmasterMint{}

func NewMsgCoinmasterMint(creator string, amount sdk.Coin) *MsgCoinmasterMint {
	return &MsgCoinmasterMint{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgCoinmasterMint) Route() string {
	return RouterKey
}

func (msg *MsgCoinmasterMint) Type() string {
	return TypeMsgCoinmasterMint
}

func (msg *MsgCoinmasterMint) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCoinmasterMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCoinmasterMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
