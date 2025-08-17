package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV              string
	PORT                 int
	DB_HOST              string
	DB_PORT              int
	DB_USERNAME          string
	DB_PASSWORD          string
	DB_NAME              string
	SSL_MODE             string
	RESEND_API_KEY       string
	RESEND_FROM_NAME     string
	RESEND_FROM_EMAIL    string
	OTP_VALIDITY_MINUTES int
	JWT_SECRET           string
	JWT_VALIDITY_MINUTES int
}

var ConfigData *Config

func getEnvValue(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}

func getEnvValueInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func LoadConfig() (*Config, error) {
	// Load environment variables
	appEnv := getEnvValue("APP_ENV")

	// Load .env file only if the environment is local
	if appEnv == "local" {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	if ConfigData == nil {
		ConfigData = &Config{
			APP_ENV:              getEnvValue("APP_ENV"),
			PORT:                 getEnvValueInt("PORT", 8080),
			DB_HOST:              getEnvValue("DB_HOST"),
			DB_PORT:              getEnvValueInt("DB_PORT", 5432),
			DB_USERNAME:          getEnvValue("DB_USERNAME"),
			DB_PASSWORD:          getEnvValue("DB_PASSWORD"),
			DB_NAME:              getEnvValue("DB_NAME"),
			SSL_MODE:             getEnvValue("DB_SSLMODE"),
			RESEND_API_KEY:       getEnvValue("RESEND_API_KEY"),
			RESEND_FROM_NAME:     getEnvValue("RESEND_FROM_NAME"),
			RESEND_FROM_EMAIL:    getEnvValue("RESEND_FROM_EMAIL"),
			OTP_VALIDITY_MINUTES: getEnvValueInt("OTP_VALIDITY_MINUTES", 10),
			JWT_SECRET:           getEnvValue("JWT_SECRET"),
			JWT_VALIDITY_MINUTES: getEnvValueInt("JWT_VALIDITY_MINUTES", 30),
		}
	}

	return ConfigData, nil
}
