package errors

import "fmt"

type Type uint8

type Layer string

const (
	Repo     Layer = "Repository"
	Service  Layer = "Service"
	Delivery Layer = "Delivery"
)

const (
	_ Type = iota
	ErrNotFound
	ErrAlreadyExists
	ErrInvalidInput
	ErrNoAccess
	ErrNoAuth
	ErrTimeout
)

type DeclaredError interface {
	Type() Type
}

// general application errors

type InvalidRoleForActionError struct {
	Need []string
}

func (e *InvalidRoleForActionError) Error() string {
	return fmt.Sprintf("your role doesn't allow you to do this action, should be %v", e.Need)
}

func (e *InvalidRoleForActionError) Type() Type {
	return ErrNoAccess
}

type NotAuthenticatedError struct{}

func (e *NotAuthenticatedError) Error() string {
	return "Auth required"
}

func (e *NotAuthenticatedError) Type() Type {
	return ErrNoAuth
}

type InternalError struct {
	Message string
	Layer   string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("Internal error occured. Message: '%s'. Layer: %s", e.Message, e.Layer)
}

type InvalidTimeFormatError struct {
}

func (e *InvalidTimeFormatError) Error() string {
	return "invalid time format, should be yyyy-mm-dd"
}

func (e *InvalidTimeFormatError) Type() Type {
	return ErrInvalidInput
}

type TimeoutExceededError struct {
}

func (e *TimeoutExceededError) Error() string {
	return "timeout exceeded"
}

func (e *TimeoutExceededError) Type() Type {
	return ErrTimeout
}
