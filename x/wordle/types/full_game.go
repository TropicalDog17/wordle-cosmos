package types

import (
	"time"

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
	_, err = storedGame.GetDeadlineAsTime()
	return err
}

func (storedGame Game) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, errDeadline := time.Parse(DeadlineLayout, storedGame.Deadline)
	return deadline, sdkerrors.Wrapf(errDeadline, ErrInvalidDeadline.Error(), storedGame.Deadline)
}

func FormatDeadline(deadline time.Time) string {
	return deadline.UTC().Format(DeadlineLayout)
}

func GetNextDeadline(ctx sdk.Context) time.Time {
	return ctx.BlockTime().Add(MaxTurnDuration)
}
