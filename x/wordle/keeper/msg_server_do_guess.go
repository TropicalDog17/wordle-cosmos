package keeper

import (
	"context"
	"strconv"

	"github.com/TropicalDog17/wordle/x/wordle/rules"
	"github.com/TropicalDog17/wordle/x/wordle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DoGuess(goCtx context.Context, msg *types.MsgDoGuess) (*types.MsgDoGuessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("System Info not found")
	}
	storedGame, found := k.Keeper.GetGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}
	if !storedGame.IsWin && storedGame.MoveCount >= uint64(rules.MaxMovesAllowed) {
		return nil, sdkerrors.Wrapf(types.ErrGameFinished, "out of move")
	}
	game := rules.Game{
		SecretWord: storedGame.Secret,
		MoveCount:  int(storedGame.MoveCount),
		IsWin:      storedGame.IsWin,
	}
	guessState, guessErr := game.Guess(msg.Letter)
	if guessErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidGuess, guessErr.Error())
	}
	storedGame.IsWin = game.IsWin
	if !storedGame.IsWin {
		k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	} else {
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &systemInfo)
	}
	// Update game to be stored
	storedGame.MoveCount = uint64(game.MoveCount)
	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx))
	k.Keeper.SetGame(ctx, storedGame)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.MovePlayedEventType,
			sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
			sdk.NewAttribute(types.MovePlayedEventGameIndex, msg.GameIndex),
			sdk.NewAttribute(types.MovePlayedEventWinner, strconv.FormatBool(game.IsWin)),
			sdk.NewAttribute(types.MovePlayedEventGuessState, rules.ParseGuessState(guessState)),
		),
	)
	return &types.MsgDoGuessResponse{
		GuessState: rules.ParseGuessState(guessState),
		Win:        storedGame.IsWin,
		MoveCount:  storedGame.MoveCount,
	}, nil
}
