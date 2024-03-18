package role

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/role"
	repository "github.com/wonderf00l/filmLib/internal/repository/role"
)

type Service interface {
	AssignRole(ctx context.Context, token string) entity.Role
	GetUserRole(ctx context.Context, userID int) (entity.Role, error)
}

type roleService struct {
	tokenRoleMap map[string]entity.Role
	repo         repository.Repository
}

// echo 'admToken' | sha256sum
func New(repo repository.Repository) *roleService {
	return &roleService{
		tokenRoleMap: map[string]entity.Role{
			"fbb89c5d1d266c5b573d69c50a56ae95ff472d43683e9ca75ddfbb5b1c098af5": entity.Administrator,
		},
		repo: repo,
	}
}

func (r *roleService) AssignRole(_ context.Context, token string) entity.Role {
	role, got := r.tokenRoleMap[token]
	if !got {
		return entity.RegularUser
	}
	return role
}

func (r *roleService) GetUserRole(ctx context.Context, userID int) (entity.Role, error) {
	role, err := r.repo.GetUserRole(ctx, userID)
	if err != nil {
		return 0, err
	}
	return role, nil
}
