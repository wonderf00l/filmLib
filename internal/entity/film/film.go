package film

import "time"

type Film struct {
	ID          int
	Name        string
	Description string
	ReleaseDate time.Time
	Rate        uint8
	Actors      map[string]struct{}
}

// []string -> set{} -- fromDelToServ
// set{} --> []string -- fromServToRepo
// convert set to []string for repository
// actors as []Actor in another struct
