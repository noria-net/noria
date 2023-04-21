package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noria-net/noria/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCoinmasterBurn_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCoinmasterBurn
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCoinmasterBurn{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCoinmasterBurn{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
