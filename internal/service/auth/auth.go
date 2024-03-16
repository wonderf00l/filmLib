package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"time"

	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Signup(ctx context.Context, profile entity.Profile, roleToken string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profile.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profile.Password = string(hashedPassword)

	hash := sha256.New()
	profile.Role = uint8(s.roleManager.AssignRole(ctx, string(hash.Sum([]byte(roleToken)))))

	if err = s.repo.AddProfile(ctx, profile); err != nil {
		return err
	}

	return nil
}

func (s *authService) CheckCredentials(ctx context.Context, username, password string) (*entity.Session, error) {
	profile, err := s.repo.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password)); err != nil {
		return nil, &InvalidPasswordError{}
	}

	cookieString := make([]byte, cookieStringLen)
	_, err = rand.Read(cookieString)
	if err != nil {
		return nil, err
	}

	session := &entity.Session{
		Key:     string(cookieString),
		UserID:  profile.ID,
		Expires: time.Now().UTC().Add(sessionLifeTime),
	}
	if err = s.repo.AddSession(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *authService) GetUserSession(ctx context.Context, key string) (*entity.Session, error) {
	sess, err := s.repo.GetSessionByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (s *authService) Logout(ctx context.Context, sessKey string) error {
	if err := s.repo.DeleteSessionByKey(ctx, sessKey); err != nil {
		return err
	}
	return nil
}
