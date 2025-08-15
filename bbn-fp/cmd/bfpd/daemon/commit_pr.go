package daemon

import (
	"context"
	"fmt"
	"math"
	"path/filepath"
	"strconv"

	fpcfg "github.com/Manta-Network/manta-fp/bbn-fp/config"
	"github.com/Manta-Network/manta-fp/bbn-fp/service"
	"github.com/Manta-Network/manta-fp/bbn-fp/store"
	fpcc "github.com/Manta-Network/manta-fp/clientcontroller"
	eotsclient "github.com/Manta-Network/manta-fp/eotsmanager/client"
	"github.com/Manta-Network/manta-fp/ethereum/node"
	"github.com/Manta-Network/manta-fp/l2chain/opstack"
	"github.com/Manta-Network/manta-fp/log"
	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/util"

	bbntypes "github.com/babylonlabs-io/babylon/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// CommandCommitPubRand returns the commit-pubrand command
func CommandCommitPubRand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unsafe-commit-pubrand [fp-eots-pk-hex] [target-height]",
		Aliases: []string{"unsafe-cpr"},
		Short:   "[UNSAFE] Manually trigger public randomness commitment for a finality provider",
		Long: `[UNSAFE] Manually trigger public randomness commitment for a finality provider.
WARNING: this can drain the finality provider's balance if the target height is too high.`,
		Example: `bfpd unsafe-commit-pubrand --home /home/user/.bfpd [fp-eots-pk-hex] [target-height]`,
		Args:    cobra.ExactArgs(2),
		RunE:    runCommandCommitPubRand,
	}
	cmd.Flags().Uint64("start-height", math.MaxUint64, "The block height to start committing pubrand from (optional)")
	return cmd
}

func runCommandCommitPubRand(cmd *cobra.Command, args []string) error {
	fpPk, err := bbntypes.NewBIP340PubKeyFromHex(args[0])
	if err != nil {
		return err
	}
	targetHeight, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return err
	}
	startHeight, err := cmd.Flags().GetUint64("start-height")
	if err != nil {
		return err
	}

	// Get homePath from context like in start.go
	clientCtx := client.GetClientContextFromCmd(cmd)
	homePath, err := filepath.Abs(clientCtx.HomeDir)
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

	db, err := cfg.DatabaseConfig.GetDBBackend()
	if err != nil {
		return fmt.Errorf("failed to create db backend: %w", err)
	}

	fpStore, err := store.NewFinalityProviderStore(db)
	if err != nil {
		return fmt.Errorf("failed to initiate finality provider store: %w", err)
	}
	pubRandStore, err := store.NewPubRandProofStore(db)
	if err != nil {
		return fmt.Errorf("failed to initiate public randomness store: %w", err)
	}
	cc, err := fpcc.NewClientController(cfg.ChainType, cfg.BabylonConfig, cfg.OpEventConfig, &cfg.BTCNetParams, logger)
	if err != nil {
		return fmt.Errorf("failed to create rpc client for the Babylon chain: %w", err)
	}
	em, err := eotsclient.NewEOTSManagerGRpcClient(cfg.EOTSManagerAddress)
	if err != nil {
		return fmt.Errorf("failed to create EOTS manager client: %w", err)
	}
	opClient, err := node.DialEthClient(context.Background(), cfg.OpEventConfig.EthRpc)
	if err != nil {
		return fmt.Errorf("failed to create op client: %w", err)
	}
	sRStore, err := store.NewOpStateRootStore(db)
	if err != nil {
		return fmt.Errorf("failed to initiate op state root store: %w", err)
	}

	ep, err := opstack.NewEventProvider(context.Background(), logger)
	if err != nil {
		return fmt.Errorf("failed to initiate op event provider: %w", err)
	}

	fp, err := service.NewFinalityProviderInstance(
		fpPk, cfg, fpStore, pubRandStore, cc, em, metrics.NewBbnFpMetrics(), "",
		make(chan<- *service.CriticalError), logger, opClient, sRStore, ep)
	if err != nil {
		return fmt.Errorf("failed to create bbn-fp %s instance: %w", fpPk.MarshalHex(), err)
	}

	if startHeight == math.MaxUint64 {
		return fp.TestCommitPubRand(targetHeight)
	}
	return fp.TestCommitPubRandWithStartHeight(startHeight, targetHeight)
}
