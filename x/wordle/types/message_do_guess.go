package types

import (
	"strconv"

	"github.com/TropicalDog17/wordle/x/wordle/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDoGuess = "do_guess"

var _ sdk.Msg = &MsgDoGuess{}

func NewMsgDoGuess(creator string, gameIndex string, letter string) *MsgDoGuess {
	return &MsgDoGuess{
		Creator:   creator,
		GameIndex: gameIndex,
		Letter:    letter,
	}
}

func (msg *MsgDoGuess) Route() string {
	return RouterKey
}

func (msg *MsgDoGuess) Type() string {
	return TypeMsgDoGuess
}

func (msg *MsgDoGuess) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDoGuess) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDoGuess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !rules.ValidWord(msg.Letter) {
		return sdkerrors.Wrapf(ErrInvalidGuess, "word should only contain English letters")
	}
	if len(msg.Letter) < 4 || len(msg.Letter) > 10 {
		return sdkerrors.Wrapf(ErrInvalidGuess, "word should be in range 4-10, please check the secret length")
	}
	gameIndex, err := strconv.ParseInt(msg.GameIndex, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidGameIndex, "not parseable (%s)", err)
	}
	if uint64(gameIndex) < DefaultIndex {
		return sdkerrors.Wrapf(ErrInvalidGameIndex, "number too low (%d)", gameIndex)
	}
	return nil
}
