package auth

type Profile struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint8  `json:"-"`
}
