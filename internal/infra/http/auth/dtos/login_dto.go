package dtos

import "errors"

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginDto) ValidateFields() error {
	if l.Email == "" || l.Password == "" {
		return errors.New("please insert email and password")
	}
	return nil
}
