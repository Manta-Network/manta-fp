package mantastaking

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
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
	kmssigner "github.com/Manta-Network/manta-fp/symbiotic-fp/kms"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/store"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/txmgr"
	types2 "github.com/Manta-Network/manta-fp/types"

	"github.com/lightningnetwork/lnd/kvdb"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
)

type MantaStakingMiddleware struct {
	Ctx                                  context.Context
	Cfg                                  *MantaStakingMiddlewareConfig
	MantaStakingMiddlewareContract       *bindings.MantaStakingMiddleware
	RawMantaStakingMiddlewareContract    *bind.BoundContract
	SymbioticOperatorRegisterContract    *bindings.SymbioticOperatorRegister
	RawSymbioticOperatorRegisterContract *bind.BoundContract
	WalletAddr                           common.Address
	PrivateKey                           *ecdsa.PrivateKey
	txMgr                                txmgr.TxManager
	log                                  *zap.Logger
	ChainPoller                          *OpChainPoller
	DAClient                             *celestia.DAClient
	metrics                              *metrics.SymbioticFpMetrics

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
	mParsed, err := abi.JSON(strings.NewReader(
		bindings.MantaStakingMiddlewareMetaData.ABI,
	))
	if err != nil {
		return nil, err
	}
	rawMantaStakingMiddlewareContract := bind.NewBoundContract(
		mCfg.MantaStakingMiddlewareAddr, mParsed, mCfg.EthClient, mCfg.EthClient,
		mCfg.EthClient,
	)

	symbioticOperatorRegisterContract, err := bindings.NewSymbioticOperatorRegister(
		mCfg.SymbioticOperatorRegisterAddr, mCfg.EthClient,
	)
	if err != nil {
		return nil, err
	}
	sParsed, err := abi.JSON(strings.NewReader(
		bindings.SymbioticOperatorRegisterMetaData.ABI,
	))
	if err != nil {
		return nil, err
	}
	rawSymbioticOperatorRegisterContract := bind.NewBoundContract(
		mCfg.SymbioticOperatorRegisterAddr, sParsed, mCfg.EthClient, mCfg.EthClient,
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
	} else if config.EnableKms {
		walletAddr, err = kmssigner.GetAddress(mCfg.KmsClient, mCfg.KmsID)
		if err != nil {
			return nil, fmt.Errorf("failed to get the kms address: %w", err)
		}
	} else {
		walletAddr = crypto.PubkeyToAddress(mCfg.PrivateKey.PublicKey)
	}

	fpMetrics := metrics.NewSymbioticFpMetrics()

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
		Ctx:                                  context.Background(),
		Cfg:                                  mCfg,
		MantaStakingMiddlewareContract:       mantaStakingMiddlewareContract,
		RawMantaStakingMiddlewareContract:    rawMantaStakingMiddlewareContract,
		SymbioticOperatorRegisterContract:    symbioticOperatorRegisterContract,
		RawSymbioticOperatorRegisterContract: rawSymbioticOperatorRegisterContract,
		WalletAddr:                           walletAddr,
		txMgr:                                txMgr,
		log:                                  log,
		ChainPoller:                          poller,
		DAClient:                             daClient,
		PrivateKey:                           mCfg.PrivateKey,
		isStarted:                            atomic.NewBool(false),
		SignatureSubmissionInterval:          config.SignatureSubmissionInterval,
		SubmissionRetryInterval:              config.SubmissionRetryInterval,
	}, nil
}

