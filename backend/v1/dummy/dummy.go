package dummy

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"boorah/email-otp-login-backend/helpers"
	"boorah/email-otp-login-backend/middlewares"
)

func RegisterRoutes(r chi.Router) chi.Router {
	dummyRouter := chi.NewRouter()

	dummyRouter.Use(middlewares.ValidateJWT)

	dummyRouter.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message": "This is dummy data",
		})
	})

	return dummyRouter
}
