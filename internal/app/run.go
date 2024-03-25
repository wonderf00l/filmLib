package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/wonderf00l/filmLib/internal/configs"
	"go.uber.org/zap"

	actorRepository "github.com/wonderf00l/filmLib/internal/repository/actor"
	authRepository "github.com/wonderf00l/filmLib/internal/repository/auth"
	roleRepository "github.com/wonderf00l/filmLib/internal/repository/role"

	actorService "github.com/wonderf00l/filmLib/internal/service/actor"
	authService "github.com/wonderf00l/filmLib/internal/service/auth"
	roleService "github.com/wonderf00l/filmLib/internal/service/role"

	actorDelivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1/actor"
	authDelivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1/auth"
)

var (
	shutdownTimeout = 5 * time.Second
	redisTimeout    = 5 * time.Second
	postgresTimeout = 5 * time.Second
)

func Run(ctx context.Context, logger *zap.SugaredLogger, cfgs *configs.Configs) error {
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

	authRepo := authRepository.New(pool, redisCl)
	roleRepo := roleRepository.New(pool)
	actorRepo := actorRepository.New(pool)

	roleService := roleService.New(roleRepo)
	authService := authService.New(authRepo, roleService)
	actorService := actorService.New(actorRepo)

	authHandler := authDelivery.New(authService)
	actorHandler := actorDelivery.New(actorService)

	router := NewRouter()
	router.RegisterRoute(HandlersHTTP{
		auth:  authHandler,
		actor: actorHandler,
	}, logger, authService, roleService)

	wg := sync.WaitGroup{}

	server := NewServer(cfgs.ServerCfg, router.Mux, logger)

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
