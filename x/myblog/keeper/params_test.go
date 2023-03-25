package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "myblog/testutil/keeper"
	"myblog/x/myblog/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MyblogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
