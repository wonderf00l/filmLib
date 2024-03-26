package film

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/film"
	repository "github.com/wonderf00l/filmLib/internal/repository/film"
)

type Service interface {
	AddFilm(ctx context.Context, film *entity.Film) error
}

type filmService struct {
	repo repository.Repository
}

func New(repo repository.Repository) *filmService {
	return &filmService{repo: repo}
}
