package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
)

var (
	// DefaultRelativePacketTimeoutTimestamp default timeout
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// const (
// 	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
// )

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdActivate())

	return cmd
}

var _ = strconv.Itoa(0)

func CmdActivate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate [did_id] [name] [countryCode]",
		Short: "Broadcast message to activate a regulator did",
		Long: `Regulators addresses are stored in the regulator genesis parameters, 
a did for each regulator is generated at genesis time but is not active, that is, 
a regulator must activate its DID document via this command.
The command will trigger the generation of a new verifiable credential for the regulator 
that activates it.`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			didId := args[0]
			name := args[1]
			countryCode := args[1]

			msg := types.NewMsgActivate(
				clientCtx.GetFromAddress().String(),
				didId,
				name,
				countryCode,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
