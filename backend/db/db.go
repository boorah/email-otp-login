package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"boorah/email-otp-login-backend/config"
	sqlcConfig "boorah/email-otp-login-backend/db/sqlc"
)

var Queries *sqlcConfig.Queries

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	config := *config.ConfigData

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME, config.SSL_MODE)

	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	Queries = sqlcConfig.New(pool)

	return pool, nil
}
