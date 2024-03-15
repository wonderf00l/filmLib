package app

import "fmt"

type initError struct {
	inner error
}

func (e *initError) Error() string {
	return fmt.Sprintf("Init error: %s\n", e.inner.Error())
}

type runError struct {
	inner error
}

func (e *runError) Error() string {
	return fmt.Sprintf("Run error: %s\n", e.inner.Error())
}
