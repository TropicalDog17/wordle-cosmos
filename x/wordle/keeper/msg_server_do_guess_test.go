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
func TestPlayToWinBeforeOutOfMove(t *testing.T) {
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
		MoveCount:  1,
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
		MoveCount:  2,
	}, *playMoveResponse)

	playMoveResponse, err = msgServer.DoGuess(context, &types.MsgDoGuess{
		Creator:   bob,
		GameIndex: "1",
		Letter:    "secret",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgDoGuessResponse{
		GuessState: "TTTTTT",
		Win:        true,
		MoveCount:  3,
	}, *playMoveResponse)

	// Win now, so can't move anymore
	_, err = msgServer.DoGuess(context, &types.MsgDoGuess{
		Creator:   bob,
		GameIndex: "1",
		Letter:    "dsadsa",
	})
	require.Error(t, types.ErrGameFinished, err)

}

type GameGuessTest struct {
	guess      string
	guessState string
	win        bool
	moveCount  int
}

func TestPlayToOutOfMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t, "secret")

	testCases := []GameGuessTest{
		{"impose", "WWWWPP", false, 1},
		{"detect", "WTPPPT", false, 2},
		{"thanks", "PWWWWP", false, 3},
		{"update", "WWWWPP", false, 4},
		{"income", "WWTWWP", false, 5},
		{"quorum", "WWWTWW", false, 6},
	}
	for _, test := range testCases {
		playMoveResponse, err := msgServer.DoGuess(context, &types.MsgDoGuess{
			Creator:   bob,
			GameIndex: "1",
			Letter:    test.guess,
		})
		require.Nil(t, err)
		require.EqualValues(t, types.MsgDoGuessResponse{
			GuessState: test.guessState,
			Win:        test.win,
			MoveCount:  uint64(test.moveCount),
		}, *playMoveResponse)
	}
	// Out of move, so return an error
	_, err := msgServer.DoGuess(context, &types.MsgDoGuess{
		Creator:   bob,
		GameIndex: "1",
		Letter:    "somewr",
	})
	require.Error(t, types.ErrGameFinished, err)
}
