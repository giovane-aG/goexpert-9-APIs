package entity

import (
	"errors"

	"github.com/giovane-aG/goexpert/9-APIs/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id" example:"1c77fd61-4f6c-4ca1-8967-e5172e25c274"`
	Name     string    `json:"name" example:"giovane"`
	Password string    `json:"password" example:"1234"`
	Email    string    `json:"email" example:"giovane@email.com"`
}

func NewUser(name, email, password string) (*User, error) {

	if name == "" || email == "" || password == "" {
		return nil, errors.New("Please fill all of user fields")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		ID:       entity.NewID(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
