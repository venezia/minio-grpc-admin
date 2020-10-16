package errors

import "fmt"

const (
	ErrNotFound     ErrorCode = "NotFound"
	ErrUnauthorized ErrorCode = "NotAuthorized"
	ErrOther        ErrorCode = "Other"
)

type ErrorCode string

type Error struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e Error) Error() string {
	output := fmt.Sprintf("Code %s", e.Code)
	if e.Message != "" {
		output += " - " + e.Message
	}
	return output
}

func (e Error) Unwrap() error {
	return e.Err
}

func IsError(err error) bool {
	if _, ok := err.(Error); ok {
		return true
	}
	return false
}

func New(code ErrorCode, message string, err error) error {
	return Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
