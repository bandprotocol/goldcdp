package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	// ModuleName is the name of the module
	ModuleName = "goldcdp"
	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// GlobalStoreKeyPrefix is a prefix for global primitive state variable
	GlobalStoreKeyPrefix = []byte{0x00}

	// OrdersCountStoreKey is a key that help getting to current orders count state variable
	OrdersCountStoreKey = append(GlobalStoreKeyPrefix, []byte("OrdersCount")...)

	// ChannelStoreKeyPrefix is a prefix for storing channel
	ChannelStoreKeyPrefix = []byte{0x01}

	// OrderStoreKeyPrefix is a prefix for storing order
	OrderStoreKeyPrefix = []byte{0x02}
)

// ChannelStoreKey is a function to generate key for each verified channel in store
func ChannelStoreKey(chainName, channelPort string) []byte {
	buf := append(ChannelStoreKeyPrefix, []byte(chainName)...)
	buf = append(buf, []byte(channelPort)...)
	return buf
}

// OrderStoreKey is a function to generate key for each order in store
func OrderStoreKey(orderID uint64) []byte {
	return append(OrderStoreKeyPrefix, uint64ToBytes(orderID)...)
}

func uint64ToBytes(num uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, num)
	return result
}

func GetEscrowAddress() sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte("COLLATERAL")))
}
