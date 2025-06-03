package config

import (
	"fmt"
	"net"
)

const (
	defaultFpApiPort = 8080
	defaultApiHost   = "127.0.0.1"
)

type ApiConfig struct {
	Host string `long:"host" description:"IP of the http server"`
	Port int    `long:"port" description:"Port of the http server"`
}

func (cfg *ApiConfig) Validate() error {
	if cfg.Port < 0 || cfg.Port > 65535 {
		return fmt.Errorf("invalid port: %d", cfg.Port)
	}

	ip := net.ParseIP(cfg.Host)
	if ip == nil {
		return fmt.Errorf("invalid host: %v", cfg.Host)
	}

	return nil
}

func (cfg *ApiConfig) Address() (string, error) {
	if err := cfg.Validate(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), nil
}

func DefaultApiConfig() *ApiConfig {
	return &ApiConfig{
		Port: defaultFpApiPort,
		Host: defaultApiHost,
	}
}
