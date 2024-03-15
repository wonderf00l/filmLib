package auth

const (
	InsertProfile           = "INSERT INTO profile (username, password, role) VALUES ($1, $2, $3);"
	SelectProfileByUsername = "SELECT id, username, password FROM profile WHERE username = $1;"
)
