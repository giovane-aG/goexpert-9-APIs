package auth_controller

import (
	"encoding/json"
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/auth/dtos"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
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
}
