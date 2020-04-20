package handler

import (
	"../auth"
	"../pkg"
	"../user"
	"encoding/json"
	"net/http"
)

type authHandler struct {
	service auth.Service
}

func NewAuthHandler(repo user.URepository) *authHandler {
	return &authHandler{service: auth.NewAuthService(repo)}
}

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Login struct {
	Token string `json:"token"`
}

// Login godoc
// @Summary  Authenticate a user and receive a JWT Token
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param Credentials body Credentials true "Credentials"
// @Success 200 {object} Login
// @Router /v1/auth/login [post]
func (ah authHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		pkg.Wrap(err, w)
		return
	}

	token, err := ah.service.Login(r.Context(), credentials.Email, credentials.Password)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
	return
}
