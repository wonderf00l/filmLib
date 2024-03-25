package auth

const (
	insertProfile           = "INSERT INTO profile (username, password, profile_role) VALUES ($1, $2, $3);"
	selectProfileByUsername = "SELECT id, username, password FROM profile WHERE username = $1;"
	selectProfileByID       = "SELECT id, username, password FROM profile WHERE id = $1;"
	updateProfile           = "UPDATE profile SET username = $1, password = $2, profile_role = $3 WHERE id = $4;"
)
