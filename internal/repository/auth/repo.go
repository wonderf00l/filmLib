package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	redis "github.com/redis/go-redis/v9"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
)

type Repository interface {
	AddProfile(ctx context.Context, profile entity.Profile) error
	GetProfile(ctx context.Context, username string) (*entity.Profile, error)
	AddSession(ctx context.Context, session *entity.Session) error
	GetSessionByKey(ctx context.Context, key string) (*entity.Session, error)
	DeleteSessionByKey(ctx context.Context, key string) error
}

// var _ Repository = (*authRepo)(nil)

type authRepo struct {
	db             pgxpool.Pool
	sessionStorage *redis.Client
}
