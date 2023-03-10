//go:build wireinject
// +build wireinject

package main

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	user_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserDatabase(db *gorm.DB) *database.User {
	wire.Build(
		database.NewUser,
	)

	return &database.User{}
}

func InitializeUserController() *user_controller.UserController {
	wire.Build(
		// InitializeUserDatabase,
		database.NewUser,
		user_controller.NewUserController,
	)
	return &user_controller.UserController{}
}
