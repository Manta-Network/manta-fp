package service

import (
	"context"
	"fmt"
	rollupfpcc "github.com/Manta-Network/manta-fp/bbn-fp/clientcontroller"
	rollupfpcfg "github.com/Manta-Network/manta-fp/bbn-fp/config"
	fpcc "github.com/Manta-Network/manta-fp/clientcontroller"
	"github.com/Manta-Network/manta-fp/ethereum/node"
	"github.com/Manta-Network/manta-fp/finality-provider/service"
	"github.com/Manta-Network/manta-fp/finality-provider/store"
	"github.com/Manta-Network/manta-fp/l2chain/opstack"
	"github.com/Manta-Network/manta-fp/metrics"
	bbntypes "github.com/babylonlabs-io/babylon/v3/types"
	"github.com/lightningnetwork/lnd/kvdb"
	"go.uber.org/zap"
)

// NewRollupBSNFinalityProviderAppFromConfig creates a new FinalityProviderApp instance from the given configuration for rollup BSN.
func NewRollupBSNFinalityProviderAppFromConfig(
	cfg *rollupfpcfg.RollupFPConfig,
	db kvdb.Backend,
	logger *zap.Logger,
	fpPkStr string,
) (*service.FinalityProviderApp, error) {
	cc, err := fpcc.NewBabylonController(cfg.Common.BabylonConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client for the Babylon chain: %w", err)
	}
	if err := cc.Start(); err != nil {
		return nil, fmt.Errorf("failed to start rpc client for the Babylon chain: %w", err)
	}

	consumerCon, err := rollupfpcc.NewRollupBSNController(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client for the consumer chain rollup: %w", err)
	}

	// if the EOTSManagerAddress is empty, run a local EOTS manager;
	// otherwise connect a remote one with a gRPC client
	em, err := service.InitEOTSManagerClient(cfg.Common.EOTSManagerAddress, cfg.Common.HMACKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create EOTS manager client: %w", err)
	}

	logger.Info("successfully connected to a remote EOTS manager", zap.String("address", cfg.Common.EOTSManagerAddress))

	fpMetrics := metrics.NewFpMetrics()

	pubRandStore, err := store.NewPubRandProofStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate public randomness store: %w", err)
	}

	// For rollup environments, always use RollupRandomnessCommitter
	contractConfig, err := consumerCon.QueryContractConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to query contract config: %w", err)
	}

	logger.Info("using RollupRandomnessCommitter for rollup environment",
		zap.Uint64("finality_signature_interval", contractConfig.FinalitySignatureInterval))

	rndCommitter := NewRollupRandomnessCommitter(
		service.NewRandomnessCommitterConfig(cfg.Common.NumPubRand, int64(cfg.Common.TimestampingDelayBlocks), cfg.Common.ContextSigningHeight),
		service.NewPubRandState(pubRandStore),
		consumerCon,
		em,
		logger,
		fpMetrics,
		contractConfig.FinalitySignatureInterval,
	)

	heightDeterminer := service.NewStartHeightDeterminer(consumerCon, cfg.Common.PollerConfig, logger)

	logger.Info("using RollupFinalitySubmitter for rollup environment",
		zap.Uint64("finality_signature_interval", contractConfig.FinalitySignatureInterval))

	// For rollup environments, use RollupFinalitySubmitter for sparse randomness generation
	finalitySubmitter := NewRollupFinalitySubmitter(consumerCon,
		em,
		rndCommitter.GetPubRandProofList,
		service.NewDefaultFinalitySubmitterConfig(cfg.Common.MaxSubmissionRetries,
			cfg.Common.ContextSigningHeight,
			cfg.Common.SubmissionRetryInterval),
		logger,
		fpMetrics,
		contractConfig.FinalitySignatureInterval,
	)

	var fpPk *bbntypes.BIP340PubKey
	if fpPkStr != "" {
		// start the finality-provider instance with the given public key
		fpPk, err = bbntypes.NewBIP340PubKeyFromHex(fpPkStr)
		if err != nil {
			return nil, fmt.Errorf("invalid finality provider public key %s: %w", fpPkStr, err)
		}
	}

	fpStore, err := store.NewFinalityProviderStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate finality provider store: %w", err)
	}
	sfp, err := fpStore.GetFinalityProvider(fpPk.MustToBTCPK())
	if err != nil {
		return nil, fmt.Errorf("failed to get finality provider: %w", err)
	}

	sRStore, err := store.NewOpStateRootStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate op state root store: %w", err)
	}

	opClient, err := node.DialEthClient(context.Background(), cfg.OpEventConfig.EthRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to create op client: %w", err)
	}

	ep, err := opstack.NewEventProvider(context.Background(), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate op event provider: %w", err)
	}

	startHeight, err := heightDeterminer.DetermineStartHeight(context.Background(), fpPk, func() (uint64, error) {
		return sfp.LastVotedHeight, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to determine start height: %w", err)
	}

	poller, err := service.NewOpChainPoller(logger, opClient, startHeight, cfg.OpEventConfig, sRStore, ep, fpMetrics)
	if err != nil {
		return nil, fmt.Errorf("failed to new op chain poller: %w", err)
	}

	fpApp, err := service.NewBsnFinalityProviderApp(cfg.Common, cc, consumerCon, em, poller, rndCommitter, heightDeterminer, finalitySubmitter, fpMetrics, db, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create finality provider app: %w", err)
	}

	return fpApp, nil
}
