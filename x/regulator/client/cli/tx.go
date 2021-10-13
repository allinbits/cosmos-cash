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
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

var (
	// DefaultRelativePacketTimeoutTimestamp default timeout
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
	activateRegulatorDID                  string
	activateRegulatorCredentialID         string
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
			// assign a did
			did := didtypes.NewKeyDID(signer.String())
			if activateRegulatorDID != "" {
				did = didtypes.DID(activateRegulatorDID)
			}

			// assign a credential id
			credID := fmt.Sprint("regulator-credential/", did)
			if activateRegulatorCredentialID != "" {
				credID = activateRegulatorCredentialID
			}
			// credentials
			vc := vctypes.NewRegulatorVerifiableCredential(
				credID,
				did.String(),
				time.Now().UTC(),
				vctypes.NewRegulatorCredentialSubject(
					did.String(),
					name,
					countryCode,
				),
			)
			// signer is the vmID
			vmID := did.NewVerificationMethodID(signer.String())

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

	cmd.Flags().StringVar(&activateRegulatorDID, "did", "", "the DID id to use for the regulator DID, otherwise the adddress of the regulator will be used")
	cmd.Flags().StringVar(&activateRegulatorCredentialID, "credential-id", "", "the credential id to use for the regulator credential, randomly generated if not present")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// IssueLicenseCredentialCmd defines the command to create a new license verifiable credential.
// This is used by ulatorsreg to define issuers and issuer permissions
func IssueLicenseCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `issue-license-credential [cred_id] [issuer_did] [credential_subject_did] [type] [country] [authority] [denom] [circulation_limit]`,
		Short:   "issues a license credential",
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
			issuerDid := didtypes.DID(args[1])
			credentialSubject := didtypes.DID(args[2])
			licenseType := args[3]
			country := args[4]
			authority := args[5]
			denom := args[6]
			circulationLimitString := args[7]
			circulationLimit, _ := sdk.NewIntFromString(circulationLimitString)
			coin := sdk.NewCoin(denom, circulationLimit)

			cs := vctypes.NewLicenseCredentialSubject(
				credentialSubject.String(),
				licenseType,
				country,
				authority,
				coin,
			)
			tm := time.Now()

			vc := vctypes.NewLicenseVerifiableCredential(
				credentialID,
				issuerDid.String(),
				tm.UTC(),
				cs,
			)

			vmID := issuerDid.NewVerificationMethodID(accAddr.String())
			signedVc, err := vc.Sign(clientCtx.Keyring, accAddr, vmID)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueLicenseCredential(signedVc, accAddrBech32)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// IssueRegistrationCredentialCmd defines the command to create a new license verifiable credential.
// This is used by regulators to define issuers and issuer permissions
func IssueRegistrationCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `issue-registration-credential [cred_id] [issuer_did] [credential_subject_did] [country] [short_name] [long_name]`,
		Short:   "issue a registration credential for a DID",
		Example: "creates a registration verifiable credential for e-money issuers",
		Args:    cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()
			accAddrBech32 := accAddr.String()

			credentialID := args[0]
			issuerDid := didtypes.DID(args[1])
			credentialSubject := didtypes.DID(args[2])
			country := args[3]
			shortName := args[4]
			longName := args[5]

			cs := vctypes.NewRegistrationCredentialSubject(
				credentialSubject.String(),
				country,
				shortName,
				longName,
			)
			tm := time.Now()

			vc := vctypes.NewRegistrationVerifiableCredential(
				credentialID,
				issuerDid.String(),
				tm.UTC(),
				cs,
			)

			vmID := issuerDid.NewVerificationMethodID(accAddr.String())
			signedVc, err := vc.Sign(clientCtx.Keyring, accAddr, vmID)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueRegistrationCredential(signedVc, accAddrBech32)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
