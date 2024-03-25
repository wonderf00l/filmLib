package actor

import (
	"unicode"

	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
)

func validateActorData(actor entity.Actor) error {
	if err := isValidName(actor.Name); err != nil {
		return err
	}

	if err := isValidGender(actor.Gender); err != nil {
		return err
	}

	return nil
}

func isValidName(name string) error {
	if len(name) < 2 || len(name) > 50 {
		return &invalidNameError{}
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && r != '\'' {
			return &invalidNameError{}
		}
	}
	return nil
}

func isValidGender(gender string) error {
	if gender != "male" && gender != "female" {
		return &invalidGenderError{}
	}
	return nil
}
