package cli

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	merkletree "github.com/wealdtech/go-merkletree"
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
		NewCreateUserVerifiableCredentialCmd(),
		NewCreateIssuerVerifiableCredentialCmd(),
	)

	return cmd
}

// NewCreateUserVerifiableCredentialCmd defines the command to create a new verifiable credential.
func NewCreateUserVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-verifiable-credential [did_url] [cred-id] [secret] [name] [address] [date_of_birth] [nationalId] [phoneNumber]",
		Short:   "create decentralized verifiable-credential",
		Example: fmt.Sprintf("creates a verifiable credential for users"),
		Args:    cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			//cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			secret := args[2]

			data := [][]byte{
				[]byte(args[3]),
				[]byte(args[4]),
				[]byte(args[5]),
				[]byte(args[6]),
				[]byte(args[7]),
			}

			tree, err := merkletree.NewUsing(data, New(secret), nil)
			if err != nil {
				return err
			}

			root := tree.Root()
			hexRoot := hex.EncodeToString(root)

			cs := types.NewUserCredentialSubject(
				args[0],
				hexRoot,
				true,
			)
			tm := time.Now()

			vc := types.NewUserVerifiableCredential(
				args[1],
				[]string{"VerifiableCredential", "KYCCredential"},
				accAddrBech32,
				fmt.Sprintf("%s", tm),
				cs,
				types.NewProof("", "", "", "", ""),
			)

			// TODO: this could be expensive review this signing method
			// TODO: we can hash this an make this less expensive
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

// NewCreateIssuerVerifiableCredentialCmd defines the command to create a new verifiable credential.
func NewCreateIssuerVerifiableCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-issuer-verifiable-credential [did_url] [cred-id]",
		Short:   "create decentralized verifiable-credential for issuer",
		Example: fmt.Sprintf("creates a verifiable credential for issuers"),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			//cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			cs := types.NewIssuerCredentialSubject(
				args[0],
				true,
			)
			tm := time.Now()

			vc := types.NewIssuerVerifiableCredential(
				args[1],
				[]string{"VerifiableCredential", "IssuerCredential"},
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
