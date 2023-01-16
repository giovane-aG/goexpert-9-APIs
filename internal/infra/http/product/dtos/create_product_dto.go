package dtos

import "errors"

type CreateProductDto struct {
	Name  string  `json:"name" example:"Galax RTX 3070"`
	Price float64 `json:"price" example:"3600.90"`
}

func (c *CreateProductDto) ValidateFields() error {
	if c.Name == "" || c.Price <= 0 {
		return errors.New("please insert valid fields")
	}
	return nil
}
