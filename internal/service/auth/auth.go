package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"

	entity "github.com/wonderf00l/filmLib/internal/entity/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) assignProfileData(ctx context.Context, profile *entity.Profile, roleToken string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profile.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profile.Password = string(hashedPassword)

	r := strings.NewReader(roleToken)
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
		return err
	}

	profile.Role = uint8(s.roleManager.AssignRole(ctx, fmt.Sprintf("%x", hash.Sum(nil))))
	return nil
}

func (s *authService) Signup(ctx context.Context, profile entity.Profile, roleToken string) error {
	if err := s.assignProfileData(ctx, &profile, roleToken); err != nil {
		return err
	}

	if err := s.repo.AddProfile(ctx, profile); err != nil {
		return err
	}

	return nil
}

func (s *authService) UpdateProfileData(ctx context.Context, profile entity.Profile, roleToken string) error {
	if err := s.assignProfileData(ctx, &profile, roleToken); err != nil {
		return err
	}

	if err := s.repo.UpdateProfile(ctx, profile); err != nil {
		return err
	}

	return nil
}

func (s *authService) CreateSessionForUser(ctx context.Context, username, password string) (*entity.Session, error) {
	profile, err := s.CheckCredentialsByUsername(ctx, username, password)
	if err != nil {
		return nil, err
	}

	sessKey := make([]byte, sessionKeyLen)
	_, err = rand.Read(sessKey)
	if err != nil {
		return nil, err
	}

	session := &entity.Session{
		Key:     base64.StdEncoding.EncodeToString(sessKey),
		UserID:  profile.ID,
		Expires: time.Now().UTC().Add(sessionLifeTime),
	}
	if err = s.repo.AddSession(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *authService) CheckCredentialsByUsername(ctx context.Context, username, password string) (*entity.Profile, error) {
	profile, err := s.repo.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password)); err != nil {
		return nil, &invalidPasswordError{}
	}

	return profile, nil
}

func (s *authService) CheckCredentialsByUserID(ctx context.Context, userID int, password string) (*entity.Profile, error) {
	profile, err := s.repo.GetProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password)); err != nil {
		return nil, &invalidPasswordError{}
	}

	return profile, nil
}

func (s *authService) GetUserSession(ctx context.Context, key string) (*entity.Session, error) {
	sess, err := s.repo.GetSessionByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (s *authService) GetProfileDataByUserID(ctx context.Context, userID int) (*entity.Profile, error) {
	profile, err := s.repo.GetProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *authService) Logout(ctx context.Context, sessKey string) error {
	if err := s.repo.DeleteSessionByKey(ctx, sessKey); err != nil {
		return err
	}
	return nil
}
