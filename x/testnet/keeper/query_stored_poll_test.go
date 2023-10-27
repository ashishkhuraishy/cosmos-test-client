package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ashishkhuraishy/test-net/testutil/keeper"
	"github.com/ashishkhuraishy/test-net/testutil/nullify"
	"github.com/ashishkhuraishy/test-net/x/testnet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStoredPollQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TestnetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStoredPoll(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStoredPollRequest
		response *types.QueryGetStoredPollResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStoredPollRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetStoredPollResponse{StoredPoll: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetStoredPollRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetStoredPollResponse{StoredPoll: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetStoredPollRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.StoredPoll(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestStoredPollQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TestnetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStoredPoll(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStoredPollRequest {
		return &types.QueryAllStoredPollRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StoredPollAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.StoredPoll), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.StoredPoll),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StoredPollAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.StoredPoll), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.StoredPoll),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StoredPollAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.StoredPoll),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StoredPollAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
