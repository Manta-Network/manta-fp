package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	fpcfg "github.com/Manta-Network/manta-fp/bbn-fp/config"
	"github.com/Manta-Network/manta-fp/bbn-fp/proto"
	"github.com/Manta-Network/manta-fp/bbn-fp/store"
	"github.com/Manta-Network/manta-fp/clientcontroller"
	"github.com/Manta-Network/manta-fp/eotsmanager"
	"github.com/Manta-Network/manta-fp/eotsmanager/client"
	"github.com/Manta-Network/manta-fp/ethereum/node"
	fpkr "github.com/Manta-Network/manta-fp/keyring"
	"github.com/Manta-Network/manta-fp/l2chain/opstack"
	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/types"

	sdkmath "cosmossdk.io/math"
	"github.com/avast/retry-go/v4"
	bbntypes "github.com/babylonlabs-io/babylon/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/lightningnetwork/lnd/kvdb"
	"go.uber.org/zap"
)

type FinalityProviderApp struct {
	startOnce sync.Once
	stopOnce  sync.Once
	wg        sync.WaitGroup
	quit      chan struct{}

	cc           clientcontroller.ClientController
	kr           keyring.Keyring
	fps          *store.FinalityProviderStore
	pubRandStore *store.PubRandProofStore
	config       *fpcfg.Config
	logger       *zap.Logger
	input        *strings.Reader

	opClient node.EthClient
	sRStore  *store.OpStateRootStore
	eP       *opstack.EventProvider

	fpIns       *FinalityProviderInstance
	eotsManager eotsmanager.EOTSManager

	metrics *metrics.FpMetrics

	createFinalityProviderRequestChan chan *CreateFinalityProviderRequest
	unjailFinalityProviderRequestChan chan *UnjailFinalityProviderRequest
	criticalErrChan                   chan *CriticalError
}

func NewFinalityProviderAppFromConfig(
	cfg *fpcfg.Config,
	db kvdb.Backend,
	logger *zap.Logger,
) (*FinalityProviderApp, error) {
	cc, err := clientcontroller.NewClientController(cfg.ChainType, cfg.BabylonConfig, cfg.OpEventConfig, &cfg.BTCNetParams, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client for the consumer chain %s: %w", cfg.ChainType, err)
	}

	// if the EOTSManagerAddress is empty, run a local EOTS manager;
	// otherwise connect a remote one with a gRPC client
	em, err := client.NewEOTSManagerGRpcClient(cfg.EOTSManagerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to create EOTS manager client: %w", err)
	}

	logger.Info("successfully connected to a remote EOTS manager", zap.String("address", cfg.EOTSManagerAddress))
	return NewFinalityProviderApp(cfg, cc, em, db, logger)
}

func NewFinalityProviderApp(
	config *fpcfg.Config,
	cc clientcontroller.ClientController,
	em eotsmanager.EOTSManager,
	db kvdb.Backend,
	logger *zap.Logger,
) (*FinalityProviderApp, error) {
	fpStore, err := store.NewFinalityProviderStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate finality provider store: %w", err)
	}
	pubRandStore, err := store.NewPubRandProofStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate public randomness store: %w", err)
	}
	sRStore, err := store.NewOpStateRootStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate op state root store: %w", err)
	}

	input := strings.NewReader("")
	kr, err := fpkr.CreateKeyring(
		config.BabylonConfig.KeyDirectory,
		config.BabylonConfig.ChainID,
		config.BabylonConfig.KeyringBackend,
		input,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyring: %w", err)
	}

	fpMetrics := metrics.NewFpMetrics()

	opClient, err := node.DialEthClient(context.Background(), config.OpEventConfig.EthRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to create op client: %w", err)
	}

	ep, err := opstack.NewEventProvider(context.Background(), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate op event provider: %w", err)
	}

	return &FinalityProviderApp{
		cc:                                cc,
		fps:                               fpStore,
		pubRandStore:                      pubRandStore,
		kr:                                kr,
		config:                            config,
		logger:                            logger,
		input:                             input,
		fpIns:                             nil,
		eotsManager:                       em,
		opClient:                          opClient,
		sRStore:                           sRStore,
		eP:                                ep,
		metrics:                           fpMetrics,
		quit:                              make(chan struct{}),
		unjailFinalityProviderRequestChan: make(chan *UnjailFinalityProviderRequest),
		createFinalityProviderRequestChan: make(chan *CreateFinalityProviderRequest),
		criticalErrChan:                   make(chan *CriticalError),
	}, nil
}

