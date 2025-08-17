package middlewares

import (
	"net/http"
	"strings"

	"boorah/email-otp-login-backend/helpers"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")

		if authorizationHeader == "" {
			helpers.RespondWithError(w, helpers.NewUnauthorizedError("Authorization header is required"))
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			helpers.RespondWithError(w, helpers.NewUnauthorizedError("Invalid authorization header format"))
			return
		}

		token := strings.TrimPrefix(authorizationHeader, "Bearer ")

		// Validate the JWT token
		_, err := helpers.ValidateJWT(token)
		if err != nil {
			helpers.RespondWithError(w, helpers.NewUnauthorizedError("Invalid or expired token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
