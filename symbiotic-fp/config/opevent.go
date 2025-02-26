package config

import (
	"time"
)

var (
	defaultScanSize                  = uint32(500)
	defaultEthRpc                    = "http://127.0.0.1:8545"
	defaultEthAddr                   = "0x0"
	defaultPollInterval              = 5 * time.Second
	defaultStartHeight               = uint64(1)
	defaultNumConfirmations          = uint64(10)
	defaultSafeAbortNonceTooLowCount = uint64(3)
)

type OpEventConfig struct {
	ChainId                       uint          `long:"chain_id" description:"The chain id of the chain"`
	StartHeight                   uint64        `long:"start_height" description:"he height from which we start polling the chain"`
	BlockStep                     uint64        `long:"block_step" description:"The block step of chain blocks scan"`
	BufferSize                    uint32        `long:"buffer_size" description:"The maximum number of ethereum blocks that can be stored in the buffer"`
	EthRpc                        string        `long:"eth_rpc" description:"The rpc uri of ethereum"`
	NumConfirmations              uint64        `long:"num_confirmations" description:"Specifies how many blocks are need to consider a transaction confirmed."`
	SafeAbortNonceTooLowCount     uint64        `long:"safe_abort_nonce_too_low_count" description:"Specifies how many ErrNonceTooLow observations are required to give up on a tx at a particular nonce without receiving confirmation."`
	L2OutputOracleAddr            string        `long:"l2outputoracleaddr" description:"The contract address of L2OutputOracle address"`
	PollInterval                  time.Duration `long:"pollinterval" description:"The interval between each polling of blocks; the value should be set depending on the block production time but could be set smaller for quick catching up"`
	EnableHsm                     bool          `long:"enable_hsm" description:"Whether to use cloud hsm"`
	HsmApiName                    string        `long:"hsm_api_name" description:"The api name of hsm"`
	HsmCreden                     string        `long:"hsm_creden" description:"The creden of hsm"`
	HsmAddress                    string        `long:"hsm_address" description:"The address of hsm"`
	MantaStakingMiddlewareAddress string        `long:"manta_staking_middleware_address" description:"the contract address of the manta-staking-middleware"`
}

func DefaultOpEventConfig() OpEventConfig {
	return OpEventConfig{
		StartHeight:                   defaultStartHeight,
		BufferSize:                    defaultScanSize,
		EthRpc:                        defaultEthRpc,
		L2OutputOracleAddr:            defaultEthAddr,
		MantaStakingMiddlewareAddress: defaultEthAddr,
		PollInterval:                  defaultPollInterval,
		NumConfirmations:              defaultNumConfirmations,
		SafeAbortNonceTooLowCount:     defaultSafeAbortNonceTooLowCount,
		EnableHsm:                     false,
		HsmApiName:                    "",
		HsmAddress:                    "",
		HsmCreden:                     "",
	}
}
