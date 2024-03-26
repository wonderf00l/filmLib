package actor

import (
	"time"

	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
)

type actorData struct {
	ID          *int    `json:"id,omitempty" example:"123"`
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
		date, err := time.Parse(delivery.LayoutISO, *d.DateOfBirth)
		if err != nil {
			return nil, &errPkg.InvalidTimeFormatError{}
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
		DateOfBirth: a.DateOfBirth.Format(delivery.LayoutISO),
	}
}

func updateActorDeliveryToService(d actorData) (map[string]any, error) {
	updFields := make(map[string]any, 3)

	if d.Name != nil {
		updFields[entity.NameAttr] = *d.Name
	}
	if d.Gender != nil {
		updFields[entity.GenderAttr] = *d.Gender
	}
	if d.DateOfBirth != nil {
		date, err := time.Parse(delivery.LayoutISO, *d.DateOfBirth)
		if err != nil {
			return nil, &errPkg.InvalidTimeFormatError{}
		}
		updFields[entity.DateAttr] = date
	}

	return updFields, nil
}
