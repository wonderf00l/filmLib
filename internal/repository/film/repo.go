package film

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	entity "github.com/wonderf00l/filmLib/internal/entity/film"
)

type Repository interface {
	AddFilm(ctx context.Context, film *entity.Film, actors []string) error
}

type filmRepo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *filmRepo {
	return &filmRepo{db: db}
}
