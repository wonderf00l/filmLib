package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/wonderf00l/filmLib/internal/configs"
	"go.uber.org/zap"
)

var (
	shutdownTimeout = 5 * time.Second
	redisTimeout    = 5 * time.Second
	postgresTimeout = 5 * time.Second
)

func Run(ctx context.Context, logger *zap.SugaredLogger, cfgs *configs.Configs) error {
	// create handler
	// create router
	// init srv, create srv

	redisCtx, redisCtxCancel := context.WithTimeout(ctx, redisTimeout)
	defer redisCtxCancel()

	redisCl, err := NewRedisClient(redisCtx, cfgs.RedisCfg)
	if err != nil {
		return &runError{inner: fmt.Errorf("create redis client: %w", err)}
	}
	defer redisCl.Close()

	postgresCtx, postgresCtxCancel := context.WithTimeout(ctx, postgresTimeout)
	defer postgresCtxCancel()

	pool, err := NewPoolPG(postgresCtx, cfgs.PostgresCfg)
	if err != nil {
		return &runError{inner: fmt.Errorf("create postgres pool: %w", err)}
	}
	defer pool.Close()

	wg := sync.WaitGroup{}

	server := NewServer(cfgs.ServerCfg, nil, logger)

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	<-ctx.Done()

	logger.Infoln("Shutting down gracefully")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		return &runError{inner: fmt.Errorf("shutdown http server: %w", err)}
	}

	wg.Wait()

	return nil
}
