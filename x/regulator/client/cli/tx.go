package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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

// const (
// 	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
// )

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
	cmd.AddCommand(CmdActivate())

	return cmd
}

var _ = strconv.Itoa(0)

func CmdActivate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate [name] [countryCode]",
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
			credId := fmt.Sprint("regulator-credential/", did)
			if activateRegulatorCredentialID != "" {
				credId = activateRegulatorCredentialID
			}
			// credentials
			vc := vctypes.NewRegulatorVerifiableCredential(
				credId,
				did.String(),
				time.Now().UTC(),
				vctypes.NewRegulatorCredentialSubject(
					credId,
					name,
					countryCode,
				),
			)
			// sign the credentials
			signedVc := vc.Sign(clientCtx.Keyring, signer, did.String())
			// compose the message
			msg := types.NewMsgActivate(
				signedVc,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVar(&activateRegulatorDID, "did", "", "the DID id to use for the regulator DID, otherwise the adddress of the regulator will be used")
	cmd.Flags().StringVar(&activateRegulatorCredentialID, "credential-id", "", "the credential id to use for the regulator credential, randomly generated if not present")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
