package keeper

import (
	"context"

	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) CreateItem(goCtx context.Context, msg *types.MsgCreateItem) (*types.MsgCreateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedPoll := types.StoredPoll{
		Index:    uuid.NewString(),
		Question: msg.Question,
		Voters:   msg.Options,
	}
	k.Keeper.SetStoredPoll(ctx, storedPoll)

	return &types.MsgCreateItemResponse{
		TestData: "sucess",
	}, nil
}
