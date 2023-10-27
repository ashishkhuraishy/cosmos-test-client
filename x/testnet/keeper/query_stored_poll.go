package keeper

import (
	"context"

	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredPollAll(goCtx context.Context, req *types.QueryAllStoredPollRequest) (*types.QueryAllStoredPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedPolls []types.StoredPoll
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storedPollStore := prefix.NewStore(store, types.KeyPrefix(types.StoredPollKeyPrefix))

	pageRes, err := query.Paginate(storedPollStore, req.Pagination, func(key []byte, value []byte) error {
		var storedPoll types.StoredPoll
		if err := k.cdc.Unmarshal(value, &storedPoll); err != nil {
			return err
		}

		storedPolls = append(storedPolls, storedPoll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredPollResponse{StoredPoll: storedPolls, Pagination: pageRes}, nil
}

func (k Keeper) StoredPoll(goCtx context.Context, req *types.QueryGetStoredPollRequest) (*types.QueryGetStoredPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStoredPoll(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredPollResponse{StoredPoll: val}, nil
}
