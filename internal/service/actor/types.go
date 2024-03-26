package actor

import (
	"errors"

	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
)

type UpdateActorData map[string]any

func (u UpdateActorData) Validate() error {
	name, got := u[entity.NameAttr]

	nameStr, ok := name.(string)
	if got && !ok {
		return errors.New("invalid type of actor name")
	}
	if got && ok && nameStr == "" {
		return &invalidNameError{}
	}

	gender, got := u[entity.GenderAttr]

	genderStr, ok := gender.(string)
	if got && !ok {
		return errors.New("invalid type of actor gender")
	}
	if got && ok && genderStr == "" {
		return &invalidGenderError{}
	}

	if err := isValidName(nameStr); err != nil {
		return err
	}
	if err := isValidGender(genderStr); err != nil {
		return err
	}

	return nil
}

func updActorDataServiceToRepository(data UpdateActorData) map[string]any {
	m := make(map[string]any, len(data))
	for field, val := range data {
		m[field] = val
	}
	return m
}
