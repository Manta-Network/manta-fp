package celestia

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/Manta-Network/manta-fp/symbiotic-fp/config"

	"github.com/celestiaorg/go-square/blob"
	"github.com/celestiaorg/go-square/inclusion"
	"github.com/celestiaorg/go-square/namespace"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"github.com/tendermint/tendermint/crypto/merkle"
)

type DAClient struct {
	Client     da.DA
	Namespace  da.Namespace
	GetTimeout time.Duration
}

func NewDAClient(cfg config.CelestiaConfig, authToken string) (*DAClient, error) {
	nsBytes, err := hex.DecodeString(cfg.Namespace)
	if err != nil {
		return nil, err
	}
	if len(nsBytes) != 10 {
		return nil, errors.New("wrong namespace length")
	}
	var client da.DA
	if cfg.DaRpc != "" {
		client, err = proxy.NewClient(cfg.DaRpc, authToken)
		if err != nil {
			return nil, err
		}
	}

	return &DAClient{
		Client:     client,
		Namespace:  append(make([]byte, 19), nsBytes...),
		GetTimeout: cfg.Timeout,
	}, nil
}

func CreateCommitment(data da.Blob, ns da.Namespace) ([]byte, error) {
	ins, err := namespace.From(ns)
	if err != nil {
		return nil, err
	}
	return inclusion.CreateCommitment(blob.New(ins, data, 0), merkle.HashFromByteSlices, 64)
}
