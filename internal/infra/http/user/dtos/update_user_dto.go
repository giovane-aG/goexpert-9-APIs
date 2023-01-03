package dtos

import "errors"

type UpdateUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UpdateUserDto) ValidateFields() error {
	if u.Email == "" && u.Name == "" {
		return errors.New("insert at least one of the fields")
	}

	return nil
}
