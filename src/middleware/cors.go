package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func CorsEveryWhere(mux http.Handler) http.Handler {
	return cors.Default().Handler(mux)
}
