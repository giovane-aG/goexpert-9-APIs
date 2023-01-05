package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	auth_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/auth"
	user_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user"

	"github.com/go-chi/chi/v5"
)

func initDb(config *configs.Conf) *gorm.DB {
	var dsn string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func initServer(port int, db *gorm.DB) {
	r := chi.NewRouter()
	portToString := fmt.Sprintf(":%v", port)

	userDb := database.NewUser(db)
	userController := user_controller.NewUserController(*userDb)
	authController := auth_controller.NewAuthController(userDb, config.JWTSecret, config.JWTExpiresIn)

	r.Post("/user", userController.CreateUser)
	r.Get("/user/findByEmail/{email}", userController.FindByEmail)
	r.Get("/user/findById/{id}", userController.FindById)
	r.Put("/user/{id}", userController.Update)
	r.Delete("/user/{id}", userController.Delete)

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
