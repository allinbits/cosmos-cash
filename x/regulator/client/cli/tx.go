package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	didtypes "github.com/allinbits/cosmos-cash/v2/x/did/types"
	"github.com/allinbits/cosmos-cash/v2/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/v2/x/verifiable-credential/types"
)

var (
	// DefaultRelativePacketTimeoutTimestamp default timeout
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
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
	cmd.AddCommand(ActivateCmd())
	cmd.AddCommand(IssueLicenseCredentialCmd())
	cmd.AddCommand(IssueRegistrationCredentialCmd())

	return cmd
}

var _ = strconv.Itoa(0)

func ActivateCmd() *cobra.Command {

	var (
		issuerDIDStr  string
		subjectDIDStr string
		credentialID  string
	)

	cmd := &cobra.Command{
		Use:   "activate-regulator-credential [name] [countryCode]",
		Short: "Broadcast message to activate a regulator did",
		Long: `Regulators addresses are stored in the regulator genesis parameters, 
a did for each regulator is generated at genesis time but is not active, that is, 
a regulator must activate its DID document via this command.
The command will trigger the generation of a new verifiable credential for the regulator 
that activates it.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			signer := clientCtx.GetFromAddress()
			// parameters
			name := args[0]
			countryCode := args[1]
			// get the issuer did
			issuerDID := didtypes.NewKeyDID(signer.String())
			if issuerDIDStr != "" {
				issuerDID = didtypes.DID(issuerDIDStr)
			}
			// get the subject did
			subjectDID := didtypes.NewKeyDID(signer.String())
			if subjectDIDStr != "" {
				subjectDID = didtypes.DID(subjectDIDStr)
			}

			// assign a credential id if not set
			if credentialID == "" {
				credentialID = fmt.Sprint("regulator-credential/", subjectDID)
			}
			// credentials
			vc := vctypes.NewRegulatorVerifiableCredential(
				credentialID,
				issuerDID.String(),
				time.Now().UTC(),
				vctypes.NewRegulatorCredentialSubject(
					subjectDID.String(),
					name,
					countryCode,
				),
			)
			// signer is the vmID
			vmID := issuerDID.NewVerificationMethodID(signer.String())

			// sign the credentials
			signedVc, err := vc.Sign(clientCtx.Keyring, signer, vmID)
			if err != nil {
				return err
			}

			// compose the message
			msg := types.NewMsgIssueRegulatorCredential(
				signedVc,
				signer.String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVar(&issuerDIDStr, "issuer-did", "", "the DID id to use for the the issuer of the regulator credential, defaults to the key did from the signing address")
	cmd.Flags().StringVar(&subjectDIDStr, "subject-did", "", "the DID id to use for the the subject of the regulator credential, defaults to the key did from the signing address")
	cmd.Flags().StringVar(&credentialID, "credential-id", "", "the credential identifier, automatically assigned if not provided")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// IssueLicenseCredentialCmd defines the command to create a new license verifiable credential.
// This is used by ulatorsreg to define issuers and issuer permissions
func IssueLicenseCredentialCmd() *cobra.Command {

	var (
		credentialID string
	)

	cmd := &cobra.Command{
		Use:   `issue-license-credential [issuer_did] [subject_did] [type] [country] [authority] [denom] [circulation_limit]`,
		Short: "issues a license credential",
		Example: `cosmos-cashd tx regulator issue-license-credential \ 
did:cosmos:net:cosmoscash-testnet:regulator \
did:cosmos:net:cosmoscash-testnet:emti \
MICAEMI IRL "Another Financial Services Body (AFFB)" sEUR 10000" `,
		Args: cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			issuerDID := didtypes.DID(args[0])
			subjectDID := didtypes.DID(args[1])
			licenseType := args[2]
			country := args[3]
			authority := args[4]
			denom := args[5]
			circulationLimitString := args[6]
			circulationLimit, _ := sdk.NewIntFromString(circulationLimitString)
			coin := sdk.NewCoin(denom, circulationLimit)

			// assign a credential id if not set
			if credentialID == "" {
				credentialID = fmt.Sprintf("license/%v/%v", denom, subjectDID)
			}

			vc := vctypes.NewLicenseVerifiableCredential(
				credentialID,
				issuerDID.String(),
				time.Now().UTC(),
				vctypes.NewLicenseCredentialSubject(
					subjectDID.String(),
					licenseType,
					country,
					authority,
					coin,
				),
			)

			vmID := issuerDID.NewVerificationMethodID(accAddr.String())
			signedVc, err := vc.Sign(clientCtx.Keyring, accAddr, vmID)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueLicenseCredential(signedVc, accAddrBech32)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVar(&credentialID, "credential-id", "", "the credential identifier, automatically assigned if not provided")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// IssueRegistrationCredentialCmd defines the command to create a new license verifiable credential.
// This is used by regulators to define issuers and issuer permissions
func IssueRegistrationCredentialCmd() *cobra.Command {

	var (
		credentialID string
	)

	cmd := &cobra.Command{
		Use:   `issue-registration-credential [issuer_did] [subject_did] [country] [short_name] [long_name]`,
		Short: "issue a registration credential for a DID",
		Example: `cosmos-cashd tx issue-registration-credential \
did:cosmos:net:cosmoscash-testnet:regulator \ 
did:cosmos:net:cosmoscash-testnet:emti \
EU EmoneyONE "EmoneyONE GmbH" `,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			issuerDID := didtypes.DID(args[0])
			subjectDID := didtypes.DID(args[1])
			country := args[2]
			shortName := args[3]
			longName := args[4]

			// assign a credential id if not set
			if credentialID == "" {
				credentialID = fmt.Sprintf("registration/%s", subjectDID)
			}

			vc := vctypes.NewRegistrationVerifiableCredential(
				credentialID,
				issuerDID.String(),
				time.Now().UTC(),
				vctypes.NewRegistrationCredentialSubject(
					subjectDID.String(),
					country,
					shortName,
					longName,
				),
			)

			vmID := issuerDID.NewVerificationMethodID(accAddr.String())
			signedVc, err := vc.Sign(clientCtx.Keyring, accAddr, vmID)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueRegistrationCredential(signedVc, accAddrBech32)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVar(&credentialID, "credential-id", "", "the credential identifier, automatically assigned if not provided")

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
