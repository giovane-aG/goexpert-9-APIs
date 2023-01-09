package auth_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/auth/dtos"
	"github.com/go-chi/jwtauth"
)

type AuthController struct {
	UserDB       *database.User
	JwtSecret    string
	JwtExpiresIn int
}

type AcessToken struct {
	Token string `json:"token"`
}

func NewAuthController(userDb *database.User, jwtExpiresIn int) *AuthController {
	return &AuthController{
		UserDB:       userDb,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto *dtos.LoginDto = &dtos.LoginDto{}
	json.NewDecoder(r.Body).Decode(loginDto)

	err := loginDto.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := a.UserDB.FindByEmail(loginDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	if !user.ValidatePassword(loginDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("incorrect password")
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
		return
	}

	accessToken := AcessToken{
		Token: token,
	}

	json.NewEncoder(w).Encode(accessToken)
}
