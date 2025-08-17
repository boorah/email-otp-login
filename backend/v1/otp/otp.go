package otp

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"boorah/email-otp-login-backend/config"
	"boorah/email-otp-login-backend/db"
	sqlcConfig "boorah/email-otp-login-backend/db/sqlc"
	"boorah/email-otp-login-backend/dtos"
	helpers "boorah/email-otp-login-backend/helpers"
	"boorah/email-otp-login-backend/validator"
)

func RegisterRoutes(r chi.Router) chi.Router {
	otpRouter := chi.NewRouter()

	otpRouter.Post("/generate", func(w http.ResponseWriter, r *http.Request) {
		var payload *dtos.GenerateOTPRequest

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Println("error decoding request body:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))
			return
		}

		// Validate the body
		if err := validator.ValidateStruct(payload); err != nil {
			log.Println("error validating request body:", err)

			helpers.RespondWithError(w, helpers.NewValidationError("Invalid request body"))
			return
		}

		ctx := r.Context()

		// See if the email is already registered
		userData, err := db.Queries.GetUserByEmail(ctx, payload.Email)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("user not found, creating new user:", payload.Email)

			// Create the user in the database
			userData, err = db.Queries.CreateUser(ctx, sqlcConfig.CreateUserParams{
				ID:    pgtype.UUID{Bytes: [16]byte(uuid.New()), Valid: true},
				Email: payload.Email,
			})
			if err != nil {
				log.Println("error occurred while creating user", err)

				helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))

				return
			}

		} else if err != nil {
			log.Println("error while fetching user data by email:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))

			return
		}

		validityDuration := time.Duration(config.ConfigData.OTP_VALIDITY_MINUTES) * time.Minute
		expiryTime := time.Now().Add(validityDuration)

		otp := helpers.GenerateOTP()

		_, err = db.Queries.CreateUserOTP(ctx, sqlcConfig.CreateUserOTPParams{
			ID:     pgtype.UUID{Bytes: [16]byte(uuid.New()), Valid: true},
			UserID: userData.ID,
			Otp:    otp,
			ExpiresAt: pgtype.Timestamptz{
				Time:  expiryTime,
				Valid: true,
			},
		})

		// Send the OTP to the user's email
		err = helpers.SendOTPEmail(payload.Email, "OTP for Login", otp)
		if err != nil {
			log.Println("error while sending email:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))

			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message": "OTP generated",
		})
	})

	otpRouter.Post("/validate", func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body
		var payload *dtos.ValidateOTPRequest

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Println("error decoding request body:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))
			return
		}

		// Validate the body
		if err := validator.ValidateStruct(payload); err != nil {
			log.Println("error validating request body:", err)

			helpers.RespondWithError(w, helpers.NewValidationError("Invalid request body"))
			return
		}

		ctx := r.Context()

		// Check if the email is valid
		userData, err := db.Queries.GetUserByEmail(ctx, payload.Email)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("user not found:", payload.Email)

			helpers.RespondWithError(w, helpers.NewNotFoundError("User not found"))
			return
		} else if err != nil {
			log.Println("error while fetching user data by email:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))
			return
		}

		// Get the latest OTP for the user
		latestOTP, err := db.Queries.GetLatestUserOTP(ctx, userData.ID)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("no OTP found for user:", userData.Email)

			helpers.RespondWithError(w, helpers.NewUnauthorizedError("Invalid OTP or OTP expired"))
			return
		} else if err != nil {
			log.Println("error while fetching latest OTP for user:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))
			return
		}

		// Check if the OTP is valid - not expired and isn't used
		if latestOTP.Otp != payload.OTP || latestOTP.ExpiresAt.Time.Before(time.Now()) || latestOTP.UsedAt.Valid {
			helpers.RespondWithError(w, helpers.NewUnauthorizedError("Invalid OTP or OTP expired"))
			return
		}

		// Mark the OTP as used
		err = db.Queries.UpdateUserOTPUsedAt(ctx, latestOTP.ID)
		if err != nil {
			log.Println("error while marking OTP as used:", err)

			helpers.RespondWithError(w, helpers.NewInternalServerError("An internal error occurred"))
			return
		}

		// If valid, create a JWT token for the user
		jwtToken, err := helpers.GenerateJWT(userData.ID, config.ConfigData.JWT_VALIDITY_MINUTES)
		if err != nil {
			log.Println("error while generating JWT token:", err)

			helpers.RespondWithError(w, helpers.NewUnauthorizedError("An internal error occurred"))
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message": "OTP validated",
			"token":   jwtToken,
		})
	})

	return otpRouter
}
