package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	user_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"net/http"
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
	var multiplexer *http.ServeMux
	portToString := fmt.Sprintf(":%v", port)
	defer http.ListenAndServe(portToString, multiplexer)

	userController := user_controller.NewUserController(db)

	multiplexer.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			// userController.CreateUser()
		}
	})

}

func main() {
	config := configs.LoadConfig("./")
	db := initDb(config)

	productModel := database.NewProduct(db)
	newProduct, err := entity.NewProduct("GALAX RTX 3060OC", 2500.00)
	if err != nil {
		panic(err)
	}

	err = newProduct.Validate()
	if err != nil {
		panic(err)
	}

	productModel.Create(newProduct)

	products, err := productModel.FindAll(2, 1, "")
	if err != nil {
		panic(err)
	}

	for _, v := range products {
		fmt.Println(v)
	}
}
