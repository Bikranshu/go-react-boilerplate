package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func Authentication(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		publicEndpoints := []string{"/v1/auth/login", "/v1/auth/forgot"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                        //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range publicEndpoints {
			if value == requestPath {
				mux.ServeHTTP(w, r)
				return
			}
		}

		authorizationHeader := r.Header.Get("Authorization") //Grab the token from the header
		if len(authorizationHeader) <= 0 {
			http.Error(w, "No authorization header provided.", http.StatusForbidden)
			return
		}

		splitAuthorizationHeader := strings.Split(authorizationHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitAuthorizationHeader) != 2  || strings.ToLower(splitAuthorizationHeader[0]) != "bearer"{
			http.Error(w, "Authorization header format must be Bearer {token}.", http.StatusUnauthorized)
			return
		}

		tokenString := splitAuthorizationHeader[1] //Grab the token part, what we are truly interested in
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("jwt_secret")), nil
		})

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			http.Error(w, "Invalid authorization token.", http.StatusForbidden)
			return
		}

		mux.ServeHTTP(w, r) //proceed in the middleware chain

	})
}
