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
