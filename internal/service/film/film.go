package film

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/film"
)

func (f *filmService) AddFilm(ctx context.Context, film *entity.Film) error {
	if err := validateFilmData(film); err != nil {
		return err
	}

	_, actors := filmFromServiceToRepository(film)
	if err := f.repo.AddFilm(ctx, film, actors); err != nil {
		return err
	}

	return nil
}
