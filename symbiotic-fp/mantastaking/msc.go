package mantastaking

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Manta-Network/manta-fp/bindings"
	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/celestia"
	common2 "github.com/Manta-Network/manta-fp/symbiotic-fp/common"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/config"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/store"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/txmgr"
	types2 "github.com/Manta-Network/manta-fp/types"

	"github.com/lightningnetwork/lnd/kvdb"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type MantaStakingMiddleware struct {
	Ctx                               context.Context
	Cfg                               *MantaStakingMiddlewareConfig
	MantaStakingMiddlewareContract    *bindings.MantaStakingMiddleware
	RawMantaStakingMiddlewareContract *bind.BoundContract
	WalletAddr                        common.Address
	PrivateKey                        *ecdsa.PrivateKey
	MantaStakingMiddlewareABI         *abi.ABI
	txMgr                             txmgr.TxManager
	log                               *zap.Logger
	ChainPoller                       *OpChainPoller
	DAClient                          *celestia.DAClient

	SignatureSubmissionInterval time.Duration
	SubmissionRetryInterval     time.Duration
	MaxSubmissionRetries        uint32

	isStarted *atomic.Bool
	wg        sync.WaitGroup
	quit      chan struct{}
}

func NewMantaStakingMiddleware(mCfg *MantaStakingMiddlewareConfig, config *config.Config, db kvdb.Backend, log *zap.Logger, authToken string) (*MantaStakingMiddleware, error) {
	mantaStakingMiddlewareContract, err := bindings.NewMantaStakingMiddleware(
		mCfg.MantaStakingMiddlewareAddr, mCfg.EthClient,
	)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(
		bindings.MantaStakingMiddlewareMetaData.ABI,
	))
	if err != nil {
		return nil, err
	}
	mantaStakingMiddlewareABI, err := bindings.MantaStakingMiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawTreasureManagerContract := bind.NewBoundContract(
		mCfg.MantaStakingMiddlewareAddr, parsed, mCfg.EthClient, mCfg.EthClient,
		mCfg.EthClient,
	)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       time.Second * 5,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          mCfg.NumConfirmations,
		SafeAbortNonceTooLowCount: mCfg.SafeAbortNonceTooLowCount,
	}

	txMgr := txmgr.NewSimpleTxManager(txManagerConfig, mCfg.EthClient)
	var walletAddr common.Address
	if mCfg.EnableHsm {
		walletAddr = common.HexToAddress(mCfg.HsmAddress)
	} else {
		walletAddr = crypto.PubkeyToAddress(mCfg.PrivateKey.PublicKey)
	}

	fpMetrics := metrics.NewFpMetrics()

	sRStore, err := store.NewOpStateRootStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate op state root store, err: %w", err)
	}

	poller, err := NewOpChainPoller(log, config.OpEventConfig, sRStore, fpMetrics)
	if err != nil {
		return nil, fmt.Errorf("failed to new op chain poller, err: %w", err)
	}

	daClient, err := celestia.NewDAClient(*config.CelestiaConfig, authToken)
	if err != nil {
		return nil, fmt.Errorf("failed to new celestia da client, err: %w", err)
	}

	return &MantaStakingMiddleware{
		Ctx:                               context.Background(),
		Cfg:                               mCfg,
		MantaStakingMiddlewareContract:    mantaStakingMiddlewareContract,
		RawMantaStakingMiddlewareContract: rawTreasureManagerContract,
		WalletAddr:                        walletAddr,
		MantaStakingMiddlewareABI:         mantaStakingMiddlewareABI,
		txMgr:                             txMgr,
		log:                               log,
		ChainPoller:                       poller,
		DAClient:                          daClient,
		PrivateKey:                        mCfg.PrivateKey,
		isStarted:                         atomic.NewBool(false),
		SignatureSubmissionInterval:       config.SignatureSubmissionInterval,
		SubmissionRetryInterval:           config.SubmissionRetryInterval,
	}, nil
}

func (msm *MantaStakingMiddleware) Start() error {
	if msm.isStarted.Swap(true) {
		return fmt.Errorf("the symbiotic-fp %s is already started", msm.WalletAddr.String())
	}
	//todo check the contract function is right
	latestBlock, err := msm.Cfg.EthClient.BlockNumber(msm.Ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest block, err: %v", err)
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        msm.WalletAddr,
	}
	operator, err := msm.MantaStakingMiddlewareContract.Operators(cOpts, msm.WalletAddr)
	if err != nil {
		return fmt.Errorf("failed to get operator info at block: %v, err: %v", latestBlock, err)
	}

	if operator.OperatorName == "" {
		receipt, err := msm.RegisterOperator()
		if err != nil {
			return fmt.Errorf("failed to register operator %w", err)
		}
		msm.log.Info("success to register operator", zap.String("tx_hash", receipt.TxHash.String()))
	} else {
		if operator.Paused {
			msm.log.Error("operator is paused", zap.String("address", msm.WalletAddr.String()))
			return errors.New("operator is paused")
		}
	}

	if err := msm.ChainPoller.Start(); err != nil {
		return fmt.Errorf("failed to start the poller %w", err)
	}

	msm.quit = make(chan struct{})
	msm.wg.Add(1)
	go msm.finalitySigSubmissionLoop()

	return nil
}

