package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		NewAddAuthenticationCmd(),
	)

	return cmd
}

// NewCreateIdentifierCmd defines the command to create a new IBC light client.
func NewCreateIdentifierCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-identifier [id]",
		Short:   "create decentralized identifier (did) document",
		Example: fmt.Sprintf("creates a did document for users"),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			//cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)
			accAddr := clientCtx.GetFromAddress()

			info, err := clientCtx.Keyring.KeyByAddress(accAddr)
			if err != nil {
				return err
			}
			pubKey := info.GetPubKey()
			accAddrBech32 := accAddr.String()
			id := types.DidPrefix + accAddrBech32

			auth := types.NewAuthentication(
				id+"#keys-1",
				"sepk256",
				accAddrBech32,
				pubKey.Address().String(),
			)

			msg := types.NewMsgCreateIdentifier(
				id,
				types.Authentications{&auth},
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

func NewAddAuthenticationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-authentication [id] [pubkey]",
		Short:   "add an authentication method to a decentralized identifier (did) document",
		Example: fmt.Sprintf("adds an authentication method for a did document"),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			accAddr := clientCtx.GetFromAddress()

			pubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, args[1])
			address := sdk.AccAddress(pubKey.Address())

			auth := types.NewAuthentication(
				"",
				"sepk256",
				address.String(),
				pubKey.Address().String(),
			)

			msg := types.NewMsgAddAuthentication(
				args[0],
				&auth,
				accAddr.String(),
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
