package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wonderf00l/filmLib/internal/configs"
)

const (
	maxConnDB = 500
	schemaDB  = "filmLib"
)

func NewPoolPG(ctx context.Context, config configs.PostgresConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsnPG(config))

	cfg.MaxConns = maxConnDB
	cfg.ConnConfig.RuntimeParams["search_path"] = schemaDB

	if err != nil {
		return nil, fmt.Errorf("parse pool config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("new postgres pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping new pool: %w", err)
	}
	return pool, err
}

func dsnPG(cfg configs.PostgresConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password,
		cfg.Host, cfg.Port,
		cfg.Database)
}
