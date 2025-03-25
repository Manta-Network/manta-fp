package mantastaking

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Manta-Network/manta-fp/ethereum/node"
	cfg "github.com/Manta-Network/manta-fp/symbiotic-fp/config"

	"go.uber.org/zap"
	"math/big"
)

type MantaStakingMiddlewareConfig struct {
	EthClient                     *ethclient.Client
	ChainID                       *big.Int
	MantaStakingMiddlewareAddr    common.Address
	SymbioticOperatorRegisterAddr common.Address
	PrivateKey                    *ecdsa.PrivateKey
	NumConfirmations              uint64
	SafeAbortNonceTooLowCount     uint64
	OperatorName                  string
	RewardAddress                 string
	Commission                    int64
	EnableHsm                     bool
	HsmApiName                    string
	HsmCreden                     string
	HsmAddress                    string
}

func NewMantaStakingMiddlewareConfig(ctx context.Context, config *cfg.Config, logger *zap.Logger, priKeyS string) (*MantaStakingMiddlewareConfig, error) {
	ethClient, err := node.DialEthClientWithTimeout(ctx, config.OpEventConfig.EthRpc, false)
	if err != nil {
		logger.Error("failed to dial eth client", zap.String("err", err.Error()))
		return nil, err
	}
	var privKey *ecdsa.PrivateKey
	if priKeyS != "" {
		privKey, err = crypto.HexToECDSA(priKeyS)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("need to config private key")
	}

	return &MantaStakingMiddlewareConfig{
		EthClient:                     ethClient,
		ChainID:                       big.NewInt(int64(config.OpEventConfig.ChainId)),
		MantaStakingMiddlewareAddr:    common.HexToAddress(config.OpEventConfig.MantaStakingMiddlewareAddress),
		SymbioticOperatorRegisterAddr: common.HexToAddress(config.OpEventConfig.SymbioticOperatorRegisterAddress),
		PrivateKey:                    privKey,
		NumConfirmations:              config.OpEventConfig.NumConfirmations,
		SafeAbortNonceTooLowCount:     config.OpEventConfig.SafeAbortNonceTooLowCount,
		OperatorName:                  config.OperatorName,
		RewardAddress:                 config.RewardAddress,
		Commission:                    int64(config.Commission),
		EnableHsm:                     config.OpEventConfig.EnableHsm,
		HsmApiName:                    config.OpEventConfig.HsmApiName,
		HsmCreden:                     config.OpEventConfig.HsmCreden,
		HsmAddress:                    config.OpEventConfig.HsmAddress,
	}, nil
}
