package types

import (
	"testing"

	"github.com/TropicalDog17/wordle/testutil/sample"
	"github.com/stretchr/testify/require"
)

var validSecret = []string{"correct", "good"}
var invalidSecret = []string{"abc", "toolongsecret", "12jl3", "k@pp@"}

func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateGame
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgCreateGame{
				Creator: "invalid_address",
				Player:  sample.AccAddress(),
				Secret:  validSecret[0],
			},
			err: ErrInvalidCreator,
		}, {
			name: "invalid Player address",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  "invalid_address",
				Secret:  validSecret[0],
			},
			err: ErrInvalidPlayer,
		},
		{
			name: "valid address",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  validSecret[0],
			},
		},
		{
			name: "invalid secret",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  invalidSecret[0],
			},
			err: ErrInvalidSecret,
		},
		{
			name: "invalid secret",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  invalidSecret[1],
			},
			err: ErrInvalidSecret,
		},
		{
			name: "invalid secret",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  invalidSecret[2],
			},
			err: ErrInvalidSecret,
		},
		{
			name: "invalid secret",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  invalidSecret[3],
			},
			err: ErrInvalidSecret,
		},
		{
			name: "valid secret",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  validSecret[0],
			},
		},
		{
			name: "valid secret",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				Player:  sample.AccAddress(),
				Secret:  validSecret[1],
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
