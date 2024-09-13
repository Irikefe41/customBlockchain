package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tictactoe/testutil/keeper"
	"tictactoe/x/tictactoe/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TictactoeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
