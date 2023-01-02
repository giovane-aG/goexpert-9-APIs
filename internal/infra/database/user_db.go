package database

import (
	"errors"

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
	tx := u.DB.Where("email = ?", email).Find(&user)
	rows := tx.RowsAffected

	if tx.Error != nil {
		return nil, tx.Error
	}

	if rows == 0 {
		return nil, nil
	}

	return &user, nil
}

func (u *User) FindById(id string) (*entity.User, error) {
	var user entity.User

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	tx := u.DB.Where("id = ?", parsedID).Find(&user)

	if tx.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func (u *User) Update(user *entity.User) error {
	user, err := u.FindById(user.ID.String())
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("No user found with this ID")
	}

	return u.DB.Save(user).Error
}
