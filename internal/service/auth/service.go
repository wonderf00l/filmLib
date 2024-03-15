package auth

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	repository "github.com/wonderf00l/filmLib/internal/repository/auth"
)

/*
reg - post{
	login
	pass
}
login - get, post
logout - delete
*checkRole
*/

type Service interface {
	Signup(ctx context.Context, profile entity.Profile) error
}

type authService struct {
	roleManager entity.RoleIdentifier
	repo        repository.Repository
}
