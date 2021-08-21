package gazuberlandia

import (
	"errors"
	"fmt"
)

const (
	CONFLICT            string = "conflict"
	INTERNAL            string = "internal"
	INVALID             string = "invalid"
	NOTFOUND            string = "not_found"
	NOTIMPLEMENTED      string = "not_implemented"
	UNAUTHORIZED        string = "unauthorized"
	UNPROCESSABLEENTITY string = "unprocessableentity"
)

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error: code=%s message=%s err=%v", e.Code, e.Message, e.Err)
}

func ErrorCode(err error) string {
	var e *AppError

	if err == nil {
		return ""
	}

	if errors.As(err, &e) {
		return e.Code
	}

	return INTERNAL
}

func ErrorMessage(err error) string {
	var e *AppError

	if err == nil {
		return ""
	}

	if errors.As(err, &e) {
		return e.Message
	}

	return "Internal error."
}

func ErrorF(code string, format string, err error, args ...interface{}) *AppError {
	return &AppError{
		Code:    code,
		Err:     err,
		Message: fmt.Sprintf(format, args...),
	}
}
