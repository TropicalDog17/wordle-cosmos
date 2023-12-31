package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/TropicalDog17/wordle/testutil/keeper"

	"github.com/TropicalDog17/wordle/x/wordle"
	"github.com/TropicalDog17/wordle/x/wordle/keeper"
	"github.com/TropicalDog17/wordle/x/wordle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.WordleKeeper(t)
	wordle.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateNewGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	newGameResponse, _ := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player:  bob,
		Secret:  "hello",
	})
	assert.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "1",
	}, newGameResponse)
	systemInfo, found := keeper.GetSystemInfo(ctx)
	assert.True(t, found)
	assert.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "1",
		FifoTailIndex: "1",
	}, systemInfo)
	anotherGameResponse, _ := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player:  bob,
		Secret:  "abcdef",
	})
	assert.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "2",
	}, anotherGameResponse)
	systemInfo, found = keeper.GetSystemInfo(ctx)
	assert.True(t, found)
	assert.EqualValues(t, types.SystemInfo{
		NextId:        3,
		FifoHeadIndex: "2",
		FifoTailIndex: "2",
	}, systemInfo)

}
