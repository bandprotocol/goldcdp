package cli

import (
	"bufio"
	"fmt"
	"strconv"
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

	"github.com/bandprotocol/band-consumer/x/consuming/types"
	"github.com/bandprotocol/bandchain/chain/x/oracle"
)

const (
	flagName                     = "name"
	flagCalldata                 = "calldata"
	flagRequestedValidatorCount  = "requested-validator-count"
	flagSufficientValidatorCount = "sufficient-validator-count"
	flagExpiration               = "expiration"
	flagPrepareGas               = "prepare-gas"
	flagExecuteGas               = "execute-gas"
	flagChannel                  = "channel"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	consumingCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "consuming transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	consumingCmd.AddCommand(flags.PostCommands(
		GetCmdRequest(cdc),
	)...)

	return consumingCmd
}

// GetCmdRequest implements the request command handler.
func GetCmdRequest(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request [oracle-script-id] (-c [calldata]) (-r [requested-validator-count]) (-v [sufficient-validator-count]) (-x [expiration]) (-w [prepare-gas]) (-g [execute-gas])",
		Short: "Make a new data request via an existing oracle script",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Make a new request via an existing oracle script with the configuration flags.
Example:
$ %s tx consuming request 1 -c 1234abcdef -r 4 -v 3 -x 20 -w 50 -g 5000 --from mykey
$ %s tx consuming request 1 --calldata 1234abcdef --requested-validator-count 4 --sufficient-validator-count 3 --expiration 20 --prepare-gas 50 --execute-gas 5000 --from mykey
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			int64OracleScriptID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			oracleScriptID := oracle.OracleScriptID(int64OracleScriptID)

			calldata, err := cmd.Flags().GetBytesHex(flagCalldata)
			if err != nil {
				return err
			}

			requestedValidatorCount, err := cmd.Flags().GetInt64(flagRequestedValidatorCount)
			if err != nil {
				return err
			}

			sufficientValidatorCount, err := cmd.Flags().GetInt64(flagSufficientValidatorCount)
			if err != nil {
				return err
			}

			expiration, err := cmd.Flags().GetInt64(flagExpiration)
			if err != nil {
				return err
			}

			prepareGas, err := cmd.Flags().GetUint64(flagPrepareGas)
			if err != nil {
				return err
			}

			executionGas, err := cmd.Flags().GetUint64(flagExecuteGas)
			if err != nil {
				return err
			}

			channel, err := cmd.Flags().GetString(flagChannel)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestData(
				oracleScriptID,
				channel,
				calldata,
				requestedValidatorCount,
				sufficientValidatorCount,
				expiration,
				prepareGas,
				executionGas,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().BytesHexP(flagCalldata, "c", nil, "Calldata used in calling the oracle script")
	cmd.Flags().Int64P(flagRequestedValidatorCount, "r", 0, "Number of top validators that need to report data for this request")
	cmd.MarkFlagRequired(flagRequestedValidatorCount)
	cmd.Flags().Int64P(flagSufficientValidatorCount, "v", 0, "Minimum number of reports sufficient to conclude the request's result")
	cmd.MarkFlagRequired(flagSufficientValidatorCount)
	cmd.Flags().Int64P(flagExpiration, "x", 0, "Maximum block count before the data request is considered expired")
	cmd.MarkFlagRequired(flagExpiration)
	cmd.Flags().Uint64P(flagPrepareGas, "w", 0, "The amount of gas that will be reserved for prepare function")
	cmd.MarkFlagRequired(flagPrepareGas)
	cmd.Flags().Uint64P(flagExecuteGas, "g", 0, "The amount of gas that will be reserved for later execution")
	cmd.MarkFlagRequired(flagExecuteGas)
	cmd.Flags().String(flagChannel, "", "The channel id.")
	cmd.MarkFlagRequired(flagChannel)

	return cmd
}
