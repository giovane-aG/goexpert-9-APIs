package database

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindById(id string) (*entity.User, error) {
	var user entity.User

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	u.DB.Where("id = ?", parsedID).Find(&user)
	return &user, nil
}
