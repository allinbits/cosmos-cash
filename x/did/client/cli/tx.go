package cli

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/did/types"
	vcstypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
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
		NewCreateDidDocumentCmd(),
		NewAddVerificationCmd(),
		NewAddServiceCmd(),
		NewRevokeVerificationCmd(),
		NewDeleteServiceCmd(),
		NewUpdateDidDocumentCmd(),
		NewSetVerificationRelationshipCmd(),
	)

	return cmd
}

// deriveVMType derive the verification method type from a public key
func deriveVMType(pubKey cryptotypes.PubKey) (vmType types.VerificationMaterialType, err error) {
	switch pubKey.(type) {
	case *ed25519.PubKey:
		vmType = types.DIDVMethodTypeEd25519VerificationKey2018
	case *secp256k1.PubKey:
		vmType = types.DIDVMethodTypeEcdsaSecp256k1RecoveryMethod2020
	default:
		err = types.ErrKeyFormatNotSupported
	}
	return
}

// NewCreateDidDocumentCmd defines the command to create a new IBC light client.
func NewCreateDidDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-did [id]",
		Short:   "create decentralized did (did) document",
		Example: "creates a did document for users",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// did
			did := types.DID(clientCtx.ChainID, args[0])
			// verification
			signer := clientCtx.GetFromAddress()
			// pubkey
			info, err := clientCtx.Keyring.KeyByAddress(signer)
			if err != nil {
				return err
			}
			pubKey := info.GetPubKey()
			// verification method id
			vmID := fmt.Sprint(did, "#", sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pubKey.Address().Bytes()))
			// understand the vmType
			vmType, err := deriveVMType(pubKey)
			if err != nil {
				return err
			}
			auth := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					did,
					hex.EncodeToString(pubKey.Bytes()),
					vmType,
				),
				[]string{types.Authentication},
				nil,
			)
			// create the message
			msg := types.NewMsgCreateDidDocument(
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
		Use:     "add-verification-method [id] [pubkey]",
		Short:   "add an verification method to a decentralized did (did) document",
		Example: "adds an verification method for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// signer address
			signer := clientCtx.GetFromAddress()
			// public key
			var pk cryptotypes.PubKey
			err = clientCtx.Codec.UnmarshalInterfaceJSON([]byte(args[1]), &pk)
			if err != nil {
				return err
			}
			// derive the public key type
			vmType, err := deriveVMType(pk)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(clientCtx.ChainID, args[0])
			// verification method id
			vmID := fmt.Sprint(did, "#",
				sdk.MustBech32ifyAddressBytes(
					sdk.GetConfig().GetBech32AccountAddrPrefix(),
					pk.Address().Bytes(),
				),
			)

			verification := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					did,
					hex.EncodeToString(pk.Bytes()),
					vmType,
				),
				[]string{types.Authentication},
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
		Short:   "add a service to a decentralized did (did) document",
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
			did := types.DID(clientCtx.ChainID, args[0])

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
		Use:     "revoke-verification-method [id] [verification-method-id]",
		Short:   "revoke a verification method from a decentralized did (did) document",
		Example: "revoke a verification method for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// verification method id
			vmID := types.DID(clientCtx.ChainID, args[1])
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

// NewDeleteServiceCmd deletes a service from a DID Document
func NewDeleteServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-service [id] [service-id]",
		Short:   "deletes a service from a decentralized did (did) document",
		Example: "delete a service for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(clientCtx.ChainID, args[0])
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

// NewUpdateDidDocumentCmd adds a controller to a did document
func NewUpdateDidDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-did-document [id] [id-key]",
		Short:   "updates a decentralized identifier (did) document to contain a controller",
		Example: "update-did-document vasp cosmos1kslgpxklq75aj96cz3qwsczr95vdtrd3p0fslp",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(clientCtx.ChainID, args[0])

			// did key to use as the controller
			didKey := types.DIDKey(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := types.NewMsgUpdateDidDocument(
				did,
				[]string{
					didKey,
				},
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

// NewSetVerificationRelationshipCmd adds a verification relationship to a verification method
func NewSetVerificationRelationshipCmd() *cobra.Command {

	// relationships
	var relationships []string
	// if true do not add the default authentication relationship
	var unsafe bool

	cmd := &cobra.Command{
		Use:     "set-verification-relationship [id] [method-id] --relationship NAME [--relationship NAME ...]",
		Short:   "adds a verification relationship to a key on a decentralized identifier (did) document.",
		Example: "add-verification-relationship vasp vasp#6f1e0700-6c86-41b6-9e05-ae3cf839cdd0",
		Args:    cobra.ExactArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.DID(clientCtx.ChainID, args[0])

			// method id
			methodID := types.DID(clientCtx.ChainID, args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := types.NewMsgSetVerificationRelationships(
				did,
				methodID,
				relationships,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// make sure that the authentication relationship is preserved
			if !unsafe {
				msg.Relationships = append(msg.Relationships, types.Authentication)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags to set did relationships
	cmd.Flags().StringSliceVarP(&relationships, "relationship", "r", []string{}, "the relationships to set for the verification method in the DID")
	cmd.Flags().BoolVar(&unsafe, "unsafe", false, fmt.Sprint("do not ensure that '", types.Authentication, "' relationship is set"))

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
