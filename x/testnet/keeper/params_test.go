package keeper_test

import (
	"testing"

	testkeeper "github.com/ashishkhuraishy/test-net/testutil/keeper"
	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestnetKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
