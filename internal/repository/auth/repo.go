package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	redis "github.com/redis/go-redis/v9"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
)

/* register -- addProfile(Profile) error
	if err - user already exists - already registered

login post -- getProfile(username, passwd - hashed) bool
	if false - no such user - not registered

checkAuth - getUserIdBySessin(cookie) (user_id, error)
	if err - session not found --> user not authenticated
	internal --> log

logout delete - delSession(cookie) error
	if error:
	no such session - not auth-ed
	internal - log

ROLES:
	getRole(user_id) user_role
*/

type Repository interface {
	AddProfile(ctx context.Context, profile entity.Profile) error
	CheckProfileExistence(ctx context.Context, username, password string) bool
	GetUserIdBySession(ctx context.Context, sessKey string) int
	DeleteSession(ctx context.Context, sessKey string) error
}

// var _ Repository = (*authRepo)(nil)

type authRepo struct {
	db             pgxpool.Pool
	sessionStorage *redis.Client
}
