package types

import (
	"github.com/bandprotocol/bandchain/chain/x/zoracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is they name of the bank module
const RouterKey = ModuleName

// MsgRequestData is a message for requesting a new data request to an existing oracle script.
type MsgRequestData struct {
	OracleScriptID           zoracle.OracleScriptID `json:"oracleScriptID"`
	SourceChannel            string                 `json:"sourceChannel"`
	Calldata                 []byte                 `json:"calldata"`
	RequestedValidatorCount  int64                  `json:"requestedValidatorCount"`
	SufficientValidatorCount int64                  `json:"sufficientValidatorCount"`
	Expiration               int64                  `json:"expiration"`
	PrepareGas               uint64                 `json:"prepareGas"`
	ExecuteGas               uint64                 `json:"executeGas"`
	Sender                   sdk.AccAddress         `json:"sender"`
}

// NewMsgRequestData creates a new MsgRequestData instance.
func NewMsgRequestData(
	oracleScriptID zoracle.OracleScriptID,
	sourceChannel string,
	calldata []byte,
	requestedValidatorCount int64,
	sufficientValidatorCount int64,
	expiration int64,
	prepareGas uint64,
	executeGas uint64,
	sender sdk.AccAddress,
) MsgRequestData {
	return MsgRequestData{
		OracleScriptID:           oracleScriptID,
		SourceChannel:            sourceChannel,
		Calldata:                 calldata,
		RequestedValidatorCount:  requestedValidatorCount,
		SufficientValidatorCount: sufficientValidatorCount,
		Expiration:               expiration,
		PrepareGas:               prepareGas,
		ExecuteGas:               executeGas,
		Sender:                   sender,
	}
}

// Route implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) Type() string { return "consuming" }

// ValidateBasic implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Sender address must not be empty.")
	}
	if msg.OracleScriptID <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Oracle script id (%d) must be positive.", msg.OracleScriptID)
	}
	if msg.SufficientValidatorCount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg,
			"MsgRequestData: Sufficient validator count (%d) must be positive.",
			msg.SufficientValidatorCount,
		)
	}
	if msg.RequestedValidatorCount < msg.SufficientValidatorCount {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg,
			"MsgRequestData: Request validator count (%d) must not be less than sufficient validator count (%d).",
			msg.RequestedValidatorCount,
			msg.SufficientValidatorCount,
		)
	}
	if msg.Expiration <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Expiration period (%d) must be positive.",
			msg.Expiration,
		)
	}
	if msg.PrepareGas <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Prepare gas (%d) must be positive.",
			msg.PrepareGas,
		)
	}
	if msg.ExecuteGas <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Execute gas (%d) must be positive.",
			msg.ExecuteGas,
		)
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
