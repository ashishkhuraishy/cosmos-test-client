package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/ashishkhuraishy/test-net/testutil/keeper"
	"github.com/ashishkhuraishy/test-net/testutil/nullify"
	"github.com/ashishkhuraishy/test-net/x/testnet/keeper"
	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredPoll(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredPoll {
	items := make([]types.StoredPoll, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredPoll(ctx, items[i])
	}
	return items
}

func TestStoredPollGet(t *testing.T) {
	keeper, ctx := keepertest.TestnetKeeper(t)
	items := createNStoredPoll(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredPoll(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredPollRemove(t *testing.T) {
	keeper, ctx := keepertest.TestnetKeeper(t)
	items := createNStoredPoll(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredPoll(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredPoll(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredPollGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestnetKeeper(t)
	items := createNStoredPoll(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredPoll(ctx)),
	)
}
