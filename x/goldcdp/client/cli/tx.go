package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/band-consumer/x/goldcdp/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	goldcdpCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "goldcdp transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	goldcdpCmd.AddCommand(flags.PostCommands(
		GetCmdRequest(cdc),
		GetCmdSetChannel(cdc),
	)...)

	return goldcdpCmd
}

// GetCmdRequest implements the request command handler.
func GetCmdRequest(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [amount]",
		Short: "Make a new order to buy gold",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Make a new order to buy gold.
Example:
$ %s tx goldcdp buy 1000000dfsbsdfdf/transfer/uatom
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			amount, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgBuyGold(
				cliCtx.GetFromAddress(),
				amount,
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdSetChannel implements the set channel command handler.
func GetCmdSetChannel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-channel [chain-id] [port] [channel-id]",
		Short: "Register a verified channel",
		Args:  cobra.ExactArgs(3),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Register a verified channel.
Example:
$ %s tx goldcdp set-cahnnel bandchain goldcdp dbdfgsdfsd
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			msg := types.NewMsgSetSourceChannel(
				args[0],
				args[1],
				args[2],
				cliCtx.GetFromAddress(),
			)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
