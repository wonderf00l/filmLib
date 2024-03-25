package auth

type Profile struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     uint8  `json:"-"`
}
