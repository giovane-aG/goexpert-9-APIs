//go:build wireinject
// +build wireinject

package main

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	product_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/product"
	user_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user"
	"github.com/google/wire"
)

func NewUserController(database.User) *user_controller.UserController {

	wire.Build(
		user_controller.NewUserController,
	)

	return &user_controller.UserController{}
}

func NewProductController(*database.Product) *product_controller.ProductController {
	wire.Build(
		product_controller.NewProductController,
	)
	return &product_controller.ProductController{}
}
