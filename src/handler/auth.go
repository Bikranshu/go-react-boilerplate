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
