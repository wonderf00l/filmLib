package film

import entity "github.com/wonderf00l/filmLib/internal/entity/film"

func validateFilmData(film *entity.Film) error {
	if err := isValidName(film.Name); err != nil {
		return err
	}
	if err := isValidDescription(film.Description); err != nil {
		return err
	}
	if err := isValidRate(film.Rate); err != nil {
		return err
	}
	return nil
}

func isValidName(filmName string) error {
	if len(filmName) < 1 || len(filmName) > 150 {
		return &invalidFilmNameError{}
	}
	return nil
}

func isValidDescription(desc string) error {
	if len(desc) > 1000 {
		return &invalidFilmDescError{}
	}
	return nil
}

func isValidRate(rate uint8) error {
	if rate < 0 || rate > 10 {
		return &invalidFilmRateError{}
	}
	return nil
}
