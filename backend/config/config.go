package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST              string
	PORT              string
	USERNAME          string
	PASSWORD          string
	DBNAME            string
	SSLMode           string
	RESEND_API_KEY    string
	RESEND_FROM_NAME  string
	RESEND_FROM_EMAIL string
}

var ConfigData *Config

func LoadConfig() (*Config, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	if ConfigData == nil {
		ConfigData = &Config{
			HOST:              os.Getenv("DB_HOST"),
			PORT:              os.Getenv("DB_PORT"),
			USERNAME:          os.Getenv("DB_USERNAME"),
			PASSWORD:          os.Getenv("DB_PASSWORD"),
			DBNAME:            os.Getenv("DB_NAME"),
			SSLMode:           os.Getenv("DB_SSLMODE"),
			RESEND_API_KEY:    os.Getenv("RESEND_API_KEY"),
			RESEND_FROM_NAME:  os.Getenv("RESEND_FROM_NAME"),
			RESEND_FROM_EMAIL: os.Getenv("RESEND_FROM_EMAIL"),
		}
	}

	return ConfigData, nil
}
