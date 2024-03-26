package film

import (
	"time"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
	entity "github.com/wonderf00l/filmLib/internal/entity/film"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

type filmData struct {
	Name        string   `json:"name" example:"Titanic"`
	Description string   `json:"description" example:"cool film"`
	ReleaseDate string   `json:"release_date" example:"2002-10-11"`
	Rate        uint8    `json:"rate" example:"8"`
	Actors      []string `json:"actors" example:"['Leonardo DiCaprio', 'Kate Winslet','Billy Zane']"`
}

func filmDataDeliveryToService(d filmData) (*entity.Film, error) {
	film := &entity.Film{}

	date, err := time.Parse(delivery.LayoutISO, d.ReleaseDate)
	if err != nil {
		return nil, &errPkg.InvalidTimeFormatError{}
	}

	film.ReleaseDate = date
	film.Name = d.Name
	film.Description = d.Description
	film.Rate = d.Rate

	actors := make(map[string]struct{}, len(d.Actors))
	for _, actor := range d.Actors {
		actors[actor] = struct{}{}
	}
	film.Actors = actors

	return film, nil
}
