package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "rollapp/testutil/keeper"
	"rollapp/x/rollapp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RollappKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
