package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"boorah/email-otp-login-backend/config"
	"boorah/email-otp-login-backend/db"
	helpers "boorah/email-otp-login-backend/helpers"
	v1 "boorah/email-otp-login-backend/v1"
)

func main() {
	// Load all environment variables
	_, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize the email client
	helpers.InitEmailClient()

	// Connect to the database
	ctx := context.Background()
	pool, err := db.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("error connecting to database: %v", err))
	}

	defer pool.Close()

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "pong"}`))
	})

	r.Mount("/api/v1", v1.RegisterRoutes(r))

	log.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(fmt.Sprintf("localhost:%d", config.ConfigData.PORT), r)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
