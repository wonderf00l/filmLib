package configs

import (
	"fmt"

	"go.uber.org/config"
)

func EnvFiles(appCfg *config.YAML, key string) ([]string, error) {
	files := make([]string, 0)
	if err := appCfg.Get(key).Populate(&files); err != nil {
		return nil, fmt.Errorf("get .env files: %w", err)
	}
	return files, nil
}
