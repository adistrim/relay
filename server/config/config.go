package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
}

var ENV *Config

func init() {
	var err error
	ENV, err = Load()
	if err != nil {
		log.Fatal(err)
	}
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	
	port := os.Getenv("PORT")
	if port == "" { 
		log.Println("PORT not set in environment variables, using default port 8080")
		port = "8080" 
	}
	
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" { return nil, fmt.Errorf("DATABASE_URL not set in environment variables") }
	
	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}, nil
}
