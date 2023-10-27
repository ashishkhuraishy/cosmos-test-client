package keeper

import (
	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredPoll set a specific storedPoll in the store from its index
func (k Keeper) SetStoredPoll(ctx sdk.Context, storedPoll types.StoredPoll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPollKeyPrefix))
	b := k.cdc.MustMarshal(&storedPoll)
	store.Set(types.StoredPollKey(
		storedPoll.Index,
	), b)
}

// GetStoredPoll returns a storedPoll from its index
func (k Keeper) GetStoredPoll(
	ctx sdk.Context,
	index string,

) (val types.StoredPoll, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPollKeyPrefix))

	b := store.Get(types.StoredPollKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredPoll removes a storedPoll from the store
func (k Keeper) RemoveStoredPoll(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPollKeyPrefix))
	store.Delete(types.StoredPollKey(
		index,
	))
}

// GetAllStoredPoll returns all storedPoll
func (k Keeper) GetAllStoredPoll(ctx sdk.Context) (list []types.StoredPoll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredPollKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredPoll
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
