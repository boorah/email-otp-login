package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"boorah/email-otp-login-backend/db"
	v1 "boorah/email-otp-login-backend/v1"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("error loading .env file: %v", err))
	}

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

	log.Println("Server running on http://localhost:8080")

	r.Mount("/api/v1", v1.RegisterRoutes(r))

	http.ListenAndServe(":8080", r)

}
