package daemon

import (
	"fmt"
	"net"

	"github.com/Manta-Network/manta-fp/eotsmanager"
	"github.com/Manta-Network/manta-fp/eotsmanager/config"
	eotsservice "github.com/Manta-Network/manta-fp/eotsmanager/service"
	"github.com/Manta-Network/manta-fp/log"

	sdkflags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lightningnetwork/lnd/signal"
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the Extractable One Time Signature Daemon",
		Long:  "Start the Extractable One Time Signature Daemon and run it until shutdown.",
		RunE:  startFn,
	}

	cmd.Flags().String(sdkflags.FlagHome, config.DefaultEOTSDir, "The path to the eotsd home directory")
	cmd.Flags().String(rpcListenerFlag, "", "The address that the RPC server listens to")

	return cmd
}

func startFn(cmd *cobra.Command, _ []string) error {
	homePath, err := getHomePath(cmd)
	if err != nil {
		return fmt.Errorf("failed to load home flag: %w", err)
	}

	cfg, err := config.LoadConfig(homePath)
	if err != nil {
		return fmt.Errorf("failed to load config at %s: %w", homePath, err)
	}

	rpcListener, err := cmd.Flags().GetString(rpcListenerFlag)
	if err != nil {
		return fmt.Errorf("failed to get RPC listener flag: %w", err)
	}
	if rpcListener != "" {
		_, err := net.ResolveTCPAddr("tcp", rpcListener)
		if err != nil {
			return fmt.Errorf("invalid RPC listener address %s: %w", rpcListener, err)
		}
		cfg.RPCListener = rpcListener
	}

	logger, err := log.NewRootLoggerWithFile(config.LogFile(homePath), cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to load the logger: %w", err)
	}

	dbBackend, err := cfg.DatabaseConfig.GetDBBackend()
	if err != nil {
		return fmt.Errorf("failed to create db backend: %w", err)
	}

	eotsManager, err := eotsmanager.NewLocalEOTSManager(homePath, cfg.KeyringBackend, dbBackend, logger)
	if err != nil {
		return fmt.Errorf("failed to create EOTS manager: %w", err)
	}

	// Hook interceptor for os signals.
	shutdownInterceptor, err := signal.Intercept()
	if err != nil {
		return fmt.Errorf("failed to set up shutdown interceptor: %w", err)
	}

	eotsServer := eotsservice.NewEOTSManagerServer(cfg, logger, eotsManager, dbBackend, shutdownInterceptor)

	return eotsServer.RunUntilShutdown()
}
