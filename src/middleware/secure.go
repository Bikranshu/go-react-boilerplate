package middleware

import (
	"net/http"
)

func SecureHeaders(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		mux.ServeHTTP(w, r)
	})
}
