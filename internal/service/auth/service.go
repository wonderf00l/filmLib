package auth

import (
	"context"
	"time"

	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	repository "github.com/wonderf00l/filmLib/internal/repository/auth"
	"github.com/wonderf00l/filmLib/internal/service/role"
)

const (
	cookieStringLen = 16
	sessionLifeTime = 24 * 30 * time.Hour
)

type Service interface {
	Signup(ctx context.Context, profile entity.Profile, roleToken string) error
	CheckCredentials(ctx context.Context, username, password string) (*entity.Session, error)
	GetUserSession(ctx context.Context, key string) (*entity.Session, error)
	Logout(ctx context.Context, sessKey string) error
	// changeCredentials
}

type authService struct {
	repo        repository.Repository
	roleManager role.Service
}
