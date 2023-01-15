package dtos

import "errors"

type LoginDto struct {
	Email    string `json:"email" example:"josvane@email.com"`
	Password string `json:"password" example:"1234"`
}

func (l *LoginDto) ValidateFields() error {
	if l.Email == "" || l.Password == "" {
		return errors.New("please insert email and password")
	}
	return nil
}
