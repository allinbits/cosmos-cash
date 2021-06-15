package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
)

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
	cmd.AddCommand(
		NewCreateIssuerCmd(),
		NewBurnTokenCmd(),
	)

	return cmd
}

// NewCreateIssuerCmd defines the command to create a new IBC light client.
func NewCreateIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-issuer [did] [token] [fee]",
		Short:   "create an issuer of an e-money token",
		Example: "creates an issuer of an e-money token",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			fee, _ := strconv.ParseInt(args[1], 0, 32)
			msg := types.NewMsgCreateIssuer(
				args[0],
				int32(fee),
				accAddrBech32,
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

// NewBurnTokenCmd defines the command to burn tokens.
func NewBurnTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-token [amount]",
		Short: "burn e-money tokens for an issuer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnToken(
				amount,
				accAddrBech32,
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

// TODO: extra commands:
//		Use:     "mint-issuer [did] [token] [fee]",
//		Use:     "deposit-emoney [did] [token] [fee]",
//		Use:     "widthdraw-emoney [did] [token] [fee]",
//		Use:     "freeze-all-emoney-tokens [did] [token] [fee]",
//		Use:     "freeze-account-with-emoney [did] [token] [fee]",
