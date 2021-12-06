package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/wealdtech/go-merkletree"

	didtypes "github.com/allinbits/cosmos-cash/v2/x/did/types"
	"github.com/allinbits/cosmos-cash/v2/x/issuer/types"
	vctypes "github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
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
		NewIssueUserVerifiableCredentialCmd(),
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

// NewIssueUserVerifiableCredentialCmd defines the command to create a new verifiable credential.
func NewIssueUserVerifiableCredentialCmd() *cobra.Command {

	var credentialID string

	cmd := &cobra.Command{
		Use:   `issue-user-credential [issuer_did] [subject_did] [secret] [amount_per_transaction] [total_number_of_transactions] [total_transaction_amount]`,
		Short: "create decentralized verifiable-credential",
		Example: `cosmos-cashd tx issuer issue-user-credential \
did:cosmos:net:cash:emti did:cosmos:cred:emti-user-alice zkp_secret 1000 1000 1000 \
--credential-id emti-alice-proof-of-kyc \
--from emti --chain-id cash -y`,
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			issuerDID := didtypes.DID(args[0])
			subjectDID := didtypes.DID(args[1])

			// assign a credential id if not set
			if credentialID == "" {
				credentialID = fmt.Sprintf("PoKYC/%s", subjectDID)
			}

			secret := args[2]
			inputs := args[3:6]

			data := make([][]byte, len(inputs))
			for i, v := range inputs {
				data[i] = []byte(v)
			}

			tree, err := merkletree.NewUsing(data, vctypes.New(secret), nil)
			if err != nil {
				return err
			}

			root := tree.Root()
			hexRoot := hex.EncodeToString(root)

			vc := vctypes.NewUserVerifiableCredential(
				credentialID,
				issuerDID.String(),
				time.Now(),
				vctypes.NewUserCredentialSubject(
					subjectDID.String(),
					hexRoot,
					true,
				),
			)

			vmID := issuerDID.NewVerificationMethodID(accAddrBech32)

			signedVc, err := vc.Sign(clientCtx.Keyring, accAddr, vmID)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueUserCredential(
				signedVc,
				accAddrBech32,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVar(&credentialID, "credential-id", "", "the credential identifier, automatically assigned if not provided")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewBurnTokenCmd defines the command to burn tokens.
func NewBurnTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-token [issuer_did] [license_cred_id] [amount]",
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
		Use:   "mint-token [issuer_did] [license_cred_id] [amount]",
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
		Use:   "pause-token [issuer_did] [license_cred_id]",
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
