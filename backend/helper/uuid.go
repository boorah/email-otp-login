package helper

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetPGXUUID() (pgtype.UUID, error) {
	var pgxUUID pgtype.UUID
	err := pgxUUID.Scan(uuid.New().String())

	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("failed to generate pgx.UUID: %v", err)
	}

	return pgxUUID, nil
}
