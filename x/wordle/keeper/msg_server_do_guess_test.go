package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/TropicalDog17/wordle/testutil/keeper"
	"github.com/TropicalDog17/wordle/x/wordle"
	"github.com/TropicalDog17/wordle/x/wordle/keeper"
	"github.com/TropicalDog17/wordle/x/wordle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneGameForPlayMove(t testing.TB, secret string) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.WordleKeeper(t)
	wordle.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player:  bob,
		Secret:  secret,
	})
	return server, *k, context
}
func TestPlayMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t, "secret")
	playMoveResponse, err := msgServer.DoGuess(context, &types.MsgDoGuess{
		Creator:   bob,
		GameIndex: "1",
		Letter:    "please",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgDoGuessResponse{
		GuessState: "WWPWPP",
		Win:        false,
	}, *playMoveResponse)
	playMoveResponse, err = msgServer.DoGuess(context, &types.MsgDoGuess{
		Creator:   bob,
		GameIndex: "1",
		Letter:    "update",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgDoGuessResponse{
		GuessState: "WWWWPP",
		Win:        false,
	}, *playMoveResponse)
}
