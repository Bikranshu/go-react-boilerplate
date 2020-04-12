package routes

import (
	"../handler"
	"../middleware"
	"../user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func InitRoute(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// middleware
	r.Use(middleware.CorsEveryWhere)
	r.Use(middleware.Authentication)
	r.Use(middleware.RequestLogger)

	userRepo := user.NewUserRepository(db)
	// auth
	authHandler := handler.NewAuthHandler(userRepo)
	r.HandleFunc("/v1/auth/login", authHandler.HandleLogin).Methods(http.MethodPost, http.MethodOptions)

	// user
	userHandler := handler.NewUserHandler(userRepo)
	r.HandleFunc("/v1/users", userHandler.HandleGetAll).Methods(http.MethodGet)
	r.HandleFunc("/v1/users", userHandler.HandleStore).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/v1/users/{id}", userHandler.HandleGetByID).Methods(http.MethodGet)
	r.HandleFunc("/v1/users/{id}", userHandler.HandleUpdate).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/v1/users/{id}/change-password", userHandler.HandleChangePassword).Methods(http.MethodPut, http.MethodOptions)

	return r
}
