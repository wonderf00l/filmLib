package film

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	entity "github.com/wonderf00l/filmLib/internal/entity/film"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"
	repo "github.com/wonderf00l/filmLib/internal/repository"
)

func convertErrorPostgres(err error) error {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return &errPkg.TimeoutExceededError{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		case strconv.Itoa(repo.PostgresUniqueViolation):
			return &filmAlreadyExistsError{}
		}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

func (f *filmRepo) AddFilm(ctx context.Context, film *entity.Film, actors []string) error {
	tx, err := f.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err := tx.Exec(ctx, insertFilmInfo, film.Name, film.Description, film.ReleaseDate, film.Rate, actors); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return convertErrorPostgres(err)
		}
		return convertErrorPostgres(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return convertErrorPostgres(err)
	}

	return nil
}
