package user_controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/user/dtos"
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
		panic(err)
	}

	var parsedBody *dtos.CreateUserDto
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		panic(err)
	}

	var user *entity.User
	user, err = entity.NewUser(parsedBody.Name, parsedBody.Email, parsedBody.Password)
	if err != nil {
		panic(err)
	}

	err = c.UserDB.Create(user)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

	}

	response, _ := json.Marshal(map[string]string{"message": "User created successfully"})
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
