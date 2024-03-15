package auth

import (
	"context"
	"crypto/sha256"

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
	profile.Role = uint8(s.roleManager.GetRole(ctx, string(hash.Sum([]byte(roleToken)))))

	if err = s.repo.AddProfile(ctx, profile); err != nil {
		return err
	}

	return nil
}
