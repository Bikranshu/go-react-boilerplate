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

// ListUsers godoc
// @Summary List all users
// @Description get users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} user.User
// @Security BearerAuth
// @Router /v1/users [get]
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

// ListUser godoc
// @Summary Find user by ID
// @Description get user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} user.User
// @Router  /v1/users/{id} [get]
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

// CreateUser godoc
// @Summary Add a new user
// @Description create a new user
// @Tags User
// @Accept  json
// @Produce  json
// @Param  user body user.User true "Create user"
// @Success 200 {object} user.User
// @Router /v1/users [post]
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

// UpdateUser godoc
// @Summary Update an existing user
// @Description update an existing user by ID
// @ID int
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body user.User true "Update user"
// @Success 200 {object} user.User
// @Router /v1/users/{id} [put]
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

// UpdateUserPassword godoc
// @Summary Update an existing user password by ID
// @Description update an existing user password by ID
// @ID int
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body user.User true "Update user password"
// @Success 200 {string} string "Password changed successfully"
// @Router /v1/users/{id}/change-password [put]
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
