package role

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/role"
	repository "github.com/wonderf00l/filmLib/internal/repository/role"
)

type Service interface {
	AssignRole(ctx context.Context, token string) entity.Role
	GetUserRole(ctx context.Context, userID int) entity.Role
}

type roleService struct {
	tokenRoleMap map[string]entity.Role
	repo         repository.Repository
}

// echo '123' | grep sha256sum
func New(repo repository.Repository) *roleService {
	return &roleService{
		tokenRoleMap: map[string]entity.Role{
			"181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b": entity.Administrator,
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

func (r *roleService) CheckUserRole(ctx context.Context, userID int) (entity.Role, error) {
	role, err := r.repo.GetUserRole(ctx, userID)
	if err != nil {
		return 0, err
	}
	return role, nil
}
