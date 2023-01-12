package errors

import (
	"errors"
)

var (
	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("invalid price")
	ErrRequiredName  = errors.New("name is required")
	ErrRequiredID    = errors.New("id is required")
	ErrInvalidID     = errors.New("invalid id")
	ErrNotFound      = errors.New("product not found")
)

var (
	ErrRequiredEmail    = errors.New("email is required")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrRequiredPassword = errors.New("password is required")
	ErrRequiredUserName = errors.New("name is required")
)

type Error struct {
	Message string `json:"message"`
}
