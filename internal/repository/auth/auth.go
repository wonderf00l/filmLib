package auth

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"

	repo "github.com/wonderf00l/filmLib/internal/repository"
)

func convertErrorPostgres(err error) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return &profileNotFoundError{}
	case errors.Is(err, context.DeadlineExceeded):
		return &errPkg.TimeoutExceededError{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		case strconv.Itoa(repo.PostgresUniqueViolation):
			return &profileAlreadyExistsError{}
		}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

func (r *authRepo) AddProfile(ctx context.Context, profile entity.Profile) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err = r.db.Exec(ctx, insertProfile, profile.Username, profile.Password, profile.Role); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return convertErrorPostgres(err)
		}
		return convertErrorPostgres(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return convertErrorPostgres(err)
	}

	return nil
}

func (r *authRepo) UpdateProfile(ctx context.Context, profile entity.Profile) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err = r.db.Exec(ctx, updateProfile, profile.Username, profile.Password, profile.Role, profile.ID); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return convertErrorPostgres(err)
		}
		return convertErrorPostgres(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return convertErrorPostgres(err)
	}

	return nil
}

func (r *authRepo) GetProfile(ctx context.Context, username string) (*entity.Profile, error) {
	profile := &entity.Profile{}
	if err := r.db.QueryRow(ctx, selectProfileByUsername, username).
		Scan(&profile.ID, &profile.Username, &profile.Password); err != nil {
		return nil, convertErrorPostgres(err)
	}
	return profile, nil
}

func (r *authRepo) GetProfileByID(ctx context.Context, id int) (*entity.Profile, error) {
	profile := &entity.Profile{}
	if err := r.db.QueryRow(ctx, selectProfileByID, id).
		Scan(&profile.ID, &profile.Username, &profile.Password); err != nil {
		return nil, convertErrorPostgres(err)
	}
	return profile, nil
}
