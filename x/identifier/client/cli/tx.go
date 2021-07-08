package cli

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
	vcstypes "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
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
		NewCreateIdentifierCmd(),
		NewAddVerificationCmd(),
		NewAddServiceCmd(),
		NewRevokeVerificationCmd(),
		NewDeleteServiceCmd(),
	)

	return cmd
}

// NewCreateIdentifierCmd defines the command to create a new IBC light client.
func NewCreateIdentifierCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-identifier [id]",
		Short:   "create decentralized identifier (did) document",
		Example: "creates a did document for users",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// did
			did := types.DID(args[0])

			// verification
			signer := clientCtx.GetFromAddress()
			signerDID := types.DID(signer.String())
			// pubkey
			info, err := clientCtx.Keyring.KeyByAddress(signer)
			if err != nil {
				return err
			}
			pubKey := info.GetPubKey()
			// verification method id
			vmID := fmt.Sprint(signerDID, "#", uuid.NewV4().String())

			auth := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					pubKey.Type(),
					signerDID,
					base58.Encode(pubKey.Bytes()),
				),
				[]string{types.RelationshipAuthentication},
				nil,
			)
			// create the message
			msg := types.NewMsgCreateIdentifier(
				did,
				[]*types.Verification{auth},
				[]*types.Service{},
				signer.String(),
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewAddVerificationCmd define the command to add a verification message
func NewAddVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-authentication [id] [pubkey]",
		Short:   "add an authentication method to a decentralized identifier (did) document",
		Example: "adds an authentication method for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// signer address
			signer := clientCtx.GetFromAddress()
			// public key
			pubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, args[1])
			if err != nil {
				return err
			}
			// document did
			did := types.DID(args[0])
			// controller did
			controllerDID := types.DID(signer.String())
			// verification method id
			vmID := fmt.Sprint(did, "#", uuid.NewV4().String())

			verification := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					pubKey.Type(),
					controllerDID,
					base58.Encode(pubKey.Bytes()),
				),
				[]string{types.RelationshipAuthentication},
				nil,
			)
			// add verification
			msg := types.NewMsgAddVerification(
				did,
				verification,
				signer.String(),
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

func NewAddServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-service [id] [service_id] [type] [endpoint]",
		Short:   "add a service to a decentralized identifier (did) document",
		Example: "adds a service to a did document",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if !vcstypes.IsValidCredentialType(args[2]) {
				return errors.New("invalid credential type")
			}
			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			serviceID, serviceType, endpoint := args[1], args[2], args[3]
			// document did
			did := types.DID(args[0])

			service := types.NewService(
				serviceID,
				serviceType,
				endpoint,
			)

			msg := types.NewMsgAddService(
				did,
				service,
				signer.String(),
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewRevokeVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "revoke-verification [id] [verification-method-id]",
		Short:   "revoke a verification method from a decentralized identifier (did) document",
		Example: "revoke a verification method for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// verification method id
			vmID := args[1]
			// build the message
			msg := types.NewMsgRevokeVerification(
				did,
				vmID,
				signer.String(),
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDeleteServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-service [id] [service-id]",
		Short:   "deletes a service from a decentralized identifier (did) document",
		Example: "delete a service for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// service id
			sID := args[1]

			msg := types.NewMsgDeleteService(
				did,
				sID,
				signer.String(),
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
