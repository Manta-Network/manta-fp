package daemon

import (
	"fmt"
	"path/filepath"

	"github.com/Manta-Network/manta-fp/log"
	fpcfg "github.com/Manta-Network/manta-fp/symbiotic-fp/config"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/mantastaking"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/service"
	"github.com/Manta-Network/manta-fp/util"

	"github.com/lightningnetwork/lnd/signal"
	"github.com/spf13/cobra"
)

// CommandStart returns the start command of bfpd daemon.
func CommandStart() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "start",
		Short:   "Start the symbiotic-fp app daemon.",
		Long:    `Start the symbiotic-fp app. Note that privateKey should be started beforehand`,
		Example: `sfpd start --home /home/user/.sfpd --private-key abcd1234 `,
		Args:    cobra.NoArgs,
		RunE:    runStartCmd,
	}
	cmd.Flags().String(PrivateKeyFlag, "", "The private key of the symbiotic-fp to sign")
	return cmd
}

func runStartCmd(cmd *cobra.Command, _ []string) error {
	home, err := cmd.Flags().GetString(HomeFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", HomeFlag, err)
	}
	priKey, err := cmd.Flags().GetString(PrivateKeyFlag)
	if err != nil {
		return fmt.Errorf("failed to read flag %s: %w", PrivateKeyFlag, err)
	}
	homePath, err := filepath.Abs(home)
	if err != nil {
		return err
	}
	homePath = util.CleanAndExpandPath(homePath)

	cfg, err := fpcfg.LoadConfig(homePath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	logger, err := log.NewRootLoggerWithFile(fpcfg.LogFile(homePath), cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to initialize the logger: %w", err)
	}

	dbBackend, err := cfg.DatabaseConfig.GetDBBackend()
	if err != nil {
		return fmt.Errorf("failed to create db backend: %w", err)
	}

	// Hook interceptor for os signals.
	shutdownInterceptor, err := signal.Intercept()
	if err != nil {
		return err
	}

	server := service.NewFinalityProviderServer(cfg, logger, dbBackend, shutdownInterceptor)

	mSMCfg, err := mantastaking.NewMantaStakingMiddlewareConfig(cmd.Context(), cfg, logger, priKey)
	if err != nil {
		return fmt.Errorf("failed to initialize the manta staking middleware config: %w", err)
	}
	mantaStakeServer, err := mantastaking.NewMantaStakingMiddleware(mSMCfg, cfg, dbBackend, logger)
	if err != nil {
		return fmt.Errorf("failed to initialize the manta staking middleware: %w", err)
	}
	err = mantaStakeServer.Start()
	if err != nil {
		return fmt.Errorf("failed to start the manta staking service: %w", err)
	}

	return server.RunUntilShutdown()
}
