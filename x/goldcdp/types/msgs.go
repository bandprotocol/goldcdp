package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is they name of the goldcdp module
const RouterKey = ModuleName

// MsgSetSoruceChannel is a message for setting source channel to other chain
type MsgSetSourceChannel struct {
	ChainName     string         `json:"chain_name"`
	SourcePort    string         `json:"source_port"`
	SourceChannel string         `json:"source_channel"`
	Signer        sdk.AccAddress `json:"signer"`
}

func NewMsgSetSourceChannel(
	chainName, sourcePort, sourceChannel string,
	signer sdk.AccAddress,
) MsgSetSourceChannel {
	return MsgSetSourceChannel{
		ChainName:     chainName,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Signer:        signer,
	}
}

// Route implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Type() string { return "set_source_channel" }

// ValidateBasic implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) ValidateBasic() error {
	// TODO: Add validate basic
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// GetSignBytes implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgBuyGold is a message for creating order to buy gold
type MsgBuyGold struct {
	Buyer  sdk.AccAddress `json:"buyer"`
	Amount sdk.Coins      `json:"amount"`
}

// NewMsgBuyGold creates a new MsgBuyGold instance.
func NewMsgBuyGold(
	buyer sdk.AccAddress,
	amount sdk.Coins,
) MsgBuyGold {
	return MsgBuyGold{
		Buyer:  buyer,
		Amount: amount,
	}
}

// Route implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) Type() string { return "buy_gold" }

// ValidateBasic implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBuyGold: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBuyGold: Amount must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// GetSignBytes implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
