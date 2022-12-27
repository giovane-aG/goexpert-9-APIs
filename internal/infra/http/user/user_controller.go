package user_controller

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"gorm.io/gorm"
)

type UserController struct {
	UserModel *database.User
}

func NewUserController(db *gorm.DB) *UserController {
	var userModel *database.User = &database.User{DB: db}
	return &UserController{UserModel: userModel}
}

func (c *UserController) CreateUser(name, email, password string) error {
	var user *entity.User
	var err error

	user, err = entity.NewUser(name, email, password)
	if err != nil {
		return err
	}

	c.UserModel.Create(user)
	return nil
}
