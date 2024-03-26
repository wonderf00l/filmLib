package actor

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
)

type Repository interface {
	AddActorData(ctx context.Context, actor entity.Actor) error
	GetActorData(ctx context.Context, actorID int) (*entity.Actor, error)
	UpdateActorData(ctx context.Context, actorID int, updFields map[string]any) error
	DeleteActorData(ctx context.Context, actorID int) error
}

type actorRepo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *actorRepo {
	return &actorRepo{
		db: db,
	}
}
