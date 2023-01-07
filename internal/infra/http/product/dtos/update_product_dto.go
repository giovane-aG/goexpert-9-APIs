package dtos

import "errors"

type UpdateProductDto struct {
	Name  string
	Price float64
}

func (u *UpdateProductDto) ValidateFields() error {
	if u.Name == "" && u.Price <= 0 {
		return errors.New("please insert valid fields")
	}
	return nil
}
