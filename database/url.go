package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
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

func GetURL(pool *pgxpool.Pool, redisClient *redis.Client, code string) (string, error) {
	ctx := context.Background()
	cacheKey := "url:" + code
	var longURL string
	var err error
	
	longURL, err = redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		return longURL, nil
	}

	sqlStatement := `SELECT long_url FROM urls WHERE short_code = $1;`
	err = pool.QueryRow(ctx, sqlStatement, code).Scan(&longURL)
	if err != nil {
		return "", err
	}
	
	redisClient.Set(ctx, cacheKey, longURL, 24*time.Hour).Err()
	
	return longURL, nil
}
