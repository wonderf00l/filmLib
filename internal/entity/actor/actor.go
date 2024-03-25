package actor

import "time"

type Actor struct {
	ID          int
	Name        string
	Gender      string
	DateOfBirth time.Time
}
