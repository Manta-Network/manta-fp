package mantastaking

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Manta-Network/manta-fp/ethereum/node"
	common2 "github.com/Manta-Network/manta-fp/symbiotic-fp/common"
	cfg "github.com/Manta-Network/manta-fp/symbiotic-fp/config"
	kmssigner "github.com/Manta-Network/manta-fp/symbiotic-fp/kms"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"go.uber.org/zap"
	"math/big"
)

type MantaStakingMiddlewareConfig struct {
	EthClient                     *ethclient.Client
	ChainID                       *big.Int
	MantaStakingMiddlewareAddr    common.Address
	SymbioticOperatorRegisterAddr common.Address
	SymbioticStakeUrl             string
	StakeLimit                    string
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
	EnableKms                     bool
	KmsID                         string
	KmsRegion                     string
	KmsClient                     *kms.Client
}

func NewMantaStakingMiddlewareConfig(ctx context.Context, config *cfg.Config, logger *zap.Logger, priKeyS string, kmsID string, kmsRegion string) (*MantaStakingMiddlewareConfig, error) {
	ethClient, err := node.DialEthClientWithTimeout(ctx, config.OpEventConfig.EthRpc, false)
	if err != nil {
		logger.Error("failed to dial eth client", zap.String("err", err.Error()))
		return nil, err
	}
	var kmsClient *kms.Client
	var privKey *ecdsa.PrivateKey
	if config.EnableKms {
		kmsClient, err = kmssigner.NewKmsClientFromConfig(context.Background(), kmsRegion)
		if err != nil {
			return nil, fmt.Errorf("failed to create the kms client: %w", err)
		}
	} else {
		if priKeyS != "" {
			privKey, err = crypto.HexToECDSA(priKeyS)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("need to config private key")
		}
	}

	return &MantaStakingMiddlewareConfig{
		EthClient:                     ethClient,
		ChainID:                       big.NewInt(int64(config.OpEventConfig.ChainId)),
		MantaStakingMiddlewareAddr:    common.HexToAddress(config.OpEventConfig.MantaStakingMiddlewareAddress),
		SymbioticOperatorRegisterAddr: common.HexToAddress(config.OpEventConfig.SymbioticOperatorRegisterAddress),
		SymbioticStakeUrl:             config.SymbioticStakeUrl,
		StakeLimit:                    common2.StakeLimit,
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
		EnableKms:                     config.EnableKms,
		KmsID:                         kmsID,
		KmsRegion:                     kmsRegion,
		KmsClient:                     kmsClient,
	}, nil
}
