package database

import (
	"context"
	"log"
	"sync"
	"relay/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbInstance *pgxpool.Pool

var once sync.Once

func GetInstance() *pgxpool.Pool {
	
	once.Do(func() {
		connStr := config.ENV.DatabaseURL
		var err error
		
		config, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			log.Fatalf("Unable to parse database connection string: %v\n", err)
		}
		
		dbInstance, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v\n", err)
		}
		
		log.Println("Database connection pool created successfully.")
	})
	
	return dbInstance
}
