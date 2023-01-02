package user_controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user/dtos"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserDB database.UserInterface
}

func NewUserController(userDB database.User) *UserController {
	var userModel *database.User = database.NewUser(userDB.DB)
	return &UserController{UserDB: userModel}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// parsing body
	var parsedBody *dtos.CreateUserDto = &dtos.CreateUserDto{}
	json.Unmarshal(body, &parsedBody)
	err = parsedBody.ValidateFields()

	if err != nil {
		message := fmt.Sprintf("message: %v", err)
		parsedMessage, _ := json.Marshal(map[string]string{"message": message})

		w.Write(parsedMessage)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// creating entity
	var user *entity.User
	user, err = entity.NewUser(parsedBody.Name, parsedBody.Email, parsedBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// saving entity
	err = c.UserDB.Create(user)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(map[string]string{"message": "User created successfully"})
	w.Write(response)
	w.WriteHeader(http.StatusBadRequest)
}

func (c *UserController) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	fmt.Println(email)

	user, err := c.UserDB.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
	return
}
