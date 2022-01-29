package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	// "strings"
	"github.com/spf13/cobra"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/allinbits/cosmos-cash/v3/x/regulator/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group regulator queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	return cmd
}
