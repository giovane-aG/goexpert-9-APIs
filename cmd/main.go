package main

import (
	"fmt"
	"log"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	auth_controller "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/auth"

	_ "github.com/giovane-aG/goexpert/9-APIs/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Go Expert API
//	@version		1.0
//	@description	Production API authentication.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Giovane Aguiar
//	@contact.email	giovaneaalmeida27@gmail.com

//	@host		localhost:8080
//	@BasePath	/
//  @securityDefinitions.apiKey ApiKeyAuth
//  @in header
//  @name Authorization

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
	r.Use(loggerMiddleware)
	r.Use(middleware.Recoverer)

	portToString := fmt.Sprintf(":%v", port)
	tokenAuth := jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	userDb := database.NewUser(db)

	userController := NewUserController(*database.NewUser(db))
	productController := NewProductController(database.NewProduct(db))

	r.Use(middleware.WithValue("jwtAuth", tokenAuth))
	authController := auth_controller.NewAuthController(userDb, config.JWTExpiresIn)

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
		r.Get("/findAll", productController.FindAll)
		r.Get("/findById/{id}", productController.FindById)
		r.Put("/{id}", productController.Update)
		r.Delete("/{id}", productController.Delete)
	})

	r.Post("/auth/login", authController.Login)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	http.ListenAndServe(portToString, r)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: IP -> %v | %v -> %v", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

var config *configs.Conf
var db *gorm.DB

func main() {
	config = configs.LoadConfig("./")
	db = initDb(config)
	initServer(8080, db)
}
