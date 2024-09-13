package tictactoe_test

import (
	"testing"

	keepertest "tictactoe/testutil/keeper"
	"tictactoe/testutil/nullify"
	tictactoe "tictactoe/x/tictactoe/module"
	"tictactoe/x/tictactoe/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, k, genesisState)
	got := tictactoe.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
