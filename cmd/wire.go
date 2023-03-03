//go:build wireinject
// +build wireinject

package main

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	user_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user"
	"github.com/google/wire"
)

func NewUserController(database.User) *user_controller.UserController {

	wire.Build(
		user_controller.NewUserController,
	)

	return &user_controller.UserController{}
}
