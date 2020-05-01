package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"strings"
)

type JSONResponse struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Bad Request"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

func JSONWrapper(w http.ResponseWriter, statusCode int, data interface{}) {
	jsonResponse := JSONResponse{
		Message: "Success",
		Code:    statusCode,
		Data:    data,
	}
	json.NewEncoder(w).Encode(jsonResponse)
}

func ErrorWrapper(w http.ResponseWriter, statusCode int, err error) {
	message := err.Error()
	w.WriteHeader(statusCode)

	errResponse := ErrorResponse{
		Message: message,
		Code:    statusCode,
	}
	log.WithFields(log.Fields{
		"message": message,
		"code":    statusCode,
	}).Error("Error occurred")

	json.NewEncoder(w).Encode(errResponse)
}

// RemoveHiddenFields if the given model is a struct pointer.
// All fields marked with the tag `model:"hide"` will be
// set to their zero value.
//
// For example, this allows to send user models to the client
// without their password field.
func RemoveHiddenFields(model interface{}) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Struct {
			value := reflect.ValueOf(model).Elem()
			for i := 0; i < t.NumField(); i++ {
				field := value.Field(i)
				fieldType := t.Field(i)

				if !field.CanSet() {
					continue
				}

				if field.Kind() == reflect.Struct && fieldType.Anonymous {
					// Check promoted fields recursively
					RemoveHiddenFields(field.Addr().Interface())
					continue
				}

				tag := strings.Split(fieldType.Tag.Get("model"), ";")
				if Contains(tag, "hide") {
					field.Set(reflect.Zero(fieldType.Type))
				}
			}
		}
	}
}
