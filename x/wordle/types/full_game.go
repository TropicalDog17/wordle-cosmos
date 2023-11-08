package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame Game) GetPlayerAddress() (black sdk.AccAddress, err error) {
	player, err := sdk.AccAddressFromBech32(storedGame.Player)
	return player, sdkerrors.Wrapf(err, ErrInvalidPlayer.Error(), storedGame.Player)
}

// TODO: add more validation logic
func (storedGame Game) Validate() (err error) {
	_, err = storedGame.GetPlayerAddress()
	if err != nil {
		return err
	}
	return err
}
