package errors

import (
	"errors"
)

var (
	ErrHttpNotFound     = errors.New("not found")
	ErrHttpInternal     = errors.New("internal error")
	ErrHttpBadRequest   = errors.New("bad request")
	ErrHttpUnauthorized = errors.New("unauthorized")
	ErrHttpForbidden    = errors.New("forbidden")
)
