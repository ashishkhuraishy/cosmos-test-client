package testnet_test

import (
	"testing"

	keepertest "github.com/ashishkhuraishy/test-net/testutil/keeper"
	"github.com/ashishkhuraishy/test-net/testutil/nullify"
	"github.com/ashishkhuraishy/test-net/x/testnet"
	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StoredPollList: []types.StoredPoll{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestnetKeeper(t)
	testnet.InitGenesis(ctx, *k, genesisState)
	got := testnet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StoredPollList, got.StoredPollList)
	// this line is used by starport scaffolding # genesis/test/assert
}
