package auth

import (
	"context"
	"time"

	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	repository "github.com/wonderf00l/filmLib/internal/repository/auth"
	"github.com/wonderf00l/filmLib/internal/service/role"
)

const (
	sessionKeyLen   = 16
	sessionLifeTime = 24 * 30 * time.Hour
)

type Service interface {
	Signup(ctx context.Context, profile entity.Profile, roleToken string) error
	CheckCredentialsByUserID(ctx context.Context, userID int, password string) (*entity.Profile, error)
	CheckCredentialsByUsername(ctx context.Context, username, password string) (*entity.Profile, error)
	UpdateProfileData(ctx context.Context, profile entity.Profile, roleToken string) error
	CreateSessionForUser(ctx context.Context, username, password string) (*entity.Session, error)
	GetUserSession(ctx context.Context, key string) (*entity.Session, error)
	Logout(ctx context.Context, sessKey string) error
}

type authService struct {
	repo        repository.Repository
	roleManager role.Service
}

func New(repo repository.Repository, roleMan role.Service) *authService {
	return &authService{
		repo:        repo,
		roleManager: roleMan,
	}
}
