package role

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	entity "github.com/wonderf00l/filmLib/internal/entity/role"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

type Repository interface {
	GetUserRole(ctx context.Context, userID int) (entity.Role, error)
}

type roleRepo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *roleRepo {
	return &roleRepo{db: db}
}

func convertErrorPostgres(err error) error {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return &errPkg.TimeoutExceededError{}
	}
	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

func (r *roleRepo) GetUserRole(ctx context.Context, userID int) (entity.Role, error) {
	role := entity.Role(0)
	if err := r.db.QueryRow(ctx, selectUserRole, userID).Scan(&role); err != nil {
		return 0, convertErrorPostgres(err)
	}
	return role, nil
}
