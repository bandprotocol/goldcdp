package goldcdp

import (
	"github.com/bandprotocol/goldcdp/x/goldcdp/keeper"
	"github.com/bandprotocol/goldcdp/x/goldcdp/types"
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
