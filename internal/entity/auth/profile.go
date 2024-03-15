package auth

import "context"

type Role uint8

const (
	Administrator Role = iota + 1
	RegularUser
)

type RoleIdentifier interface {
	GetRole(ctx context.Context, token string) Role
}

type roleIdentifier struct {
	tokenRoleMap map[string]Role
}

func (r *roleIdentifier) GetRole(_ context.Context, token string) Role {
	role, got := r.tokenRoleMap[token]
	if !got {
		return RegularUser
	}
	return role
}

// echo '123' | grep sha256sum
func NewRoleIdentifier() *roleIdentifier {
	return &roleIdentifier{
		map[string]Role{
			"181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b": Administrator,
		},
	}
}

type Profile struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint8  `json:"-"`
}
