package keeper

import (
	"context"
	"strconv"

	"github.com/TropicalDog17/wordle/x/wordle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)
	storedGame := types.Game{
		Index:       newIndex,
		Player:      msg.Player,
		Secret:      msg.Secret,
		IsWin:       false,
		Deadline:    types.FormatDeadline(types.GetNextDeadline(ctx)),
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
	}
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	k.Keeper.SetGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, newIndex),
			sdk.NewAttribute(types.GameCreatedEventPlayer, msg.Player),
		),
	)

	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
