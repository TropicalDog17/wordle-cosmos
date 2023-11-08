package types

import (
	"testing"

	"github.com/TropicalDog17/wordle/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDoGuess_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDoGuess
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDoGuess{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDoGuess{
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
