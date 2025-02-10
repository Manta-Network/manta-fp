package main

import (
	"fmt"
	"os"

	incentivecli "github.com/babylonlabs-io/babylon/x/incentive/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	fpcmd "github.com/Manta-Network/manta-fp/bbn-fp/cmd"
	"github.com/Manta-Network/manta-fp/bbn-fp/cmd/bfpd/daemon"
	fpcfg "github.com/Manta-Network/manta-fp/bbn-fp/config"
	"github.com/Manta-Network/manta-fp/version"
)

// NewRootCmd creates a new root command for bfpd. It is called once in the main function.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "bfpd",
		Short:             "bfpd - Finality Provider Daemon (bfpd).",
		Long:              `bfpd is the daemon to create and manage finality providers.`,
		SilenceErrors:     false,
		PersistentPreRunE: fpcmd.PersistClientCtx(client.Context{}),
	}
	rootCmd.PersistentFlags().String(flags.FlagHome, fpcfg.DefaultFpdDir, "The application home directory")

	return rootCmd
}

func main() {
	cmd := NewRootCmd()
	cmd.AddCommand(
		daemon.CommandInit(), daemon.CommandStart(), daemon.CommandKeys(), daemon.CommandAddEotsKey(),
		daemon.CommandGetDaemonInfo(), daemon.CommandCreateFP(), daemon.CommandLsFP(),
		daemon.CommandInfoFP(), daemon.CommandAddFinalitySig(), daemon.CommandUnjailFP(),
		daemon.CommandEditFinalityDescription(), daemon.CommandCommitPubRand(),
		incentivecli.NewWithdrawRewardCmd(),
		version.CommandVersion("bfpd"),
	)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your bfpd CLI '%s'", err)
		os.Exit(1)
	}
}
