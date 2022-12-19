package entity

import (
	"errors"

	"github.com/giovane-aG/goexpert/9-APIs/pkg/entity"
)

var (
	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("invalid price")
	ErrRequiredName  = errors.New("name is required")
	ErrRequiredID    = errors.New("id is required")
	ErrInvalidID     = errors.New("invalid id")
	ErrNotFound      = errors.New("product not found")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
}

func ValidateProduct(p *Product) error {
	if p.ID.String() == "" {
		return ErrRequiredID
	}
	if p.Name == "" {
		return ErrRequiredName
	}
	if p.Price == 0 {
		return ErrRequiredPrice
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