func (app *FinalityProviderApp) GetConfig() *fpcfg.Config {
	return app.config
}

func (app *FinalityProviderApp) GetFinalityProviderStore() *store.FinalityProviderStore {
	return app.fps
}

func (app *FinalityProviderApp) GetPubRandProofStore() *store.PubRandProofStore {
	return app.pubRandStore
}

func (app *FinalityProviderApp) GetFinalityProviderInfo(fpPk *bbntypes.BIP340PubKey) (*proto.FinalityProviderInfo, error) {
	storedFp, err := app.fps.GetFinalityProvider(fpPk.MustToBTCPK())
	if err != nil {
		return nil, err
	}

	fpInfo := storedFp.ToFinalityProviderInfo()

	if app.IsFinalityProviderRunning(fpPk) {
		fpInfo.IsRunning = true
	}

	return fpInfo, nil
}

func (app *FinalityProviderApp) ListAllFinalityProvidersInfo() ([]*proto.FinalityProviderInfo, error) {
	storedFps, err := app.fps.GetAllStoredFinalityProviders()
	if err != nil {
		return nil, err
	}

	fpsInfo := make([]*proto.FinalityProviderInfo, 0, len(storedFps))
	for _, fp := range storedFps {
		fpInfo := fp.ToFinalityProviderInfo()

		if app.IsFinalityProviderRunning(fp.GetBIP340BTCPK()) {
			fpInfo.IsRunning = true
		}

		fpsInfo = append(fpsInfo, fpInfo)
	}

	return fpsInfo, nil
}

// GetFinalityProviderInstance returns the bbn-fp instance with the given Babylon public key
func (app *FinalityProviderApp) GetFinalityProviderInstance() (*FinalityProviderInstance, error) {
	if app.fpIns == nil {
		return nil, fmt.Errorf("finality provider does not exist")
	}

	return app.fpIns, nil
}

// StartFinalityProvider starts a finality provider instance with the given EOTS public key
// Note: this should be called right after the bbn-fp is registered
func (app *FinalityProviderApp) StartFinalityProvider(fpPk *bbntypes.BIP340PubKey, passphrase string) error {
	app.logger.Info("starting finality provider", zap.String("pk", fpPk.MarshalHex()))

	if err := app.startFinalityProviderInstance(fpPk, passphrase); err != nil {
		return err
	}

	app.logger.Info("finality provider is started", zap.String("pk", fpPk.MarshalHex()))

	return nil
}

