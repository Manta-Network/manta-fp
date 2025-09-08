package config

import (
	"time"
)

var (
	defaultScanSize           = uint32(500)
	defaultL2OutputOracleAddr = "0x0"
	defaultPollInterval       = 5 * time.Second
	defaultStartHeight        = uint64(1)
)

type OpEventConfig struct {
	ChainId            uint          `long:"chain_id" description:"The chain id of the chain"`
	BlockStep          uint64        `long:"block_step" description:"The block step of chain blocks scan"`
	BufferSize         uint32        `long:"buffersize" description:"The maximum number of ethereum blocks that can be stored in the buffer"`
	L2OutputOracleAddr string        `long:"l2outputoracleaddr" description:"The contract address of L2OutputOracle address"`
	PollInterval       time.Duration `long:"pollinterval" description:"The interval between each polling of blocks; the value should be set depending on the block production time but could be set smaller for quick catching up"`
	ScanStartHeight    uint64        `long:"scantartheight" description:"The height from which we start polling the chain"`
}

func DefaultOpEventConfig() OpEventConfig {
	return OpEventConfig{
		BufferSize:         defaultScanSize,
		L2OutputOracleAddr: defaultL2OutputOracleAddr,
		PollInterval:       defaultPollInterval,
		ScanStartHeight:    defaultStartHeight,
	}
}
