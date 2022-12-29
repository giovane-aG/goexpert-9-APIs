package user_controller

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
)

type UserController struct {
	UserDB database.UserInterface
}

func NewUserController(userDB database.User) *UserController {
	var userModel *database.User = database.NewUser(userDB.DB)
	return &UserController{UserDB: userModel}
}

func (c *UserController) CreateUser(name, email, password string) error {
	var user *entity.User
	var err error

	user, err = entity.NewUser(name, email, password)
	if err != nil {
		return err
	}

	c.UserDB.Create(user)
	return nil
}