// SyncAllFinalityProvidersStatus syncs the status of all the stored finality providers with the chain.
// it should be called before a fp instance is started
func (app *FinalityProviderApp) SyncAllFinalityProvidersStatus() error {
	fps, err := app.fps.GetAllStoredFinalityProviders()
	if err != nil {
		return err
	}

	for _, fp := range fps {
		latestBlock, err := app.cc.QueryBestBlock()
		if err != nil {
			return err
		}

		pkHex := fp.GetBIP340BTCPK().MarshalHex()
		hasPower, err := app.cc.QueryFinalityProviderVotingPower(fp.BtcPk, latestBlock.Height)
		if err != nil {
			return fmt.Errorf("failed to query voting power for finality provider %s at height %d: %w",
				fp.GetBIP340BTCPK().MarshalHex(), latestBlock.Height, err)
		}

		// power > 0 (slashed_height must > 0), set status to ACTIVE
		oldStatus := fp.Status
		if hasPower {
			if oldStatus != proto.FinalityProviderStatus_ACTIVE {
				fp.Status = proto.FinalityProviderStatus_ACTIVE
				app.fps.MustSetFpStatus(fp.BtcPk, proto.FinalityProviderStatus_ACTIVE)
				app.logger.Debug(
					"the bbn-fp status is changed to ACTIVE",
					zap.String("fp_btc_pk", pkHex),
					zap.String("old_status", oldStatus.String()),
					zap.Bool("has_power", hasPower),
				)
			}
			continue
		}
		slashed, jailed, err := app.cc.QueryFinalityProviderSlashedOrJailed(fp.BtcPk)
		if err != nil {
			return err
		}
		if slashed {
			app.fps.MustSetFpStatus(fp.BtcPk, proto.FinalityProviderStatus_SLASHED)

			app.logger.Debug(
				"the bbn-fp status is changed to SLAHED",
				zap.String("fp_btc_pk", pkHex),
				zap.String("old_status", oldStatus.String()),
			)

			continue
		}
		if jailed {
			app.fps.MustSetFpStatus(fp.BtcPk, proto.FinalityProviderStatus_JAILED)

			app.logger.Debug(
				"the bbn-fp status is changed to JAILED",
				zap.String("fp_btc_pk", pkHex),
				zap.String("old_status", oldStatus.String()),
			)

			continue
		}
		// power == 0 and slashed_height == 0, change to INACTIVE if the current status is ACTIVE
		if oldStatus == proto.FinalityProviderStatus_ACTIVE {
			app.fps.MustSetFpStatus(fp.BtcPk, proto.FinalityProviderStatus_INACTIVE)

			app.logger.Debug(
				"the bbn-fp status is changed to INACTIVE",
				zap.String("fp_btc_pk", pkHex),
				zap.String("old_status", oldStatus.String()),
			)

			continue
		}
	}

	return nil
}

// Start starts only the bbn-fp daemon without any bbn-fp instances
func (app *FinalityProviderApp) Start() error {
	var startErr error
	app.startOnce.Do(func() {
		app.logger.Info("Starting FinalityProviderApp")

		startErr = app.SyncAllFinalityProvidersStatus()
		if startErr != nil {
			return
		}

		app.wg.Add(5)
		go app.metricsUpdateLoop()
		go app.monitorCriticalErr()
		go app.monitorStatusUpdate()
		go app.registrationLoop()
		go app.unjailFpLoop()
	})

	return startErr
}

func (app *FinalityProviderApp) Stop() error {
	var stopErr error
	app.stopOnce.Do(func() {
		app.logger.Info("Stopping FinalityProviderApp")

		close(app.quit)
		app.wg.Wait()

		if app.fpIns != nil && app.fpIns.IsRunning() {
			pkHex := app.fpIns.GetBtcPkHex()
			app.logger.Info("stopping finality provider", zap.String("pk", pkHex))

			if err := app.fpIns.Stop(); err != nil {
				stopErr = fmt.Errorf("failed to close the fp instance: %w", err)
				return
			}

			app.logger.Info("finality provider is stopped", zap.String("pk", pkHex))
		}

		app.logger.Debug("Stopping EOTS manager")
		if err := app.eotsManager.Close(); err != nil {
			stopErr = fmt.Errorf("failed to close the EOTS manager: %w", err)
			return
		}

		app.logger.Debug("FinalityProviderApp successfully stopped")
	})
	return stopErr
}

