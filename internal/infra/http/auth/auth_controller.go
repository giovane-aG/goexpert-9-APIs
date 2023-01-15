package auth_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/errors"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/auth/dtos"
	"github.com/go-chi/jwtauth"
)

type AuthController struct {
	UserDB       *database.User
	JwtSecret    string
	JwtExpiresIn int
}

func NewAuthController(userDb *database.User, jwtExpiresIn int) *AuthController {
	return &AuthController{
		UserDB:       userDb,
		JwtExpiresIn: jwtExpiresIn,
	}
}

// @Summary 		Login user
// @Description	Login user
// @Tags				auth
// @Accept			json
// @Produce 		json
// @Param 			request	body dtos.LoginDto true	"user request"
// @Success			200	{object}	dtos.AcessToken
// @Failure			404	{object}	errors.Error
// @Router			/auth/login	[post]
func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	jsonEncoder := json.NewEncoder(w)
	var loginDto *dtos.LoginDto = &dtos.LoginDto{}
	json.NewDecoder(r.Body).Decode(loginDto)

	err := loginDto.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder.Encode(errors.Error{Message: err.Error()})
		return
	}

	user, err := a.UserDB.FindByEmail(loginDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(errors.Error{Message: err.Error()})
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(errors.Error{Message: "incorret email or password"})
	}

	if !user.ValidatePassword(loginDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		jsonEncoder.Encode(errors.Error{Message: "incorret email or password"})
		return
	}

	w.WriteHeader(http.StatusOK)
	tokenAuth := r.Context().Value("jwtAuth").(*jwtauth.JWTAuth)
	_, token, err := tokenAuth.Encode(map[string]interface{}{
		"user_id": user.ID.String(),
		"email":   user.Email,
	})

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(errors.Error{Message: err.Error()})
		return
	}

	accessToken := dtos.AcessToken{
		Token: token,
	}

	json.NewEncoder(w).Encode(accessToken)
}
