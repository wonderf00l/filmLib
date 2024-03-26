package actor

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	entity "github.com/wonderf00l/filmLib/internal/entity/actor"

	errPkg "github.com/wonderf00l/filmLib/internal/errors"
	repo "github.com/wonderf00l/filmLib/internal/repository"

	sq "github.com/Masterminds/squirrel"
)

const (
	relationNamePostgres = "actor"
)

func convertErrorPostgres(err error) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return &actorNotFoundError{}
	case errors.Is(err, context.DeadlineExceeded):
		return &errPkg.TimeoutExceededError{}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.SQLState() {
		case strconv.Itoa(repo.PostgresUniqueViolation):
			return &actorAlreadyExistsError{}
		}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

func (a *actorRepo) AddActorData(ctx context.Context, actor entity.Actor) error {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err = a.db.Exec(ctx, insertActorData, actor.Name, actor.Gender, actor.DateOfBirth); err != nil {
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

func (a *actorRepo) GetActorData(ctx context.Context, actorID int) (*entity.Actor, error) {
	var actor entity.Actor
	if err := a.db.QueryRow(ctx, selectActorData, actorID).Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.DateOfBirth); err != nil {
		return nil, convertErrorPostgres(err)
	}
	return &actor, nil
}

func (a *actorRepo) UpdateActorData(ctx context.Context, actorID int, updFields map[string]any) error {
	sqlRow, args, err := sq.Update(relationNamePostgres).
		SetMap(updFields).
		Where(sq.Eq{"id": actorID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return convertErrorPostgres(err)
	}

	tx, err := a.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err := a.db.Exec(ctx, sqlRow, args...); err != nil {
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

func (a *actorRepo) DeleteActorData(ctx context.Context, actorID int) error {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return convertErrorPostgres(err)
	}

	if _, err := tx.Exec(ctx, deleteActorData, actorID); err != nil {
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
