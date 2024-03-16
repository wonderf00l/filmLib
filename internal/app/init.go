package app

import (
	"flag"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/wonderf00l/filmLib/internal/configs"
	"github.com/wonderf00l/filmLib/pkg/logger"
	"go.uber.org/zap"
)

func Init() (*zap.SugaredLogger, *configs.Configs, error) {
	flag.Parse()

	appCfg, err := configs.NewYAML(configs.AppCfgPath)
	if err != nil {
		return nil, nil, &initError{inner: fmt.Errorf("app cfg: %w", err)}
	}

	env, err := configs.EnvFiles(appCfg, configs.EnvKey)
	if err != nil {
		return nil, nil, &initError{inner: err}
	}

	err = godotenv.Load(env...)
	if err != nil {
		return nil, nil, &initError{inner: fmt.Errorf("load .env: %w", err)}
	}

	sConfig, err := configs.NewServerConfig(appCfg, configs.ServerKey)
	if err != nil {
		return nil, nil, &initError{inner: fmt.Errorf("create srv cfg: %w", err)}
	}

	serviceLogger, err := logger.New(logger.NewConfig(
		logger.ConfigureTimeKey(logger.TimeKey),
		logger.ConfigureEncoding(*logger.LogEncoding),
		logger.ConfigureOutput(strings.Split(*logger.LogOutputPaths, ",")),
		logger.ConfigureErrorOutput(strings.Split(*logger.LogErrorOutputPaths, ",")),
	))
	if err != nil {
		return nil, nil, &initError{inner: fmt.Errorf("create logger: %w", err)}
	}

	return serviceLogger, &configs.Configs{
		ServerCfg:   *sConfig,
		RedisCfg:    configs.NewRedisConfig(),
		PostgresCfg: configs.NewPostgresConfig(),
	}, nil
}
