package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
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
		NewDeleteVerifiableCredentialCmd(),
		NewRevokeCredentialCmd(),
	)

	return cmd
}

// NewDeleteVerifiableCredentialCmd defines the command to delete a verifiable credential.
// TODO: to me moved to the issuer module
func NewDeleteVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `delete-verifiable-credential [cred_id] [issuer_did]`,
		Short:   "delete a decentralized verifiable-credential",
		Example: "deletes a license verifiable credential for issuers",
		Args:    cobra.ExactArgs(2),
		RunE:    revokeCredential,
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRevokeCredentialCmd defines the command to create a new license verifiable credential.
// This is used by regulators to define issuers and issuer permissions
func NewRevokeCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `revoke-credential [cred_id]`,
		Short:   "revoke a verifiable credential",
		Example: "",
		Args:    cobra.ExactArgs(1),
		RunE:    revokeCredential,
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func revokeCredential(cmd *cobra.Command, args []string) error {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	accAddr := clientCtx.GetFromAddress()
	accAddrBech32 := accAddr.String()

	credentialID := args[0]

	msg := types.NewMsgRevokeVerifiableCredential(credentialID, accAddrBech32)

	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
}
