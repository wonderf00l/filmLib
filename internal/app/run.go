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
	filmRepository "github.com/wonderf00l/filmLib/internal/repository/film"
	roleRepository "github.com/wonderf00l/filmLib/internal/repository/role"

	actorService "github.com/wonderf00l/filmLib/internal/service/actor"
	authService "github.com/wonderf00l/filmLib/internal/service/auth"
	filmService "github.com/wonderf00l/filmLib/internal/service/film"
	roleService "github.com/wonderf00l/filmLib/internal/service/role"

	actorDelivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1/actor"
	authDelivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1/auth"
	filmDelivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1/film"
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
	filmRepo := filmRepository.New(pool)

	roleService := roleService.New(roleRepo)
	authService := authService.New(authRepo, roleService)
	actorService := actorService.New(actorRepo)
	filmService := filmService.New(filmRepo)

	authHandler := authDelivery.New(authService)
	actorHandler := actorDelivery.New(actorService)
	filmHandler := filmDelivery.New(filmService)

	router := NewRouter()
	router.RegisterRoute(HandlersHTTP{
		auth:  authHandler,
		actor: actorHandler,
		film:  filmHandler,
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
