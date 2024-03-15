package auth

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis"
	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

func convertErrorRedis(err error) error {
	switch {
	case errors.Is(err, redis.Nil):
		return &UserNotAuthenticatedError{}
	case errors.Is(err, context.DeadlineExceeded):
		return &errPkg.TimeoutExceededError{}
	}

	return &errPkg.InternalError{Message: err.Error(), Layer: string(errPkg.Repo)}
}

// login post
func (s *authRepo) AddSession(ctx context.Context, session *entity.Session) error {
	res := s.sessionStorage.SetNX(ctx, session.Key, session.UserID, time.Duration(session.Expires.Sub(time.Now().UTC())))

	if err := res.Err(); err != nil {
		return convertErrorRedis(err)
	}

	return nil
}

// auth mw
func (s *authRepo) GetSessionByKey(ctx context.Context, key string) (*entity.Session, error) {
	res := s.sessionStorage.Get(ctx, key)

	if err := res.Err(); err != nil {
		return nil, convertErrorRedis(err)
	}

	var err error
	sess := &entity.Session{Key: key}
	sess.UserID, err = res.Int()
	if err != nil {
		return nil, convertErrorRedis(err)
	}

	return sess, nil
}

// logout delete
func (s *authRepo) DeleteSessionByKey(ctx context.Context, key string) error {
	res := s.sessionStorage.Del(ctx, key)
	if err := res.Err(); err != nil {
		return convertErrorRedis(err)
	}
	return nil
}
