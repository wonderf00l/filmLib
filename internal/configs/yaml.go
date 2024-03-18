package configs

import (
	"fmt"

	"go.uber.org/config"
)

var (
	AppCfgPath = "configs/app.yml"
	ServerKey  = "app.server"
	EnvKey     = "app.env"
)

func NewYAML(filename string) (*config.YAML, error) {
	cfg, err := config.NewYAML(config.File(filename))
	if err != nil {
		return nil, fmt.Errorf("new yaml cfg: %w", err)
	}
	return cfg, nil
}
