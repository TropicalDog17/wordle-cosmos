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
		Index:  newIndex,
		Player: msg.Player,
		Secret: msg.Secret,
		IsWin:  false,
	}
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetGame(ctx, storedGame)
	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
