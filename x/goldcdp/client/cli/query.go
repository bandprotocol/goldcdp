package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/goldcdp/x/goldcdp/types"
)

// GetQueryCmd returns
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	goldcdpCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the goldcdp module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	goldcdpCmd.AddCommand(flags.GetCommands(
		GetCmdReadOrder(storeKey, cdc),
	)...)

	return goldcdpCmd
}

// GetCmdReadOrder queries order by orderID
func GetCmdReadOrder(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "order",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			orderID := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/order/%s", queryRoute, orderID),
				nil,
			)
			if err != nil {
				fmt.Printf("read request fail - %s \n", orderID)
				return nil
			}

			var order types.Order
			if err := cdc.UnmarshalJSON(res, &order); err != nil {
				return err
			}
			return cliCtx.PrintOutput(order)
		},
	}
}
