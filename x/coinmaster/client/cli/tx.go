package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/noria-net/noria/x/coinmaster/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Interact with the Coinmaster to mint and burn coins.",
		Long:                       `Coinmaster allows minting and burning of whitelisted native coins, by whitelisted minters. The amount denomination may be restricted by a whiteist too. These parameters may be adjusted through governance proposals.`,
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdMint())
	cmd.AddCommand(CmdBurn())

	return cmd
}
