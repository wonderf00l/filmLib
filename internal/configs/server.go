package configs

import (
	"fmt"

	"go.uber.org/config"
)

type SrvConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewServerConfig(appCfg *config.YAML, key string) (*SrvConfig, error) {
	cfg := &SrvConfig{}
	if err := appCfg.Get(key).Populate(cfg); err != nil {
		return nil, fmt.Errorf("new srv config: %w", err)
	}
	return cfg, nil
}
