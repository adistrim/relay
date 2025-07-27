package services

import (
	"relay/database"
)

func GetShortUrl(url string) (string, error) {
	pool := database.GetInstance()

	shortCode, err := database.SaveURL(pool, url)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}
