package types

import (
	"testing"

	"github.com/TropicalDog17/wordle/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)
const validSecret = []string{"correct", "good"}
const invalidSecret = []string{"abc", "toolongsecret", "12jl3", "k@pp@"}
func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateGame
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateGame{
				Creator: "invalid_address",

			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
			},
			{
				name: "invalid secret",
				msg: MsgCreateGame{

				}
			}
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
