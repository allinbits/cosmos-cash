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
		NewMintTokenCmd(),
		NewPauseTokenCmd(),
	)

	return cmd
}

// NewCreateIssuerCmd defines the command to create a new IBC light client.
func NewCreateIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-issuer [did] [license_cred_id] [token] [fee]",
		Short:   "create an issuer of an e-money token",
		Example: "creates an issuer of an e-money token",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()
			issuerDid := args[0]
			licenseCred := args[1]
			token := args[2]
			fee, _ := strconv.ParseInt(args[3], 0, 32)

			msg := types.NewMsgCreateIssuer(
				issuerDid,
				licenseCred,
				token,
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
		Use:   "burn-token [did] [license_cred_id] [amount]",
		Short: "burn e-money tokens for an issuer",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			didID := args[0]
			vc := args[1]

			// read the amount to burn
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			// build the burn message
			msg := types.NewMsgBurnToken(
				didID,
				vc,
				amount,
				accAddrBech32,
			)
			// validate the message
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// submit the transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

// NewMintTokenCmd defines the command to mint tokens.
func NewMintTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-token [did] [license_cred_id] [amount]",
		Short: "mint e-money tokens for an issuer",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			didID := args[0]
			vc := args[1]

			// read the amount to mint
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			// build the message
			msg := types.NewMsgMintToken(
				didID,
				vc,
				amount,
				accAddrBech32,
			)
			// validate the message
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// submit the transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

//NewPauseTokenCmd defines the command to pause all transfers of an emoney token.
func NewPauseTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-token [did] [license_cred_id]",
		Short: "pauses all transfers of an emoney token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()
			issuerDid := args[0]
			licenseCred := args[1]

			msg := types.NewMsgPauseToken(
				issuerDid,
				licenseCred,
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
//		Use:     "withdraw-emoney [did] [token] [amount]",
//		Use:     "freeze-account-with-emoney [did] [token] [fee]",
