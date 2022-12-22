package database

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	Find(id string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
