package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.ConfigData.PORT),
		Handler: r,
	}

	go func() {
		log.Printf("Starting server on port %d\n", config.ConfigData.PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not start server: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exited gracefully")
}
