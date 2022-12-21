package database

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	Find(id string) (*entity.User, error)
}
