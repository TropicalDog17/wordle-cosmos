package keeper_test

import (
	"testing"

	testkeeper "github.com/TropicalDog17/wordle/testutil/keeper"
	"github.com/TropicalDog17/wordle/x/wordle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WordleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
