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
type ErrNotAuthenticated struct{}

func (e *ErrNotAuthenticated) Error() string {
	return "Auth required"
}

func (e *ErrNotAuthenticated) Type() Type {
	return ErrNoAuth
}

type InternalError struct {
	Message string
	Layer   string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("Internal error occured. Message: '%s'. Layer: %s", e.Message, e.Layer)
}

type ErrTimeoutExceeded struct {
}

func (e *ErrTimeoutExceeded) Error() string {
	return "timeout exceeded"
}

func (e *ErrTimeoutExceeded) Type() Type {
	return ErrTimeout
}
