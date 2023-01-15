package dtos

import "errors"

type CreateUserDto struct {
	Name     string `json:"name" example:"Giovane"`
	Email    string `json:"email" example:"giovane@email.com"`
	Password string `json:"password" example:"1234"`
}

func (c CreateUserDto) ValidateFields() error {
	if c.Email == "" || c.Name == "" || c.Password == "" {
		return errors.New("insert all of the fields")
	}

	return nil
}
