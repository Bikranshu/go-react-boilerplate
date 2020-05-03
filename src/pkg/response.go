package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type response struct {
	Body       *responseBody
	StatusCode int
}

type responseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	ErrorInvalidUser      = errors.New("invalid username or password")
	ErrorNotFound         = errors.New("record not found")
	ErrorExists           = errors.New("record already exists")
	ErrorInvalidParameter = errors.New("invalid parameter")
	ErrorInvalidSyntax    = errors.New("invalid syntax")
)

var ErrorHTTPStatusMap = map[string]int{
	ErrorInvalidUser.Error():      http.StatusBadRequest,
	ErrorNotFound.Error():         http.StatusNotFound,
	ErrorExists.Error():           http.StatusConflict,
	ErrorInvalidParameter.Error(): http.StatusBadRequest,
	ErrorInvalidSyntax.Error():    http.StatusBadRequest,
}

// ToJSON writes the response to the given http.ResponseWriter with an application/json Content-Type header.
func (r response) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r.Body)
}

// OK returns a successful response.
func OK(message string, data interface{}) *response {
	return newResponse(true, message, data, http.StatusOK)
}

// Fail returns a failed response.
func Fail(err error) *response {
	message := err.Error()
	statusCode := ErrorHTTPStatusMap[message]
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return newResponse(false, toSentenceCase(message), nil, statusCode)
}

func newResponse(success bool, message string, data interface{}, statusCode int) *response {
	return &response{
		Body: &responseBody{
			Success: success,
			Message: message,
			Data:    data,
		},
		StatusCode: statusCode,
	}
}

func toSentenceCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
