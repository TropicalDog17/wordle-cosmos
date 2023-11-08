package wordle_test

import (
	"testing"

	keepertest "github.com/TropicalDog17/wordle/testutil/keeper"
	"github.com/TropicalDog17/wordle/testutil/nullify"
	"github.com/TropicalDog17/wordle/x/wordle"
	"github.com/TropicalDog17/wordle/x/wordle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 81,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WordleKeeper(t)
	wordle.InitGenesis(ctx, *k, genesisState)
	got := wordle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
