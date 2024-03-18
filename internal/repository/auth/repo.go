package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	redis "github.com/redis/go-redis/v9"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
)

type Repository interface {
	AddProfile(ctx context.Context, profile entity.Profile) error
	UpdateProfile(ctx context.Context, profile entity.Profile) error
	GetProfile(ctx context.Context, username string) (*entity.Profile, error)
	GetProfileByID(ctx context.Context, id int) (*entity.Profile, error)
	AddSession(ctx context.Context, session *entity.Session) error
	GetSessionByKey(ctx context.Context, key string) (*entity.Session, error)
	DeleteSessionByKey(ctx context.Context, key string) error
}

type authRepo struct {
	db             *pgxpool.Pool
	sessionStorage *redis.Client
}

func New(db *pgxpool.Pool, sessStorage *redis.Client) *authRepo {
	return &authRepo{
		db:             db,
		sessionStorage: sessStorage,
	}
}
