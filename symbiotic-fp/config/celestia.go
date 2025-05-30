package config

import "time"

var (
	defaultNamespace = "00000000000000000000"
	defaultTimeout   = time.Minute
)

type CelestiaConfig struct {
	Namespace string        `long:"namespace" description:"Namespace ID for DA node"`
	DaRpc     string        `long:"da_rpc" description:"Dial address of data availability grpc client"`
	Timeout   time.Duration `long:"time_out" description:"Timeout for celestia requests"`
}

func DefaultCelestiaConfig() CelestiaConfig {
	return CelestiaConfig{
		Namespace: defaultNamespace,
		DaRpc:     "",
		Timeout:   defaultTimeout,
	}
}
