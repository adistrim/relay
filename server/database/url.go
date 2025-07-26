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

func GetURL(pool *pgxpool.Pool, code string) (string, error) {
	var longURL string

	sqlStatement := `SELECT long_url FROM urls WHERE short_code = $1;`

	err := pool.QueryRow(context.Background(), sqlStatement, code).Scan(&longURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}
