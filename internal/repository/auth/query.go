package auth

const (
	InsertProfile           = "INSERT INTO profile (username, password, profile_role) VALUES ($1, $2, $3);"
	SelectProfileByUsername = "SELECT id, username, password FROM profile WHERE username = $1;"
	SelectProfileByID       = "SELECT id, username, password FROM profile WHERE id = $1;"
	UpdateProfile           = "UPDATE profile SET username = $1, password = $2, profile_role = $3 WHERE id = $4;"
)
