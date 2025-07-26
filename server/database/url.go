package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)


func SaveURL(pool *pgxpool.Pool, longURL string) (string, error) {
	var shortCode string

	sqlStatement := `INSERT INTO urls (long_url) VALUES ($1) RETURNING short_code;`

	err := pool.QueryRow(context.Background(), sqlStatement, longURL).Scan(&shortCode)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}
