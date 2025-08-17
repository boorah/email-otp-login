package v1

import (
	"boorah/email-otp-login-backend/v1/dummy"
	"boorah/email-otp-login-backend/v1/otp"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) chi.Router {
	v1Router := chi.NewRouter()

	v1Router.Mount("/otp", otp.RegisterRoutes(v1Router))
	v1Router.Mount("/dummy", dummy.RegisterRoutes(v1Router))

	return v1Router
}
