package actor

import (
	"time"

	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
)

const (
	layoutISO = "2006-01-02"
)

type actorData struct {
	Name        *string `json:"name" example:"Ryan Gosling"`
	Gender      *string `json:"gender" example:"male"`
	DateOfBirth *string `json:"date_of_birth" example:"2006-01-02"`
}

func actorDataDeliveryToService(d actorData) (*entity.Actor, error) {
	actor := entity.Actor{}
	if d.Name != nil {
		actor.Name = *d.Name
	}
	if d.Gender != nil {
		actor.Gender = *d.Gender
	}
	if d.DateOfBirth != nil {
		date, err := time.Parse(layoutISO, *d.DateOfBirth)
		if err != nil {
			return nil, &invalidTimeFormatError{}
		}
		actor.DateOfBirth = date
	}
	return &actor, nil
}

type getActorData struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

func getActorServiceToDelivery(a entity.Actor) getActorData {
	return getActorData{
		ID:          a.ID,
		Name:        a.Name,
		Gender:      a.Gender,
		DateOfBirth: a.DateOfBirth.Format(layoutISO),
	}
}

// func actorDataDeliveryToService(d actorData) map[string]any {
// 	updFields := make(map[string]any, 3)

// 	if d.Name != nil {
// 		updFields["name"] = *d.Name
// 	}
// 	if d.Gender != nil {
// 		updFields["gender"] = *d.Gender
// 	}
// 	if d.DateOfBirth != nil {
// 		updFields["date_of_birth"] = *d.DateOfBirth
// 	}

// 	return updFields
// }
