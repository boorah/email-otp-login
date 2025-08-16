package db

import (
	db "boorah/email-otp-login-backend/db/sqlc"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Queries *db.Queries

type Config struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
	DBNAME   string
	SSLMode  string
}

func loadConfig() *Config {
	return &Config{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USERNAME: os.Getenv("DB_USERNAME"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

}

func Connect(ctx context.Context) (*pgx.Conn, error) {
	config := loadConfig()

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.USERNAME, config.PASSWORD, config.HOST, config.PORT, config.DBNAME, config.SSLMode)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	Queries = db.New(conn)

	return conn, nil
}
