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
	defaultCommission                  = 1000
)

var (
	DefaultFpdDir = btcutil.AppDataDir("sfpd", false)
)

type Config struct {
	SubmissionRetryInterval     time.Duration `long:"submissionretryinterval" description:"The interval between each attempt to submit finality signature or public randomness after a failure"`
	SignatureSubmissionInterval time.Duration `long:"signaturesubmissioninterval" description:"The interval between each finality signature(s) submission"`
	MaxSubmissionRetries        uint32        `long:"maxsubmissionretries" description:"The maximum number of retries to submit finality signature or public randomness"`
	OperatorName                string        `long:"operatorname" description:"The name of operator; The name needs to be registered in the contract"`
	RewardAddress               string        `long:"rewardaddress" description:"The manta address to receive fp rewards"`
	Commission                  uint64        `long:"commission" description:"The custom commission, 10000 = 100%"`
	LogLevel                    string        `long:"loglevel" description:"Logging level for all subsystems" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal"`
	SymbioticStakeUrl           string        `long:"symbioticstakeurl" description:"The url to get the symbiotic stake amount"`
	StakeLimit                  string        `long:"stakelimit" description:"The limit of the total stake required to start symbiotic fp"`
	EnableKms                   bool          `long:"enablekms" description:"Whether to use aws kms"`

	OpEventConfig *OpEventConfig `group:"opeventconfig" namespace:"opeventconfig"`

	DatabaseConfig *DBConfig `group:"dbconfig" namespace:"dbconfig"`

	CelestiaConfig *CelestiaConfig `group:"celestiaconfig" namespace:"celestiaconfig"`

	Metrics *metrics.Config `group:"metrics" namespace:"metrics"`

	Api *ApiConfig `group:"api" namespace:"api"`
}

func DefaultConfigWithHome(homePath string) Config {
	opEventConfig := DefaultOpEventConfig()
	celestiaConfig := DefaultCelestiaConfig()
	cfg := Config{
		SignatureSubmissionInterval: defaultSignatureSubmissionInterval,
		SubmissionRetryInterval:     defaultSubmitRetryInterval,
		MaxSubmissionRetries:        defaultMaxSubmissionRetries,
		OperatorName:                "",
		SymbioticStakeUrl:           "",
		StakeLimit:                  "",
		RewardAddress:               defaultEthAddr,
		Commission:                  defaultCommission,
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
