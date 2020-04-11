package pkg

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var (
	ErrNotFound         = errors.New("Record not found.")
	ErrNoContent        = errors.New("Record not found.")
	ErrExists           = errors.New("Record already exists in system.")
	ErrDatabase         = errors.New("Database error occured.")
	ErrUnauthorized     = errors.New("You are not allowed to perform this action.")
	ErrForbidden        = errors.New("Forbidden")
	ErrMethodNotAllowed = errors.New("Method is not allowed.")
	ErrInvalidToken     = errors.New("Invalid authorization token.")
	ErrInvalidParameter = errors.New("Invalid parameter.")
	ErrInvalidUser      = errors.New("Invalid username or password.")
)

var ErrHTTPStatusMap = map[string]int{
	ErrNotFound.Error():         http.StatusNotFound,
	ErrExists.Error():           http.StatusConflict,
	ErrNoContent.Error():        http.StatusNotFound,
	ErrDatabase.Error():         http.StatusInternalServerError,
	ErrUnauthorized.Error():     http.StatusUnauthorized,
	ErrInvalidUser.Error():      http.StatusUnauthorized,
	ErrForbidden.Error():        http.StatusForbidden,
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrInvalidToken.Error():     http.StatusForbidden,
	ErrInvalidParameter.Error(): http.StatusBadRequest,
}

func Wrap(err error, w http.ResponseWriter) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	// If error code is not found
	if code == 0 {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)

	errResponse := ErrorResponse{
		Message: msg,
		Status:  code,
	}
	log.WithFields(log.Fields{
		"message": msg,
		"code":    code,
	}).Error("Error occurred")

	json.NewEncoder(w).Encode(errResponse)
}