func (msm *MantaStakingMiddleware) Start() error {
	if msm.isStarted.Swap(true) {
		return fmt.Errorf("the symbiotic-fp %s is already started", msm.WalletAddr.String())
	}
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
		msm.log.Info(fmt.Sprintf("address %s is not operator, start register", msm.WalletAddr.String()))

		isExist, err := msm.MantaStakingMiddlewareContract.OperatorNameExists(cOpts, Keccak256Hash([]byte(msm.Cfg.OperatorName)))
		if err != nil {
			return fmt.Errorf("failed to get operator operator name isExists: %v, err: %v", latestBlock, err)
		}
		if isExist {
			return fmt.Errorf("operator name %s is exist, please choose a new name", msm.Cfg.OperatorName)
		}

		isEntity, err := msm.SymbioticOperatorRegisterContract.IsEntity(cOpts, msm.WalletAddr)
		if err != nil {
			return fmt.Errorf("failed to get symbiotic operator info, err:%w", err)
		}
		if !isEntity {
			receipt, err := msm.RegisterSymbioticOperator()
			if err != nil {
				return fmt.Errorf("failed to register symbiotic operator %w", err)
			}
			msm.log.Info("success to register symbiotic operator", zap.String("tx_hash", receipt.TxHash.String()))
		}
		receipt, err := msm.RegisterOperator()
		if err != nil {
			return fmt.Errorf("failed to register manta staking operator %w", err)
		}
		msm.log.Info("success to register manta staking operator", zap.String("tx_hash", receipt.TxHash.String()))
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
				continue
			}

			isActive, err := msm.checkOperatorIsDelegateActive()
			if err != nil {
				msm.log.Error("the symbiotic-fp failed to check operator is delegate active",
					zap.String("address", msm.WalletAddr.String()),
					zap.String("error", err.Error()),
				)
				continue
			}

			if !isActive {
				msm.log.Warn("the symbiotic-fp is not delegate active, skip sign")
				msm.metrics.RecordFpStatus(msm.WalletAddr.String(), common2.InActive)
				continue
			}

			msm.metrics.RecordFpStatus(msm.WalletAddr.String(), common2.Active)

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
			err = msm.retrySubmitSigsUntilFinalized(pollerBlocks)
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
				msm.metrics.RecordFpLastProcessedHeight(msm.WalletAddr.String(), targetHeight)
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

	stateRoot := blocks[len(blocks)-1].StateRoot
	signature, err := crypto.Sign(stateRoot.StateRoot[:], msm.PrivateKey)
	if err != nil {
		msm.log.Error("failed to sign data", zap.String("err", err.Error()))
		return err
	}

	signRequest := types2.SignRequest{
		StateRoot:     hex.EncodeToString(stateRoot.StateRoot[:]),
		L1BlockNumber: stateRoot.L1BlockNumber,
		L1BlockHash:   stateRoot.L1BlockHash.String(),
		L2BlockNumber: stateRoot.L2BlockNumber.Uint64(),
		Signature:     signature,
		SignAddress:   msm.WalletAddr.String(),
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
			}
		} else {
			msm.log.Info("celestia: failed to create commitment", zap.String("err", err.Error()))
			msm.metrics.IncrementFpTotalFailedVotes(msm.WalletAddr.String())
		}
	}

	msm.metrics.RecordFpLastVotedL1Height(msm.WalletAddr.String(), stateRoot.L1BlockNumber)
	msm.metrics.RecordFpLastVotedL2Height(msm.WalletAddr.String(), stateRoot.L2BlockNumber.Uint64())
	msm.metrics.IncrementFpTotalVotedBlocks(msm.WalletAddr.String())

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

func (msm *MantaStakingMiddleware) UpdateMantaStakingGasPrice(opts *bind.TransactOpts, tx *types.Transaction) (*types.Transaction, error) {
	finalTx, err := msm.RawMantaStakingMiddlewareContract.RawTransact(opts, tx.Data())
	if err != nil {
		return nil, err
	}
	return finalTx, nil
}

