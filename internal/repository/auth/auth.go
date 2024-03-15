package auth

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgconn"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

/*func convertErrorPostgres(err error) error {

	switch err {
	case context.DeadlineExceeded:
		return &errPkg.ErrTimeoutExceeded{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		}
	}

	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}*/

// convert err redis

// sql builder?

// вмесете с username | pass также можно ввести секретный код --> он прокидывается в service
// параметр role устанавливается в service, repo просто работает со структурой

/*
type Repository interface {
	AddProfile(profile entity.Profile) error
	CheckProfileExistence(username, password string) bool
	GetUserIdBySession(sessKey string) int
	DeleteSession(sessKey string) error
}
*/

/*
func convertErrorPostgres(err error) error {

	switch err {
	case context.DeadlineExceeded:
		return &errPkg.ErrTimeoutExceeded{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		case strconv.Itoa(23505):
			return &subRepo.ErrSubscriptionAlreadyExist{}
		}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

func (r *subscriptionRepoPG) CreateSubscriptionUser(ctx context.Context, from, to int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err = tx.Exec(ctx, CreateSubscriptionUser, from, to); err != nil {
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
*/

func convertErrorPostgres(err error) error {
	switch err {
	case context.DeadlineExceeded:
		return &errPkg.ErrTimeoutExceeded{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		case strconv.Itoa(23505):
			return &ProfileAlreadyExistsError{}
		}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

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
