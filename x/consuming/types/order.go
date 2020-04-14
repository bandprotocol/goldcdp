package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderStatus uint8

const (
	Pending OrderStatus = iota
	Active
	Completed
)

type Order struct {
	Owner  sdk.AccAddress `json:"owner"`
	Amount sdk.Coins      `json:"amount"`
	Gold   sdk.Coin       `json:"gold"`
	Status OrderStatus    `json:"status"`
}

func NewOrder(owner sdk.AccAddress, amount sdk.Coins) Order {
	return Order{
		Owner:  owner,
		Amount: amount,
		Status: Pending,
	}
}
