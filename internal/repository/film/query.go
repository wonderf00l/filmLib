package film

const (
	insertFilmInfo = "INSERT INTO film (name, description, release_date, rating, actors) VALUES ($1, $2, $3, $4, $5);"
)
