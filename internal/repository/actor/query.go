package actor

const (
	insertActorData = "INSERT INTO actor (name, gender, date_of_birth) VALUES ($1,$2,$3);"
	selectActorData = "SELECT id, name, gender, date_of_birth FROM actor WHERE id = $1;"
	deleteActorData = "DELETE FROM actor WHERE id = $1;"
)
