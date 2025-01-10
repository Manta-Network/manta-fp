package service

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp/ethereum/node"
	cfg "github.com/Manta-Network/manta-fp/finality-provider/config"
	"github.com/Manta-Network/manta-fp/finality-provider/store"
	"github.com/Manta-Network/manta-fp/l2chain/opstack"
	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/types"
)

type OpChainPoller struct {
	isStarted *atomic.Bool
	wg        sync.WaitGroup
	logger    *zap.Logger
	stateRoot *store.OpStateRootStore

	opClient       node.EthClient
	cfg            *cfg.OpEventConfig
	headers        []ctypes.Header
	latestBlock    *big.Int
	blockTraversal *node.BlockTraversal
	eventProvider  *opstack.EventProvider
	blockInfoChan  chan *types.BlockInfo

	metrics *metrics.FpMetrics
	quit    chan struct{}
}

func NewOpChainPoller(logger *zap.Logger, opClient node.EthClient, cfg *cfg.OpEventConfig, stateRoot *store.OpStateRootStore, eventProvider *opstack.EventProvider, metrics *metrics.FpMetrics) (*OpChainPoller, error) {
	dbLatestBlock, err := stateRoot.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	var fromBlock *big.Int
	if dbLatestBlock != nil {
		log.Info("sync detected last indexed block", "blockNumber", dbLatestBlock)
		fromBlock = dbLatestBlock
	} else if cfg.ScanStartHeight > 0 {
		log.Info("no sync indexed state starting from supplied ethereum height", "height", cfg.ScanStartHeight)
		header, err := opClient.BlockHeaderByNumber(big.NewInt(int64(cfg.ScanStartHeight)))
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromBlock = header.Number
	} else {
		log.Info("no ethereum block indexed state")
	}

	blockTraversal := node.NewBlockTraversal(opClient, fromBlock, big.NewInt(0), cfg.ChainId)

	return &OpChainPoller{
		isStarted:      atomic.NewBool(false),
		logger:         logger,
		opClient:       opClient,
		cfg:            cfg,
		latestBlock:    fromBlock,
		blockTraversal: blockTraversal,
		eventProvider:  eventProvider,
		blockInfoChan:  make(chan *types.BlockInfo, cfg.BufferSize),
		metrics:        metrics,
		quit:           make(chan struct{}),
	}, nil
}

func (ocp *OpChainPoller) Start(startHeight uint64) error {
	if ocp.isStarted.Swap(true) {
		return fmt.Errorf("the op chain poller is already started")
	}

	ocp.logger.Info("starting the op chain poller")

	go ocp.opPollChain()

	ocp.metrics.RecordPollerStartingHeight(startHeight) // todo: change to op stack
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
				log.Info("retrying previous batch")
			} else {
				newHeaders, err := ocp.blockTraversal.NextHeaders(uint64(ocp.cfg.BlockStep))
				if err != nil {
					log.Error("error querying for headers", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					log.Warn("no new headers. syncer at head?")
				} else {
					ocp.headers = newHeaders
				}
				latestBlock := ocp.blockTraversal.LatestBlock()
				if latestBlock != nil {
					log.Info("Latest header", "latestHeader Number", latestBlock)
				}
				err = ocp.stateRoot.AddLatestBlock(latestBlock)
				if err != nil {
					log.Error("Add latest block fail", "err", err)
					return
				}
			}
			err := ocp.processBatch(ocp.headers, ocp.cfg)
			if err == nil {
				ocp.headers = nil
			}
		case <-ocp.quit:
			return
		}
	}
}

func (ocp *OpChainPoller) processBatch(headers []ctypes.Header, chainCfg *cfg.OpEventConfig) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	log.Info("extracting batch", "size", len(headers), "startBlock", firstHeader.Number.String(), "endBlock", lastHeader.Number.String())

	headerMap := make(map[common.Hash]*ctypes.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}

	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: ocp.cfg.Contracts}
	logs, err := ocp.opClient.FilterLogs(filterQuery)
	if err != nil {
		log.Info("failed to extract logs", "err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash")
	}

	if len(logs.Logs) > 0 {
		log.Info("detected logs", "size", len(logs.Logs))
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
	for i := range logs.Logs {
		stateRootEvent, err := ocp.eventProvider.ProcessStateRootEvent(logs.Logs[i])
		if err != nil {
			return err
		}
		log.Info("event list", "stateroot", stateRootEvent.StateRoot)

		ocp.blockInfoChan <- &types.BlockInfo{
			Height:    logs.Logs[i].BlockNumber,
			Hash:      logs.Logs[i].BlockHash.Bytes(),
			Finalized: false,
			StateRoot: *stateRootEvent,
		}
	}
	return nil
}
