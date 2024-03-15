package auth

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"

	repo "github.com/wonderf00l/filmLib/internal/repository"
)

// convert err redis

// sql builder?

// вмесете с username | pass также можно ввести секретный код --> он прокидывается в service
// параметр role устанавливается в service, repo просто работает со структурой

func convertErrorPostgres(err error) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return &ProfileNotFoundError{}
	case errors.Is(err, context.DeadlineExceeded):
		return &errPkg.TimeoutExceededError{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		case strconv.Itoa(repo.PostgresUniqueViolation):
			return &ProfileAlreadyExistsError{}
		}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

// reg
func (r *authRepo) AddProfile(ctx context.Context, profile entity.Profile) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err = r.db.Exec(ctx, InsertProfile, profile.Username, profile.Password, profile.Role); err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return convertErrorPostgres(err)
		}
		return convertErrorPostgres(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return convertErrorPostgres(err)
	}

	return nil
}

// login post
func (r *authRepo) GetProfile(ctx context.Context, username string) (*entity.Profile, error) {
	profile := &entity.Profile{}
	if err := r.db.QueryRow(ctx, SelectProfileByUsername, username).Scan(&profile.ID, &profile.Username, &profile.Password); err != nil {
		return nil, convertErrorPostgres(err)
	}
	return profile, nil
}
