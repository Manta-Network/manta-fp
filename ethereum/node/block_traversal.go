package node

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Manta-Network/manta-fp/ethereum/bigint"

	"go.uber.org/zap"
)

var (
	ErrBlockTraversalAheadOfProvider            = errors.New("the BlockTraversal's internal state is ahead of the provider")
	ErrBlockTraversalAndProviderMismatchedState = errors.New("the BlockTraversal and provider have diverged in state")
	ErrBlockTraversalCheckBlockFail             = errors.New("the BlockTraversal check block height fail")
)

type BlockTraversal struct {
	ethClient EthClient
	chainId   uint
	log       *zap.Logger

	latestBlock        *big.Int
	lastTraversedBlock *big.Int

	blockConfirmationDepth *big.Int
}

func NewBlockTraversal(ethClient EthClient, fromBlock *big.Int, confDepth *big.Int, chainId uint, logger *zap.Logger) *BlockTraversal {
	return &BlockTraversal{
		ethClient:              ethClient,
		lastTraversedBlock:     fromBlock,
		blockConfirmationDepth: confDepth,
		chainId:                chainId,
		log:                    logger,
	}
}

func (f *BlockTraversal) LatestBlock() *big.Int {
	return f.latestBlock
}

func (f *BlockTraversal) LastTraversedHeader() *big.Int {
	return f.lastTraversedBlock
}

func (f *BlockTraversal) NextHeaders(maxSize uint64) ([]types.Header, error) {
	latestHeader, err := f.ethClient.BlockHeaderByNumber(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		f.latestBlock = latestHeader.Number
	}

	f.log.Info("header traversal db latest header: ", zap.Uint64("latestBlock", latestHeader.Number.Uint64()))

	endHeight := new(big.Int).Sub(latestHeader.Number, f.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		return nil, nil
	}

	f.log.Info("header traversal last traversed deader to json: ", zap.Uint64("lastTraversedBlock", f.lastTraversedBlock.Uint64()))

	if f.lastTraversedBlock != nil {
		cmp := f.lastTraversedBlock.Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrBlockTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero
	if f.lastTraversedBlock != nil {
		nextHeight = new(big.Int).Add(f.lastTraversedBlock, bigint.One)
	}

	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.ethClient.BlockHeadersByRange(nextHeight, endHeight, f.chainId)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}
	if len(headers) == 0 {
		return nil, nil
	}
	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if f.lastTraversedBlock != nil && headers[0].Number.Uint64() != new(big.Int).Add(f.lastTraversedBlock, big.NewInt(1)).Uint64() {
		f.log.Error("Err header traversal and provider mismatched state", zap.Uint64("parentNumber = ", headers[0].Number.Uint64()), zap.Uint64("number = ", f.lastTraversedBlock.Uint64()))
		return nil, ErrBlockTraversalAndProviderMismatchedState
	}
	f.lastTraversedBlock = headers[numHeaders-1].Number
	return headers, nil
}

func (f *BlockTraversal) ChangeLastTraversedHeaderByDelAfter(dbLatestBlock *big.Int) {
	f.lastTraversedBlock = dbLatestBlock
}
