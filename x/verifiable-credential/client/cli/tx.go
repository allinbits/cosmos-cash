package cli

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/wealdtech/go-merkletree"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
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
		NewCreateKYCVerifiableCredentialCmd(),
		NewCreateLicenseVerifiableCredentialCmd(),
		NewDeleteVerifiableCredentialCmd(),
	)

	return cmd
}

// NewCreateKYCVerifiableCredentialCmd defines the command to create a new verifiable credential.
// TODO: to be moved to the issuer module
func NewCreateKYCVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `create-kyc-verifiable-credential [cred_subject] [cred_id] [issuer_did] [secret] [amount_per_transaction] [total_number_of_transactions] [total_transaction_amount]`,
		Short:   "create decentralized verifiable-credential",
		Example: "creates a verifiable credential for users",
		Args:    cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			credentialSubject := args[0]
			credentialID := args[1]
			issuerDid := args[2]
			secret := args[3]

			inputs := args[4:7]

			data := make([][]byte, len(inputs))
			for i, v := range inputs {
				data[i] = []byte(v)
			}

			tree, err := merkletree.NewUsing(data, New(secret), nil)
			if err != nil {
				return err
			}

			root := tree.Root()
			hexRoot := hex.EncodeToString(root)

			cs := types.NewUserCredentialSubject(
				credentialSubject,
				hexRoot,
				true,
			)
			tm := time.Now()

			vc := types.NewUserVerifiableCredential(
				credentialID,
				issuerDid,
				tm,
				cs,
			)

			signedVc := vc.Sign(clientCtx.Keyring, accAddr, issuerDid)

			msg := types.NewMsgCreateVerifiableCredential(
				signedVc,
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

// NewCreateLicenseVerifiableCredentialCmd defines the command to create a new license verifiable credential.
// This is used by regulators to define issuers and issuer permissions
// TODO: to be moved to regulator module
func NewCreateLicenseVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `create-license-verifiable-credential [cred_id] [issuer_did] [credential_subject_did] [type] [country] [authority] [denom] [circulation_limit]`,
		Short:   "create decentralized  verifiable-credential",
		Example: "creates a license verifiable credential for issuers",
		Args:    cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			credentialID := args[0]
			issuerDid := args[1]
			credentialSubject := args[2]
			licenseType := args[3]
			country := args[4]
			authority := args[5]
			denom := args[6]
			circulationLimitString := args[7]
			circulationLimit, _ := sdk.NewIntFromString(circulationLimitString)
			coin := sdk.NewCoin(denom, circulationLimit)

			cs := types.NewLicenseCredentialSubject(
				credentialSubject,
				licenseType,
				country,
				authority,
				coin,
			)
			tm := time.Now()

			vc := types.NewLicenseVerifiableCredential(
				credentialID,
				issuerDid,
				tm.UTC(),
				cs,
			)

			signedVc := vc.Sign(clientCtx.Keyring, accAddr, issuerDid)

			msg := types.NewMsgCreateVerifiableCredential(
				signedVc,
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

// NewDeleteVerifiableCredentialCmd defines the command to delete a verifiable credential.
// TODO: to me moved to the issuer module
func NewDeleteVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `delete-verifiable-credential [cred_id] [issuer_did]`,
		Short:   "delete a decentralized verifiable-credential",
		Example: "deletes a license verifiable credential for issuers",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			credentialID := args[0]
			issuerDid := args[1]

			msg := types.NewMsgDeleteVerifiableCredential(
				credentialID,
				issuerDid,
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
