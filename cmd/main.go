package main

import (
	"fmt"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	auth_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/auth"
	product_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/product"
	user_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

func initDb(config *configs.Conf) *gorm.DB {
	var dsn string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initServer(port int, db *gorm.DB) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	portToString := fmt.Sprintf(":%v", port)
	tokenAuth := jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	userDb := database.NewUser(db)
	productDb := database.NewProduct(db)

	userController := user_controller.NewUserController(*userDb)
	productController := product_controller.NewProductController(productDb)
	authController := auth_controller.NewAuthController(userDb, config.JWTSecret, config.JWTExpiresIn)

	r.Route("/user", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", userController.CreateUser)
		r.Get("/findByEmail/{email}", userController.FindByEmail)
		r.Get("/findById/{id}", userController.FindById)
		r.Put("/{id}", userController.Update)
		r.Delete("/{id}", userController.Delete)
	})

	r.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productController.Create)
	})

	r.Post("/auth/login", authController.Login)
	http.ListenAndServe(portToString, r)
}

var config *configs.Conf
var db *gorm.DB

func main() {
	config = configs.LoadConfig("./")
	db = initDb(config)
	initServer(8080, db)
}
