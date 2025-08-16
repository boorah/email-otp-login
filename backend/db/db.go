package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"boorah/email-otp-login-backend/config"
	db "boorah/email-otp-login-backend/db/sqlc"
)

var Queries *db.Queries

func Connect(ctx context.Context) (*pgx.Conn, error) {
	config := *config.ConfigData

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.USERNAME, config.PASSWORD, config.HOST, config.PORT, config.DBNAME, config.SSLMode)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	Queries = db.New(conn)

	return conn, nil
}
