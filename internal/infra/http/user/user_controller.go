package user_controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/errors"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	http_errors "github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/errors"
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

// @Summary 		Create new user
// @Description	Create new user
// @Tags			users
// @Accept			json
// @Produce 		json
// @Param 			request	body 		dtos.CreateUserDto	true	"user request"
// @Success			201
// @Failure			500		{object}	errors.Error
// @Router			/user	[post]
// @Security		ApiKeyAuth
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	jsonEnconder := json.NewEncoder(w)
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
		jsonEnconder.Encode(errors.Error{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// creating entity
	var user *entity.User
	user, err = entity.NewUser(parsedBody.Name, parsedBody.Email, parsedBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonEnconder.Encode(errors.Error{Message: err.Error()})
	}

	// saving entity
	err = c.UserDB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonEnconder.Encode(errors.Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsonEnconder.Encode(errors.Error{Message: "User created successfully"})
	return
}

// @Summary 		Find user by email
// @Description	Find user by email
// @Tags				users
// @Accept			json
// @Produce 		json
// @Param 			email path string true "the email of the user"
// @Success			200	{object}	entity.User
// @Failure			500	{object}	errors.Error
// @Failure			400	{object}	errors.Error
// @Router			/user/findByEmail/{email}	[get]
// @Security		ApiKeyAuth
func (c *UserController) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	jsonEncoder := json.NewEncoder(w)
	email, err := url.QueryUnescape(email)

	user, err := c.UserDB.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(errors.Error{
			Message: "No user with that email was found",
		})
		return
	}

	jsonEncoder.Encode(user)
	return
}

// @Summary 		Find user by id
// @Description	Find user by id
// @Tags				users
// @Produce 		json
// @Param 			id path string true "the id of the user"
// @Success			200	{object}	entity.User
// @Failure			500	{object}	errors.Error
// @Failure			400	{object}	errors.Error
// @Router			/user/findById/{id}	[get]
// @Security		ApiKeyAuth
func (c *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	jsonEncoder := json.NewEncoder(w)
	user, err := c.UserDB.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(errors.Error{
			Message: err.Error(),
		})
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(errors.Error{Message: "No user with that id was found"})
		return
	}

	jsonEncoder.Encode(user)
	return
}

// @Summary 		Update user
// @Description	Update user
// @Tags				users
// @Produce 		json
// @Param 			id path string true "the id of the user"
// @Param				request	body	dtos.UpdateUserDto true "payload to update the user"
// @Success			200
// @Failure			500	{object}	errors.Error
// @Failure			404	{object}	errors.Error
// @Router			/user/{id}	[put]
// @Security		ApiKeyAuth
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := c.UserDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var parsedBody *dtos.UpdateUserDto = &dtos.UpdateUserDto{}
	json.Unmarshal(body, &parsedBody)
	err = parsedBody.ValidateFields()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := fmt.Sprintf("message: %v", err)
		parsedMessage, _ := json.Marshal(map[string]string{"message": message})

		w.Write(parsedMessage)
		return
	}

	if parsedBody.Name != "" {
		user.Name = parsedBody.Name
	}
	if parsedBody.Email != "" {
		user.Email = parsedBody.Email
	}

	err = c.UserDB.Update(user)
}

// @Summary 		Delete user
// @Description	Delete user
// @Tags				users
// @Produce 		json
// @Param 			id path string true "the id of the user"
// @Success			200
// @Failure			500	{object}	errors.Error
// @Failure			404	{object}	errors.Error
// @Router			/user/{id}	[delete]
// @Security		ApiKeyAuth
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := c.UserDB.Delete(id)
	if err != nil {
		if err == http_errors.ErrHttpNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
