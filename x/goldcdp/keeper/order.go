package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"

	"github.com/bandprotocol/goldcdp/x/goldcdp/types"
)

func (k Keeper) AddOrder(ctx sdk.Context, buyer sdk.AccAddress, amount sdk.Coins) (uint64, error) {
	orderID := k.GetNextOrderCount(ctx)
	// TODO: Config chain name
	collateralChain := "band-cosmoshub"

	// TODO: Support only 1 coin
	if len(amount) != 1 {
		return 0, sdkerrors.Wrapf(types.ErrOnlyOneDenomAllowed, "%d denoms included", len(amount))
	}
	channelID, err := k.GetChannel(ctx, collateralChain, "transfer")
	if err != nil {
		return 0, err
	}
	prefix := transfer.GetDenomPrefix("transfer", channelID)
	if !strings.HasPrefix(amount[0].Denom, prefix) {
		return 0, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom was: %s", amount[0].Denom)
	}

	// escrow source tokens. It fails if balance insufficient.
	escrowAddress := types.GetEscrowAddress()
	err = k.BankKeeper.SendCoins(ctx, buyer, escrowAddress, amount)
	if err != nil {
		return 0, err
	}
	k.SetOrder(ctx, orderID, types.NewOrder(buyer, amount))

	return orderID, nil
}

// SetOrder saves the given order to the store without performing any validation.
func (k Keeper) SetOrder(ctx sdk.Context, id uint64, order types.Order) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.OrderStoreKey(id), k.cdc.MustMarshalBinaryBare(order))
}

// GetOrder gets the given order from the store
func (k Keeper) GetOrder(ctx sdk.Context, id uint64) (types.Order, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.OrderStoreKey(id)) {
		return types.Order{}, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "order %d not found", id)
	}
	bz := store.Get(types.OrderStoreKey(id))
	var order types.Order
	k.cdc.MustUnmarshalBinaryBare(bz, &order)
	return order, nil
}
