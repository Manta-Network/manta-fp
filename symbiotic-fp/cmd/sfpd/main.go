package main

import (
	"fmt"
	"os"

	"github.com/Manta-Network/manta-fp/symbiotic-fp/cmd/sfpd/daemon"
	fpcfg "github.com/Manta-Network/manta-fp/symbiotic-fp/config"
	"github.com/Manta-Network/manta-fp/version"

	"github.com/spf13/cobra"
)

// NewRootCmd creates a new root command for bfpd. It is called once in the main function.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "sfpd",
		Short:         "sfpd - Finality Provider Daemon (sfpd).",
		Long:          `sfpd is the daemon to create and manage finality providers.`,
		SilenceErrors: false,
	}
	rootCmd.PersistentFlags().String(daemon.HomeFlag, fpcfg.DefaultFpdDir, "The application home directory")

	return rootCmd
}

func main() {
	cmd := NewRootCmd()
	cmd.AddCommand(
		daemon.CommandInit(), daemon.CommandStart(),
		version.CommandVersion("sfpd"),
	)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your sfpd CLI '%s'", err)
		os.Exit(1)
	}
}
