package errors

import (
	"fmt"
)

type Error struct {
	Err        error
	StatusCode int
	Message    string
}

func NewError(err error, statusCode int, message string) *Error {
	if err == nil {
		err = fmt.Errorf(message)
	}

	e := &Error{
		Err:        err,
		StatusCode: statusCode,
		Message:    message,
	}

	return e
}

func (e *Error) Error() string {
	return e.Err.Error()
}
