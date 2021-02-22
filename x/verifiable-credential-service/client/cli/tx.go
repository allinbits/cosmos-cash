package cli

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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
		NewCreateVerifiableCredentialCmd(),
	)

	return cmd
}

// NewCreateVerifiableCredentialCmd defines the command to create a new verifiable credential.
func NewCreateVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-verifiable-credential [id]",
		Short:   "create decentralized verifiable-credential",
		Example: fmt.Sprintf("creates a verifiable credential for users"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			//cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			cs := types.NewCredentialSubject(
				args[0],
				true,
			)
			tm := time.Now()

			vc := types.NewVerifiableCredential(
				"new-verifiable-cred-3",
				[]string{"VerifiableCredential", "KYCCredential"},
				accAddrBech32,
				fmt.Sprintf("%s", tm),
				cs,
				types.NewProof("", "", "", "", ""),
			)

			signature, pubKey, err := clientCtx.Keyring.SignByAddress(accAddr, vc.GetBytes())
			if err != nil {
				return err
			}

			p := types.NewProof(
				pubKey.Type(),
				fmt.Sprintf("%s", tm),
				"assertionMethod",
				accAddrBech32+"#keys-1",
				base64.StdEncoding.EncodeToString(signature),
			)
			vc.Proof = &p

			msg := types.NewMsgCreateVerifiableCredential(
				vc,
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
