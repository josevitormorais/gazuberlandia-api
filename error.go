package gazuberlandia

import (
	"errors"
	"fmt"
)

const (
	CONFLICT       = "conflict"
	INTERNAL       = "internal"
	INVALID        = "invalid"
	NOTFOUND       = "not_found"
	NOTIMPLEMENTED = "not_implemented"
	UNAUTHORIZED   = "unauthorized"
)

type Error struct {
	Code    string
	Message string
	Err     error
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: code=%s message=%s err=%v", e.Code, e.Message, e.Err)
}

func ErrorCode(err error) string {
	var e *Error

	if err == nil {
		return ""
	}

	if errors.As(err, &e) {
		return e.Code
	}

	return INTERNAL
}

func ErrorMessage(err error) string {
	var e *Error

	if err == nil {
		return ""
	}

	if errors.As(err, &e) {
		return e.Message
	}

	return "Internal error."
}
