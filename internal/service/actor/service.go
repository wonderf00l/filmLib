package actor

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
	repository "github.com/wonderf00l/filmLib/internal/repository/actor"
)

type Service interface {
	AddActor(ctx context.Context, actor entity.Actor) error
	GetActor(ctx context.Context, actorID int) (*entity.Actor, error)
	UpdateActorData(ctx context.Context, actorID int, updData UpdateActorData) error
	DeleteActorData(ctx context.Context, actorID int) error
}

type actorService struct {
	repo repository.Repository
}

func New(repo repository.Repository) *actorService {
	return &actorService{
		repo: repo,
	}
}