func (app *FinalityProviderApp) CreateFinalityProvider(
	keyName, chainID, passPhrase string,
	eotsPk *bbntypes.BIP340PubKey,
	description *stakingtypes.Description,
	commission *sdkmath.LegacyDec,
) (*CreateFinalityProviderResult, error) {
	// 1. check if the chain key exists
	kr, err := fpkr.NewChainKeyringControllerWithKeyring(app.kr, keyName, app.input)
	if err != nil {
		return nil, err
	}

	fpAddr, err := kr.Address(passPhrase)
	if err != nil {
		// the chain key does not exist, should create the chain key first
		return nil, fmt.Errorf("the keyname %s does not exist, add the key first: %w", keyName, err)
	}

	// 2. create proof-of-possession
	if eotsPk == nil {
		return nil, fmt.Errorf("eots pk cannot be nil")
	}
	pop, err := app.CreatePop(fpAddr, eotsPk, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("failed to create proof-of-possession of the bbn-fp: %w", err)
	}

	// TODO: query consumer chain to check if the fp is already registered
	// if true, update db with the fp info from the consumer chain
	// otherwise, proceed registration

	// 3. register the finality provider on the consumer chain
	request := &CreateFinalityProviderRequest{
		fpAddr:          fpAddr,
		btcPubKey:       eotsPk,
		pop:             pop,
		description:     description,
		commission:      commission,
		errResponse:     make(chan error, 1),
		successResponse: make(chan *RegisterFinalityProviderResponse, 1),
	}

	app.createFinalityProviderRequestChan <- request

	select {
	case err := <-request.errResponse:
		return nil, err
	case successResponse := <-request.successResponse:
		pkHex := eotsPk.MarshalHex()
		btcPk := eotsPk.MustToBTCPK()
		// save the fp info to db after successful registration
		// this ensures the data saved in db is consistent with that on the consumer chain
		// if the program crashes in the middle, the user can retry registration
		// which will update db use the information from the consumer chain without
		// submitting a registration again
		if err := app.fps.CreateFinalityProvider(fpAddr, btcPk, description, commission, chainID); err != nil {
			return nil, fmt.Errorf("failed to save bbn-fp: %w", err)
		}

		app.metrics.RecordFpStatus(pkHex, proto.FinalityProviderStatus_REGISTERED)

		app.logger.Info("successfully saved the bbn-fp",
			zap.String("eots_pk", pkHex),
			zap.String("addr", fpAddr.String()),
		)

		storedFp, err := app.fps.GetFinalityProvider(btcPk)
		if err != nil {
			return nil, err
		}

		return &CreateFinalityProviderResult{
			FpInfo: storedFp.ToFinalityProviderInfo(),
			TxHash: successResponse.txHash,
		}, nil
	case <-app.quit:
		return nil, fmt.Errorf("bbn-fp app is shutting down")
	}
}

// UnjailFinalityProvider sends a transaction to unjail a bbn-fp
func (app *FinalityProviderApp) UnjailFinalityProvider(fpPk *bbntypes.BIP340PubKey) (*UnjailFinalityProviderResponse, error) {
	// send request to the loop to avoid blocking the main thread
	request := &UnjailFinalityProviderRequest{
		btcPubKey:       fpPk,
		errResponse:     make(chan error, 1),
		successResponse: make(chan *UnjailFinalityProviderResponse, 1),
	}

	app.unjailFinalityProviderRequestChan <- request

	select {
	case err := <-request.errResponse:
		return nil, err
	case successResponse := <-request.successResponse:
		_, err := app.fps.GetFinalityProvider(fpPk.MustToBTCPK())
		if err != nil {
			return nil, fmt.Errorf("failed to get finality provider from db: %w", err)
		}

		// Update bbn-fp status in the local store
		// set it to INACTIVE for now and it will be updated to
		// ACTIVE if the fp has voting power
		err = app.fps.SetFpStatus(fpPk.MustToBTCPK(), proto.FinalityProviderStatus_INACTIVE)
		if err != nil {
			return nil, fmt.Errorf("failed to update bbn-fp status after unjailing: %w", err)
		}

		app.metrics.RecordFpStatus(fpPk.MarshalHex(), proto.FinalityProviderStatus_INACTIVE)

		return successResponse, nil
	case <-app.quit:
		return nil, fmt.Errorf("bbn-fp app is shutting down")
	}
}

func (app *FinalityProviderApp) CreatePop(fpAddress sdk.AccAddress, fpPk *bbntypes.BIP340PubKey, passphrase string) (*bstypes.ProofOfPossessionBTC, error) {
	pop := &bstypes.ProofOfPossessionBTC{
		BtcSigType: bstypes.BTCSigType_BIP340, // by default, we use BIP-340 encoding for BTC signature
	}

	// generate pop.BtcSig = schnorr_sign(sk_BTC, hash(bbnAddress))
	// NOTE: *schnorr.Sign has to take the hash of the message.
	// So we have to hash the address before signing
	hash := tmhash.Sum(fpAddress.Bytes())

	sig, err := app.eotsManager.SignSchnorrSig(fpPk.MustMarshal(), hash, passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to get schnorr signature from the EOTS manager: %w", err)
	}

	pop.BtcSig = bbntypes.NewBIP340SignatureFromBTCSig(sig).MustMarshal()

	return pop, nil
}

