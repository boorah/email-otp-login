package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST                 string
	PORT                 int
	USERNAME             string
	PASSWORD             string
	DBNAME               string
	SSLMode              string
	RESEND_API_KEY       string
	RESEND_FROM_NAME     string
	RESEND_FROM_EMAIL    string
	OTP_VALIDITY_MINUTES int
	JWT_SECRET           string
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
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	if ConfigData == nil {
		ConfigData = &Config{
			HOST:                 getEnvValue("DB_HOST"),
			PORT:                 getEnvValueInt("DB_PORT", 5432),
			USERNAME:             getEnvValue("DB_USERNAME"),
			PASSWORD:             getEnvValue("DB_PASSWORD"),
			DBNAME:               getEnvValue("DB_NAME"),
			SSLMode:              getEnvValue("DB_SSLMODE"),
			RESEND_API_KEY:       getEnvValue("RESEND_API_KEY"),
			RESEND_FROM_NAME:     getEnvValue("RESEND_FROM_NAME"),
			RESEND_FROM_EMAIL:    getEnvValue("RESEND_FROM_EMAIL"),
			OTP_VALIDITY_MINUTES: getEnvValueInt("OTP_VALIDITY_MINUTES", 10),
			JWT_SECRET:           getEnvValue("JWT_SECRET"),
		}
	}

	return ConfigData, nil
}
