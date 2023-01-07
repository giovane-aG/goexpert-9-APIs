package dtos

import "errors"

type CreateProductDto struct {
	Name  string
	Price float64
}

func (c *CreateProductDto) ValidateFields() error {
	if c.Name == "" || c.Price <= 0 {
		return errors.New("please insert valid fields")
	}
	return nil
}
