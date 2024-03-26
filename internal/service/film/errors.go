package film

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type invalidFilmNameError struct {
}

func (e *invalidFilmNameError) Error() string {
	return "invalid film name was provided: should be - len[1:150]"
}

func (e *invalidFilmNameError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type invalidFilmDescError struct {
}

func (e *invalidFilmDescError) Error() string {
	return "invalid film description was provided: should be - len[0:1000]"
}

func (e *invalidFilmDescError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type invalidFilmRateError struct {
}

func (e *invalidFilmRateError) Error() string {
	return "invalid film rating was provided: should be - from 0 to 10"
}

func (e *invalidFilmRateError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
