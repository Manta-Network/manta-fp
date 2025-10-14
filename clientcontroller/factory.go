package clientcontroller

import (
	"fmt"
	"github.com/Manta-Network/manta-fp/clientcontroller/api"
	"github.com/Manta-Network/manta-fp/clientcontroller/babylon"
	fpcfg "github.com/Manta-Network/manta-fp/finality-provider/config"
	bbnclient "github.com/babylonlabs-io/babylon/v3/client/client"
	"go.uber.org/zap"
)

func NewBabylonController(bbnConfig *fpcfg.BBNConfig, logger *zap.Logger) (api.BabylonController, error) {
	bbnCfg := bbnConfig.ToBabylonConfig()
	bbnClient, err := bbnclient.New(
		&bbnCfg,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Babylon rpc client: %w", err)
	}
	cc, err := babylon.NewBabylonController(bbnClient, bbnConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create Babylon rpc client: %w", err)
	}

	return cc, nil
}
