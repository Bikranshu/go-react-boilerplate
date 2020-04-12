package handler

import (
	"../pkg"
	"../user"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(repo user.URepository) *userHandler {
	return &userHandler{service: user.NewUserService(repo)}
}

func (uh userHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	u, err := uh.service.FindByAll(r.Context())
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": u,
	})

	return
}

func (uh userHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	u, err := uh.service.FindByID(r.Context(), uint(id))
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": u,
	})
	return
}

func (uh userHandler) HandleStore(w http.ResponseWriter, r *http.Request) {
	userModel := user.User{}
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		pkg.Wrap(err, w)
		return
	}

	u, err := uh.service.Insert(r.Context(), userModel)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   u,
		"message": "User created successfully.",
	})
	return
}

func (uh userHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	userModel := user.User{}
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		pkg.Wrap(err, w)
		return
	}

	u, err := uh.service.Update(r.Context(), uint(id), userModel)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   u,
		"message": "User updated successfully.",
	})
	return
}

func (uh userHandler) HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	userModel := user.User{}
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		pkg.Wrap(err, w)
		return
	}

	err = uh.service.ChangePassword(r.Context(), uint(id), userModel.Email, userModel.Password)
	if err != nil {
		pkg.Wrap(err, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Password changed successfully.",
	})
	return
}
