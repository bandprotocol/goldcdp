package goldcdp

import (
	"github.com/bandprotocol/band-consumer/x/goldcdp/keeper"
	"github.com/bandprotocol/band-consumer/x/goldcdp/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier
)

type (
	Keeper              = keeper.Keeper
	MsgBuyGold          = types.MsgBuyGold
	MsgSetSourceChannel = types.MsgSetSourceChannel
)
