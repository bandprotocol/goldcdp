package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/goldcdp/x/goldcdp/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	BankKeeper    types.BankKeeper
	ChannelKeeper types.ChannelKeeper
}

// NewKeeper creates a new band consumer Keeper instance.
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		BankKeeper:    bankKeeper,
		ChannelKeeper: channelKeeper,
	}
}

// GetOrderCount returns the current number of all orders ever exist.
func (k Keeper) GetOrderCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.OrdersCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetNextOrderCount increments and returns the current number of orders.
// If the global order count is not set, it initializes it with value 0.
func (k Keeper) GetNextOrderCount(ctx sdk.Context) uint64 {
	orderCount := k.GetOrderCount(ctx)
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(orderCount + 1)
	store.Set(types.OrdersCountStoreKey, bz)
	return orderCount + 1
}
