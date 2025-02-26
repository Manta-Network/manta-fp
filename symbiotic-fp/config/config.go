package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/util"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLogLevel                    = zapcore.InfoLevel
	defaultLogDirname                  = "logs"
	defaultLogFilename                 = "sfpd.log"
	defaultConfigFileName              = "sfpd.conf"
	defaultDataDirname                 = "data"
	defaultSubmitRetryInterval         = 1 * time.Second
	defaultSignatureSubmissionInterval = 1 * time.Second
	defaultMaxSubmissionRetries        = 20
	defaultGrpcHost                    = "http://127.0.0.1"
	defaultGrpcPort                    = "34500"
)

var (
	DefaultFpdDir = btcutil.AppDataDir("sfpd", false)
)

type Config struct {
	SubmissionRetryInterval     time.Duration `long:"submissionretryinterval" description:"The interval between each attempt to submit finality signature or public randomness after a failure"`
	SignatureSubmissionInterval time.Duration `long:"signaturesubmissioninterval" description:"The interval between each finality signature(s) submission"`
	MaxSubmissionRetries        uint32        `long:"maxsubmissionretries" description:"The maximum number of retries to submit finality signature or public randomness"`
	GrpcHost                    string        `long:"grpchost" description:"The grpc host of manta relayer"`
	GrpcPort                    string        `long:"grpcport" description:"The grpc port of manta relayer"`

	LogLevel string `long:"loglevel" description:"Logging level for all subsystems" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal"`

	OpEventConfig *OpEventConfig `group:"opeventconfig" namespace:"opeventconfig"`

	DatabaseConfig *DBConfig `group:"dbconfig" namespace:"dbconfig"`

	CelestiaConfig *CelestiaConfig `group:"celestiaconfig" namespace:"celestiaconfig"`

	Metrics *metrics.Config `group:"metrics" namespace:"metrics"`
}

func DefaultConfigWithHome(homePath string) Config {
	opEventConfig := DefaultOpEventConfig()
	celestiaConfig := DefaultCelestiaConfig()
	cfg := Config{
		SignatureSubmissionInterval: defaultSignatureSubmissionInterval,
		SubmissionRetryInterval:     defaultSubmitRetryInterval,
		MaxSubmissionRetries:        defaultMaxSubmissionRetries,
		GrpcPort:                    defaultGrpcPort,
		GrpcHost:                    defaultGrpcHost,
		LogLevel:                    defaultLogLevel.String(),
		DatabaseConfig:              DefaultDBConfigWithHomePath(homePath),
		OpEventConfig:               &opEventConfig,
		CelestiaConfig:              &celestiaConfig,
		Metrics:                     metrics.DefaultFpConfig(),
	}

	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	return cfg
}

func CfgFile(homePath string) string {
	return filepath.Join(homePath, defaultConfigFileName)
}

func LogDir(homePath string) string {
	return filepath.Join(homePath, defaultLogDirname)
}

func LogFile(homePath string) string {
	return filepath.Join(LogDir(homePath), defaultLogFilename)
}

func DataDir(homePath string) string {
	return filepath.Join(homePath, defaultDataDirname)
}

func LoadConfig(homePath string) (*Config, error) {
	cfgFile := CfgFile(homePath)
	if !util.FileExists(cfgFile) {
		return nil, fmt.Errorf("specified config file does "+
			"not exist in %s", cfgFile)
	}

	var cfg Config
	fileParser := flags.NewParser(&cfg, flags.Default)
	err := flags.NewIniParser(fileParser).ParseFile(cfgFile)
	if err != nil {
		return nil, err
	}

	// Make sure everything we just loaded makes sense.
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Validate() error {
	if cfg.Metrics == nil {
		return fmt.Errorf("empty metrics config")
	}

	if err := cfg.Metrics.Validate(); err != nil {
		return fmt.Errorf("invalid metrics config")
	}

	return nil
}
