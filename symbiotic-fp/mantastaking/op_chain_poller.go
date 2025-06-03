package mantastaking

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Manta-Network/manta-fp/ethereum/node"
	"github.com/Manta-Network/manta-fp/l2chain/opstack"
	"github.com/Manta-Network/manta-fp/metrics"
	cfg "github.com/Manta-Network/manta-fp/symbiotic-fp/config"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/store"
	"github.com/Manta-Network/manta-fp/types"

	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type OpChainPoller struct {
	isStarted *atomic.Bool
	wg        sync.WaitGroup
	logger    *zap.Logger
	sRStore   *store.OpStateRootStore

	opClient       node.EthClient
	cfg            *cfg.OpEventConfig
	contracts      []common.Address
	headers        []ctypes.Header
	latestBlock    *big.Int
	blockTraversal *node.BlockTraversal
	eventProvider  *opstack.EventProvider
	blockInfoChan  chan *types.BlockInfo

	metrics *metrics.SymbioticFpMetrics
	quit    chan struct{}
}

func NewOpChainPoller(logger *zap.Logger, cfg *cfg.OpEventConfig, sRStore *store.OpStateRootStore, metrics *metrics.SymbioticFpMetrics) (*OpChainPoller, error) {
	var contracts []common.Address
	contracts = append(contracts, common.HexToAddress(cfg.L2OutputOracleAddr))

	opClient, err := node.DialEthClient(context.Background(), cfg.EthRpc)
	if err != nil {
		logger.Error("failed to dial eth client", zap.String("err", err.Error()))
		return nil, err
	}

	dbLatestBlock, err := sRStore.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	var fromBlock *big.Int
	if dbLatestBlock != nil {
		logger.Info("sync detected last indexed block", zap.String("blockNumber", dbLatestBlock.String()))
		fromBlock = dbLatestBlock
	} else if cfg.StartHeight > 0 {
		logger.Info("no sync indexed state starting from supplied ethereum height", zap.Uint64("height", cfg.StartHeight))
		header, err := opClient.BlockHeaderByNumber(big.NewInt(int64(cfg.StartHeight)))
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromBlock = header.Number
	} else {
		logger.Info("no ethereum block indexed state")
	}

	blockTraversal := node.NewBlockTraversal(opClient, fromBlock, big.NewInt(0), cfg.ChainId, logger)

	eventProvider, err := opstack.NewEventProvider(context.Background(), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate op event provider: %w", err)
	}

	return &OpChainPoller{
		isStarted:      atomic.NewBool(false),
		logger:         logger,
		opClient:       opClient,
		sRStore:        sRStore,
		cfg:            cfg,
		latestBlock:    fromBlock,
		blockTraversal: blockTraversal,
		eventProvider:  eventProvider,
		blockInfoChan:  make(chan *types.BlockInfo, cfg.BufferSize),
		metrics:        metrics,
		contracts:      contracts,
		quit:           make(chan struct{}),
	}, nil
}

func (ocp *OpChainPoller) Start() error {
	if ocp.isStarted.Swap(true) {
		return fmt.Errorf("the op chain poller is already started")
	}

	ocp.logger.Info("starting the op chain poller")

	ocp.wg.Add(1)
	go ocp.opPollChain()

	ocp.metrics.RecordPollerStartingHeight(ocp.latestBlock.Uint64())
	ocp.logger.Info("the chain poller is successfully started")

	return nil
}

func (ocp *OpChainPoller) Stop() error {
	if !ocp.isStarted.Swap(false) {
		return fmt.Errorf("the op chain poller has already stopped")
	}

	ocp.logger.Info("stopping the op chain poller")

	close(ocp.quit)
	ocp.wg.Wait()

	ocp.logger.Info("the op chain poller is successfully stopped")

	return nil
}

func (ocp *OpChainPoller) GetBlockInfoChan() <-chan *types.BlockInfo {
	return ocp.blockInfoChan
}

func (ocp *OpChainPoller) opPollChain() {
	defer ocp.wg.Done()
	for {
		select {
		case <-time.After(ocp.cfg.PollInterval):
			if len(ocp.headers) > 0 {
				ocp.logger.Info("retrying previous batch")
			} else {
				newHeaders, err := ocp.blockTraversal.NextHeaders(ocp.cfg.BlockStep)
				if err != nil {
					ocp.logger.Error("error querying for headers", zap.String("err", err.Error()))
					continue
				} else if len(newHeaders) == 0 {
					ocp.logger.Warn("no new headers. syncer at head?")
				} else {
					ocp.headers = newHeaders
				}
				latestBlock := ocp.blockTraversal.LatestBlock()
				if latestBlock != nil {
					ocp.logger.Info("Latest header", zap.String("latestHeader Number", latestBlock.String()))
				}
				err = ocp.sRStore.AddLatestBlock(latestBlock)
				if err != nil {
					ocp.logger.Error("Add latest block fail", zap.String("err", err.Error()))
					return
				}
				ocp.metrics.RecordLastPolledHeight(latestBlock.Uint64())
			}
			err := ocp.processBatch(ocp.headers)
			if err == nil {
				ocp.headers = nil
			}
		case <-ocp.quit:
			return
		}
	}
}

func (ocp *OpChainPoller) processBatch(headers []ctypes.Header) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	ocp.logger.Info("extracting batch", zap.Int("size", len(headers)), zap.String("startBlock", firstHeader.Number.String()), zap.String("endBlock", lastHeader.Number.String()))
	headerMap := make(map[common.Hash]*ctypes.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}

	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: ocp.contracts}
	logs, err := ocp.opClient.FilterLogs(filterQuery)
	if err != nil {
		ocp.logger.Error("failed to extract logs", zap.String("err", err.Error()))
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash")
	}

	if len(logs.Logs) > 0 {
		ocp.logger.Info("detected logs", zap.Int("size", len(logs.Logs)))
	}

	blockList := make([]types.Block, 0, len(headers))
	for i := range headers {
		if headers[i].Number == nil {
			continue
		}
		blockItem := types.Block{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].ParentHash,
			Number:     headers[i].Number,
			Timestamp:  headers[i].Time,
		}
		blockList = append(blockList, blockItem)
	}

	// contracts parse
	for i, log := range logs.Logs {
		if log.Topics[0].String() == ocp.eventProvider.L2ooABI.Events["OutputProposed"].ID.String() {
			stateRootEvent, err := ocp.eventProvider.ProcessStateRootEvent(logs.Logs[i])
			if err != nil {
				return err
			}
			ocp.logger.Info("event list", zap.String("stateroot", hex.EncodeToString(stateRootEvent.StateRoot[:])))

			err = ocp.sRStore.SaveStateRoot(big.NewInt(int64(stateRootEvent.L1BlockNumber)), stateRootEvent.StateRoot,
				stateRootEvent.L2BlockNumber, stateRootEvent.L1BlockHash, stateRootEvent.L2OutputIndex, stateRootEvent.DisputeGameType)
			if err != nil {
				ocp.logger.Error("failed to store state root", zap.String("err", err.Error()))
				return err
			}
			ocp.blockInfoChan <- &types.BlockInfo{
				Height:    logs.Logs[i].BlockNumber,
				Hash:      logs.Logs[i].BlockHash.Bytes(),
				Finalized: false,
				StateRoot: *stateRootEvent,
			}
		} else {
			return nil
		}
	}
	return nil
}