func (msm *MantaStakingMiddleware) Stop() error {
	if !msm.isStarted.Swap(false) {
		return fmt.Errorf("the symbiotic-fp %s has already stopped", msm.WalletAddr.String())
	}

	if err := msm.ChainPoller.Stop(); err != nil {
		return fmt.Errorf("failed to stop the poller: %w", err)
	}

	msm.log.Info("stopping symbiotic-fp service", zap.String("address", msm.WalletAddr.String()))

	close(msm.quit)
	msm.wg.Wait()

	msm.log.Info("the symbiotic-fp service is successfully stopped", zap.String("address", msm.WalletAddr.String()))

	return nil
}

func (msm *MantaStakingMiddleware) finalitySigSubmissionLoop() {
	defer msm.wg.Done()

	for {
		select {
		case <-time.After(msm.SignatureSubmissionInterval):
			if err := msm.checkOperatorIsPaused(); err != nil {
				msm.log.Error("the symbiotic-fp failed to check operator is paused",
					zap.String("address", msm.WalletAddr.String()),
					zap.String("error", err.Error()),
				)
			}

			pollerBlocks := msm.getAllBlocksFromChan()
			if len(pollerBlocks) == 0 {
				continue
			}
			targetHeight := pollerBlocks[len(pollerBlocks)-1].Height
			msm.log.Debug("the symbiotic-fp received new block(s), start processing",
				zap.String("address", msm.WalletAddr.String()),
				zap.Uint64("start_height", pollerBlocks[0].Height),
				zap.Uint64("end_height", targetHeight),
			)
			err := msm.retrySubmitSigsUntilFinalized(pollerBlocks)
			if err != nil {
				msm.log.Error("the symbiotic-fp failed to submit signature",
					zap.String("address", msm.WalletAddr.String()),
					zap.String("err", err.Error()))
				continue
			}

			msm.log.Info(
				"successfully submitted the finality signature to the consumer chain",
				zap.String("address", msm.WalletAddr.String()),
				zap.Uint64("start_height", pollerBlocks[0].Height),
				zap.Uint64("end_height", targetHeight),
			)

		case <-msm.quit:
			msm.log.Info("the finality signature submission loop is closing")
			return
		}
	}
}

// retrySubmitSigsUntilFinalized periodically tries to submit finality signature until success or the block is finalized
// error will be returned if maximum retries have been reached or the query to the consumer chain fails
func (msm *MantaStakingMiddleware) retrySubmitSigsUntilFinalized(targetBlocks []*types2.BlockInfo) error {
	if len(targetBlocks) == 0 {
		return fmt.Errorf("cannot send signatures for empty blocks")
	}

	var failedCycles uint32
	targetHeight := targetBlocks[len(targetBlocks)-1].Height

	// we break the for loop if the block is finalized or the signature is successfully submitted
	// error will be returned if maximum retries have been reached or the query to the consumer chain fails
	for {
		select {
		case <-time.After(msm.SubmissionRetryInterval):
			// error will be returned if max retries have been reached
			var err error
			var ctx = context.Background()
			err = msm.SubmitBatchFinalitySignatures(ctx, targetBlocks)
			if err != nil {
				msm.log.Debug(
					"failed to submit finality signature to the consumer chain",
					zap.String("address", msm.WalletAddr.String()),
					zap.Uint32("current_failures", failedCycles),
					zap.Uint64("target_start_height", targetBlocks[0].Height),
					zap.Uint64("target_end_height", targetHeight),
					zap.Error(err),
				)

				failedCycles++
				if failedCycles > msm.MaxSubmissionRetries {
					return fmt.Errorf("reached max failed cycles with err: %w", err)
				}
			} else {
				// the signature has been successfully submitted
				return nil
			}

		case <-msm.quit:
			msm.log.Debug("the symbiotic-fp instance is closing", zap.String("address", msm.WalletAddr.String()))
			return errors.New("the finality provider instance is shutting down")
		}
	}
}

