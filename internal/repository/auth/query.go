package auth

const (
	InsertProfile = "INSERT INTO profile (username, password, role) VALUES ($1, $2, $3);"
)
