package types

import (
	"github.com/TropicalDog17/wordle/x/wordle/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGame = "create_game"

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, player string, secret string) *MsgCreateGame {
	return &MsgCreateGame{
		Creator: creator,
		Player:  player,
		Secret:  secret,
	}
}

func (msg *MsgCreateGame) Route() string {
	return RouterKey
}

func (msg *MsgCreateGame) Type() string {
	return TypeMsgCreateGame
}

func (msg *MsgCreateGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidCreator, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Player)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidPlayer, "invalid player address (%s)", err)
	}
	if !rules.ValidWord(msg.Secret) {
		return sdkerrors.Wrapf(ErrInvalidSecret, "secret word should contains English letters")

	}
	if len(msg.Secret) < 4 || len(msg.Secret) > 10 {
		return sdkerrors.Wrapf(ErrInvalidSecret, "secret word should be between 4 and 10 letters")
	}
	return nil
}
