package auth

import (
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// auth session

		next.ServeHTTP(w, r)
	})
}
