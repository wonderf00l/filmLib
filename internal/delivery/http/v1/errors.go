package v1

import (
	"errors"
	"fmt"
	"net/http"

	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

// general http delivery errors

const (
	ApplicationJSON = "application/json"
)

type NoAuthCookieError struct{}

func (e *NoAuthCookieError) Error() string {
	return "no auth cookie has been provided"
}

func (e *NoAuthCookieError) Type() errPkg.Type {
	return errPkg.ErrNoAuth
}

type MiddlewareNotSetError struct {
	MWTypes []MiddlewareType
}

func (e *MiddlewareNotSetError) Error() string {
	return fmt.Sprintf("necessary middlewares %v aren't set", e.MWTypes)
}

type InvalidContentTypeError struct {
	PreferredType string
}

func (e *InvalidContentTypeError) Error() string {
	return fmt.Sprintf("invalid content type, should be %s", e.PreferredType)
}

func (e *InvalidContentTypeError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type InvalidBodyError struct{}

func (e *InvalidBodyError) Error() string {
	return "invalid body"
}

func (e *InvalidBodyError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type InvalidQueryParamError struct {
	Params map[string]string
	Need   string
}

func (e *InvalidQueryParamError) Error() string {
	return fmt.Sprintf("invalid query params: %v, provide %s", e.Params, e.Need)
}

func (e *InvalidQueryParamError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type InvalidUrlParamsError struct {
	Params map[string]string
}

func (e *InvalidUrlParamsError) Error() string {
	return fmt.Sprintf("invalid URL params: %v", e.Params)
}

func (e *InvalidUrlParamsError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type MissingBodyParamsError struct {
	Params []string
}

func (e *MissingBodyParamsError) Error() string {
	return fmt.Sprintf("missing body params: %v", e.Params)
}

func (e *MissingBodyParamsError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

func getCodeStatusHTTP(err error) (errCode string, httpStatus int) {
	var declaredErr errPkg.DeclaredError
	if errors.As(err, &declaredErr) {
		switch declaredErr.Type() {
		case errPkg.ErrInvalidInput:
			return "bad_input", http.StatusBadRequest
		case errPkg.ErrNotFound:
			return "not_found", http.StatusNotFound
		case errPkg.ErrAlreadyExists:
			return "already_exists", http.StatusConflict
		case errPkg.ErrNoAuth:
			return "no_auth", http.StatusUnauthorized
		case errPkg.ErrNoAccess:
			return "no_access", http.StatusForbidden
		case errPkg.ErrTimeout:
			return "timeout", http.StatusRequestTimeout
		}
	}

	return "internal_error", http.StatusInternalServerError
}
