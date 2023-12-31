package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/TropicalDog17/wordle/testutil/keeper"
	"github.com/TropicalDog17/wordle/x/wordle/keeper"
	"github.com/TropicalDog17/wordle/x/wordle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WordleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
