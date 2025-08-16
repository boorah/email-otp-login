package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"boorah/email-otp-login-backend/config"
	"boorah/email-otp-login-backend/db"
	"boorah/email-otp-login-backend/helper"
	v1 "boorah/email-otp-login-backend/v1"
)

func main() {
	// Load all environment variables
	_, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize the email client
	helper.InitEmailClient()

	// Connect to the database
	ctx := context.Background()
	conn, err := db.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("error connecting to database: %v", err))
	}

	defer conn.Close(ctx)

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "pong"}`))
	})

	r.Mount("/api/v1", v1.RegisterRoutes(r))

	log.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", r)

}
