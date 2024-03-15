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
	SessionLifeTime = 24 * 30 * time.Hour
)

type Service interface {
	Signup(ctx context.Context, profile entity.Profile) error
	CheckCredentials(ctx context.Context, username, password string) error
	GetUserSession(ctx context.Context, key string) (*entity.Session, error)
	Logout(ctx context.Context, sessKey string) error
}

type authService struct {
	repo        repository.Repository
	roleManager role.Service
}
