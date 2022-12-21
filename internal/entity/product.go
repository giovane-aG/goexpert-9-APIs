package entity

import (
	"time"

	"github.com/giovane-aG/goexpert/9-APIs/internal/errors"
	"github.com/giovane-aG/goexpert/9-APIs/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errors.ErrRequiredID
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errors.ErrInvalidID
	}
	if p.Name == "" {
		return errors.ErrRequiredName
	}
	if p.Price == 0 {
		return errors.ErrRequiredPrice
	}
	if p.Price < 0 {
		return errors.ErrInvalidPrice
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}
