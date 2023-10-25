package keeper

import (
	"github.com/ashishkhuraishy/test-net/x/testnet/types"
)

var _ types.QueryServer = Keeper{}