// SubmitBatchFinalitySignatures builds and sends a finality signature over the given block to the consumer chain
// NOTE: the input blocks should be in the ascending order of height
func (msm *MantaStakingMiddleware) SubmitBatchFinalitySignatures(ctx context.Context, blocks []*types2.BlockInfo) error {
	if len(blocks) == 0 {
		return fmt.Errorf("should not submit batch finality signature with zero block")
	}

	if len(blocks) > math.MaxUint32 {
		return fmt.Errorf("should not submit batch finality signature with too many blocks")
	}

	stateRoot := blocks[len(blocks)-1].StateRoot.StateRoot
	signature, err := crypto.Sign(stateRoot[:], msm.PrivateKey)
	if err != nil {
		msm.log.Error("failed to sign data", zap.String("err", err.Error()))
		return err
	}

	signRequest := types2.SignRequest{
		StateRoot:   hex.EncodeToString(stateRoot[:]),
		Signature:   signature,
		SignAddress: msm.WalletAddr.String(),
	}

	data, err := json.Marshal(signRequest)
	if err != nil {
		msm.log.Error("failed to marshal data", zap.String("err", err.Error()))
		return err
	}

	if msm.DAClient != nil && msm.DAClient.Client != nil {
		commit, err := celestia.CreateCommitment(data, msm.DAClient.Namespace)
		if err == nil {
			ctx2, cancel := context.WithTimeout(ctx, msm.DAClient.GetTimeout)
			ids, err := msm.DAClient.Client.Submit(ctx2, [][]byte{data}, -1, msm.DAClient.Namespace)
			cancel()
			if err == nil && len(ids) == 1 && len(ids[0]) == 40 && bytes.Equal(commit, ids[0][8:]) {
				msm.log.Info("celestia: blob successfully submitted", zap.String("id", hex.EncodeToString(ids[0])))
				ctx2, cancel := context.WithTimeout(ctx, msm.DAClient.GetTimeout)
				proofs, err := msm.DAClient.Client.GetProofs(ctx2, ids, msm.DAClient.Namespace)
				cancel()
				if err == nil && len(proofs) == 1 {
					ctx2, cancel := context.WithTimeout(ctx, msm.DAClient.GetTimeout)
					valids, err := msm.DAClient.Client.Validate(ctx2, ids, proofs, msm.DAClient.Namespace)
					cancel()
					if err == nil && len(valids) == 1 && valids[0] == true {
						msm.log.Info("success to send finality signature to celestia")
					} else {
						msm.log.Error("celestia: failed to validate proof",
							zap.String("err", err.Error()),
							zap.Any("valid", valids))
					}
				} else {
					msm.log.Error("celestia: failed to get proof", zap.String("err", err.Error()))
				}
			} else {
				msm.log.Info("celestia: blob submission failed; falling back to eth",
					zap.String("err", err.Error()),
					zap.Any("ids", ids),
					zap.ByteString("commit", commit))
			}
		} else {
			msm.log.Info("celestia: failed to create commitment", zap.String("err", err.Error()))
		}
	}

	return nil
}

func (msm *MantaStakingMiddleware) getAllBlocksFromChan() []*types2.BlockInfo {
	var pollerBlocks []*types2.BlockInfo
	for {
		select {
		case b := <-msm.ChainPoller.GetBlockInfoChan():
			pollerBlocks = append(pollerBlocks, b)
		case <-msm.quit:
			msm.log.Info("the get all blocks loop is closing")
			return nil
		default:
			return pollerBlocks
		}
	}
}

func (msm *MantaStakingMiddleware) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	if !msm.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			msm.Cfg.PrivateKey, msm.Cfg.ChainID,
		)
	} else {
		opts, err = common2.NewHSMTransactOpts(ctx, msm.Cfg.HsmApiName,
			msm.Cfg.HsmAddress, msm.Cfg.ChainID, msm.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true
	finalTx, err := msm.RawMantaStakingMiddlewareContract.RawTransact(opts, tx.Data())
	if err != nil {
		return nil, err
	}
	return finalTx, nil
}

func (msm *MantaStakingMiddleware) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return msm.Cfg.EthClient.SendTransaction(ctx, tx)
}

func (msm *MantaStakingMiddleware) IsMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(
		err.Error(), common2.ErrMaxPriorityFeePerGasNotFound.Error(),
	)
}

func (msm *MantaStakingMiddleware) registerOperator(ctx context.Context) (*types.Transaction, error) {
	balance, err := msm.Cfg.EthClient.BalanceAt(
		msm.Ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, err
	}
	msm.log.Info("manta wallet address balance", zap.String("balance", balance.String()))

	nonce64, err := msm.Cfg.EthClient.NonceAt(
		msm.Ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if !msm.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			msm.Cfg.PrivateKey, msm.Cfg.ChainID,
		)
	} else {
		opts, err = common2.NewHSMTransactOpts(ctx, msm.Cfg.HsmApiName,
			msm.Cfg.HsmAddress, msm.Cfg.ChainID, msm.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true
	tx, err := msm.MantaStakingMiddlewareContract.RegisterOperator(opts, msm.Cfg.OperatorName, common.HexToAddress(msm.Cfg.RewardAddress), big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (msm *MantaStakingMiddleware) RegisterOperator() (*types.Receipt, error) {
	tx, err := msm.registerOperator(msm.Ctx)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return msm.UpdateGasPrice(ctx, tx)
	}
	receipt, err := msm.txMgr.Send(
		msm.Ctx, updateGasPrice, msm.SendTransaction,
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (msm *MantaStakingMiddleware) checkOperatorIsPaused() error {
	latestBlock, err := msm.Cfg.EthClient.BlockNumber(msm.Ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest block, err: %v", err)
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        msm.WalletAddr,
	}
	operator, err := msm.MantaStakingMiddlewareContract.Operators(cOpts, msm.WalletAddr)
	if err != nil {
		return fmt.Errorf("failed to get operator info at block: %v, err: %v", latestBlock, err)
	}

	if operator.Paused {
		msm.log.Info("the operator is paused, stopping service")
		return msm.Stop()
	}

	return nil
}
