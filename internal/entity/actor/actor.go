package actor

import "time"

const (
	NameAttr   = "name"
	GenderAttr = "gender"
	DateAttr   = "date_of_birth"
)

type Actor struct {
	ID          int
	Name        string
	Gender      string
	DateOfBirth time.Time
}


