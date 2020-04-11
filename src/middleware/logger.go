package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RequestLogger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		mux.ServeHTTP(w, r)
	})
}
