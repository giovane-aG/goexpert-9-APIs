package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
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

type createUserDto struct {
	Name     string
	Email    string
	Password string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userController := user_controller.NewUserController(db)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var parsedBody *createUserDto
		err = json.Unmarshal(body, &parsedBody)
		if err != nil {
			panic(err)
		}

		err = userController.CreateUser(parsedBody.Name, parsedBody.Email, parsedBody.Password)
		if err != nil {
			response, _ := json.Marshal(map[string]string{"message": err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		response, _ := json.Marshal(map[string]string{"message": "User created successfully"})
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func initServer(port int, db *gorm.DB) {
	var multiplexer http.ServeMux
	portToString := fmt.Sprintf(":%v", port)

	multiplexer.HandleFunc("/user", CreateUser)
	http.ListenAndServe(portToString, &multiplexer)
}

var config *configs.Conf
var db *gorm.DB

func main() {
	config = configs.LoadConfig("./")
	db = initDb(config)
	initServer(8080, db)
}