func (msm *MantaStakingMiddleware) UpdateSymbioticGasPrice(opts *bind.TransactOpts, tx *types.Transaction) (*types.Transaction, error) {
	finalTx, err := msm.RawSymbioticOperatorRegisterContract.RawTransact(opts, tx.Data())
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

func (msm *MantaStakingMiddleware) registerSymbioticOperator(ctx context.Context) (*types.Transaction, *bind.TransactOpts, error) {
	balance, err := msm.Cfg.EthClient.BalanceAt(
		ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, nil, err
	}
	msm.log.Info("manta wallet address balance", zap.String("balance", balance.String()))

	nonce64, err := msm.Cfg.EthClient.NonceAt(
		ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if msm.Cfg.EnableHsm {
		opts, err = common2.NewHSMTransactOpts(ctx, msm.Cfg.HsmApiName,
			msm.Cfg.HsmAddress, msm.Cfg.ChainID, msm.Cfg.HsmCreden)
	} else if msm.Cfg.EnableKms {
		opts, err = kmssigner.NewAwsKmsTransactorWithChainIDCtx(ctx, msm.Cfg.KmsClient,
			msm.Cfg.KmsID, msm.Cfg.ChainID)
	} else {
		opts, err = bind.NewKeyedTransactorWithChainID(
			msm.Cfg.PrivateKey, msm.Cfg.ChainID,
		)
	}
	if err != nil {
		return nil, nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true
	tx, err := msm.SymbioticOperatorRegisterContract.RegisterOperator(opts)
	if err != nil {
		return nil, nil, err
	}
	return tx, opts, nil
}

func (msm *MantaStakingMiddleware) RegisterSymbioticOperator() (*types.Receipt, error) {
	ctx := context.Background()
	tx, opts, err := msm.registerSymbioticOperator(ctx)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return msm.UpdateSymbioticGasPrice(opts, tx)
	}
	receipt, err := msm.txMgr.Send(
		ctx, updateGasPrice, msm.SendTransaction,
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (msm *MantaStakingMiddleware) registerOperator(ctx context.Context) (*types.Transaction, *bind.TransactOpts, error) {
	balance, err := msm.Cfg.EthClient.BalanceAt(
		ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, nil, err
	}
	msm.log.Info("manta wallet address balance", zap.String("balance", balance.String()))

	nonce64, err := msm.Cfg.EthClient.NonceAt(
		ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if msm.Cfg.EnableHsm {
		opts, err = common2.NewHSMTransactOpts(ctx, msm.Cfg.HsmApiName,
			msm.Cfg.HsmAddress, msm.Cfg.ChainID, msm.Cfg.HsmCreden)
	} else if msm.Cfg.EnableKms {
		opts, err = kmssigner.NewAwsKmsTransactorWithChainIDCtx(ctx, msm.Cfg.KmsClient,
			msm.Cfg.KmsID, msm.Cfg.ChainID)
	} else {
		opts, err = bind.NewKeyedTransactorWithChainID(
			msm.Cfg.PrivateKey, msm.Cfg.ChainID,
		)
	}
	if err != nil {
		return nil, nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	xBytes := msm.PrivateKey.PublicKey.X.Bytes()
	yBytes := msm.PrivateKey.PublicKey.Y.Bytes()
	paddedX := make([]byte, 32)
	copy(paddedX[32-len(xBytes):], xBytes)
	paddedY := make([]byte, 32)
	copy(paddedY[32-len(yBytes):], yBytes)
	publicKeyBytes := append(paddedX, paddedY...)

	tx, err := msm.MantaStakingMiddlewareContract.RegisterOperator(opts, publicKeyBytes, msm.Cfg.OperatorName, common.HexToAddress(msm.Cfg.RewardAddress), big.NewInt(msm.Cfg.Commission))
	if err != nil {
		return nil, nil, err
	}
	return tx, opts, nil
}

func (msm *MantaStakingMiddleware) RegisterOperator() (*types.Receipt, error) {
	ctx := context.Background()
	tx, opts, err := msm.registerOperator(ctx)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return msm.UpdateMantaStakingGasPrice(opts, tx)
	}
	receipt, err := msm.txMgr.Send(
		ctx, updateGasPrice, msm.SendTransaction,
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
		msm.log.Error("the operator is paused")
		msm.metrics.RecordFpStatus(msm.WalletAddr.String(), common2.Paused)
		return fmt.Errorf("the operator is paused at block: %v, address: %v", latestBlock, msm.WalletAddr.String())
	}

	return nil
}

func (msm *MantaStakingMiddleware) checkOperatorIsDelegateActive() (bool, error) {
	stakeAmount, err := msm.getSymbioticOperatorStakeAmount(strings.ToLower(msm.WalletAddr.String()))
	if err != nil {
		msm.log.Error("failed to get operator stake amount", zap.String("address", msm.WalletAddr.String()), zap.Error(err))
		return false, err
	}

	stakeLimit, _ := new(big.Int).SetString(msm.Cfg.StakeLimit, 10)
	if stakeAmount.Cmp(stakeLimit) < 0 {
		msm.log.Error("the total stake amount is insufficient", zap.String("staked", stakeAmount.String()), zap.String("required", msm.Cfg.StakeLimit))
		return false, nil
	}

	return true, nil
}

func (msm *MantaStakingMiddleware) getSymbioticOperatorStakeAmount(operator string) (*big.Int, error) {
	query := fmt.Sprintf(`{"query":"query {\n  vaultUpdates(first: 1, where: {operator: \"%s\"}, orderBy: timestamp, orderDirection: desc) {\n    vaultTotalActiveStaked\n  }\n}"}`, operator)
	jsonQuery := []byte(query)

	req, err := http.NewRequest("POST", msm.Cfg.SymbioticStakeUrl, bytes.NewBuffer(jsonQuery))
	if err != nil {
		msm.log.Error("Error creating HTTP request:", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, multipart/mixed")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		msm.log.Error("Error sending HTTP request:", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		msm.log.Error("Error reading response body:", zap.Error(err))
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		msm.log.Error("Error parsing JSON response:", zap.Error(err))
		return nil, err
	}

	var totalStaked = big.NewInt(0)
	if data, exists := result["data"]; exists {
		if vaultUpdates, exists := data.(map[string]interface{})["vaultUpdates"]; exists {
			if len(vaultUpdates.([]interface{})) > 0 {
				vaultTotalActiveStaked := vaultUpdates.([]interface{})[0].(map[string]interface{})["vaultTotalActiveStaked"]
				totalStaked, _ = new(big.Int).SetString(vaultTotalActiveStaked.(string), 10)
				msm.log.Info(fmt.Sprintf("operator %s vaultTotalActiveStaked: %s", operator, vaultTotalActiveStaked))
			} else {
				msm.log.Warn(fmt.Sprintf("operator %s no vault updates found", operator))
			}
		} else {
			msm.log.Warn(fmt.Sprintf("operator %s no vaultUpdates field found in response data", operator))
		}
	} else {
		msm.log.Warn(fmt.Sprintf("operator %s no data field found in JSON response", operator))
	}

	return totalStaked, nil
}

func Keccak256Hash(data []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	var result [32]byte
	hash.Sum(result[:0])
	return result
}
