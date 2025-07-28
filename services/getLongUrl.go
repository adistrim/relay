package services

import (
	"relay/database"
)

func GetLongUrl(code string) (string, error) {
	
	pool := database.GetInstance()
	redisClient := database.GetRedisInstance()
	
	url, err := database.GetURL(pool, redisClient, code)
	if err != nil {
		return "", err
	}
	
	return url, nil
}
