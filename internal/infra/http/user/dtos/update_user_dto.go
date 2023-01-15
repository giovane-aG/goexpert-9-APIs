package dtos

import "errors"

type UpdateUserDto struct {
	Name  string `json:"name" example:"Giovane"`
	Email string `json:"email" example:"giovane@email.com"`
}

func (u UpdateUserDto) ValidateFields() error {
	if u.Email == "" && u.Name == "" {
		return errors.New("insert at least one of the fields")
	}

	return nil
}
