package otp

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"boorah/email-otp-login-backend/db"
	queryTypes "boorah/email-otp-login-backend/db/sqlc"
	"boorah/email-otp-login-backend/dtos"
	"boorah/email-otp-login-backend/helper"
	"boorah/email-otp-login-backend/validator"
)

func RegisterRoutes(r chi.Router) chi.Router {
	otpRouter := chi.NewRouter()

	otpRouter.Post("/generate", func(w http.ResponseWriter, r *http.Request) {
		var payload *dtos.GenerateOTPRequest

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Println("error decoding request body:", err)

			helper.RespondWithError(w, helper.NewInternalServerError("An internal error occurred"))
			return
		}

		// Validate the body
		if err := validator.ValidateStruct(payload); err != nil {
			log.Println("error validating request body:", err)

			helper.RespondWithError(w, helper.NewValidationError("Invalid request body"))
			return
		}

		ctx := r.Context()

		// See if the email is already registered
		_, err := db.Queries.GetUserByEmail(ctx, payload.Email)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("user not found, creating new user:", payload.Email)

			// Create a new user in the database
			userId, err := helper.GetPGXUUID()
			if err != nil {
				log.Println("error while generating uuid:", err)

				helper.RespondWithError(w, helper.NewInternalServerError("An internal error occurred"))

				return
			}

			// Create the user in the database
			_, err = db.Queries.CreateUser(ctx, queryTypes.CreateUserParams{
				ID:    userId,
				Email: payload.Email,
			})
			if err != nil {
				log.Println("error occurred while creating user", err)

				helper.RespondWithError(w, helper.NewInternalServerError("An internal error occurred"))

				return
			}

		} else if err != nil {
			log.Println("error while fetching user data by email:", err)

			helper.RespondWithError(w, helper.NewInternalServerError("An internal error occurred"))

			return
		}

		helper.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message": "OTP generated",
		})
	})

	otpRouter.Post("/validate", func(w http.ResponseWriter, r *http.Request) {
		helper.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message": "OTP validated",
		})
	})

	return otpRouter
}
