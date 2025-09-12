package daemon

import (
	"fmt"
	"path/filepath"

	rollupfpcfg "github.com/Manta-Network/manta-fp/bbn-fp/config"
	clientctx "github.com/Manta-Network/manta-fp/finality-provider/cmd/fpd/clientctx"
	fpdaemon "github.com/Manta-Network/manta-fp/finality-provider/cmd/fpd/daemon"
	"github.com/Manta-Network/manta-fp/util"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func CommandRecoverProof(binaryName string) *cobra.Command {
	cmd := fpdaemon.CommandRecoverProofTemplate(binaryName)
	cmd.RunE = clientctx.RunEWithClientCtx(runCommandRecoverProof)

	return cmd
}

func runCommandRecoverProof(ctx client.Context, cmd *cobra.Command, args []string) error {
	// Get homePath from context like in start.go
	homePath, err := filepath.Abs(ctx.HomeDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	homePath = util.CleanAndExpandPath(homePath)

	cfg, err := rollupfpcfg.LoadConfig(homePath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if err := fpdaemon.RunCommandRecoverProofWithConfig(ctx, cmd, homePath, cfg.Common, args); err != nil {
		return fmt.Errorf("failed to run recover proof command: %w", err)
	}

	return nil
}