func (app *FinalityProviderApp) startFinalityProviderInstance(
	pk *bbntypes.BIP340PubKey,
	passphrase string,
) error {
	pkHex := pk.MarshalHex()
	if app.fpIns == nil {
		fpIns, err := NewFinalityProviderInstance(
			pk, app.config, app.fps, app.pubRandStore, app.cc, app.eotsManager,
			app.metrics, passphrase, app.criticalErrChan, app.logger, app.opClient,
			app.sRStore, app.eP,
		)
		if err != nil {
			return fmt.Errorf("failed to create finality provider instance %s: %w", pkHex, err)
		}

		app.fpIns = fpIns
	} else if !pk.Equals(app.fpIns.btcPk) {
		return fmt.Errorf("the finality provider daemon is already bonded with the finality provider %s,"+
			"please restart the daemon to switch to another instance", app.fpIns.btcPk.MarshalHex())
	}

	return app.fpIns.Start()
}

func (app *FinalityProviderApp) IsFinalityProviderRunning(fpPk *bbntypes.BIP340PubKey) bool {
	if app.fpIns == nil {
		return false
	}

	if app.fpIns.GetBtcPkHex() != fpPk.MarshalHex() {
		return false
	}

	return app.fpIns.IsRunning()
}

func (app *FinalityProviderApp) removeFinalityProviderInstance() error {
	fpi := app.fpIns
	if fpi == nil {
		return fmt.Errorf("the finality provider instance does not exist")
	}
	if fpi.IsRunning() {
		if err := fpi.Stop(); err != nil {
			return fmt.Errorf("failed to stop the finality provider instance %s", fpi.GetBtcPkHex())
		}
	}

	app.fpIns = nil

	return nil
}

func (app *FinalityProviderApp) setFinalityProviderSlashed(fpi *FinalityProviderInstance) {
	fpi.MustSetStatus(proto.FinalityProviderStatus_SLASHED)
	if err := app.removeFinalityProviderInstance(); err != nil {
		panic(fmt.Errorf("failed to terminate a slashed bbn-fp %s: %w", fpi.GetBtcPkHex(), err))
	}
}

func (app *FinalityProviderApp) setFinalityProviderJailed(fpi *FinalityProviderInstance) {
	fpi.MustSetStatus(proto.FinalityProviderStatus_JAILED)
	if err := app.removeFinalityProviderInstance(); err != nil {
		panic(fmt.Errorf("failed to terminate a jailed bbn-fp %s: %w", fpi.GetBtcPkHex(), err))
	}
}

func (app *FinalityProviderApp) getLatestBlockWithRetry() (*types.BlockInfo, error) {
	var (
		latestBlock *types.BlockInfo
		err         error
	)

	if err := retry.Do(func() error {
		latestBlock, err = app.cc.QueryBestBlock()
		if err != nil {
			return err
		}
		return nil
	}, RtyAtt, RtyDel, RtyErr, retry.OnRetry(func(n uint, err error) {
		app.logger.Debug(
			"failed to query the consumer chain for the latest block",
			zap.Uint("attempt", n+1),
			zap.Uint("max_attempts", RtyAttNum),
			zap.Error(err),
		)
	})); err != nil {
		return nil, err
	}

	return latestBlock, nil
}

// NOTE: this is not safe in production, so only used for testing purpose
func (app *FinalityProviderApp) getFpPrivKey(fpPk []byte) (*btcec.PrivateKey, error) {
	record, err := app.eotsManager.KeyRecord(fpPk, "")
	if err != nil {
		return nil, err
	}

	return record.PrivKey, nil
}
