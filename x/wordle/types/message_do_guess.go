package types

import (
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
	return nil
}
