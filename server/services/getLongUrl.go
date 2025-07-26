package services

import (
	"relay/database"
)

func GetLongUrl(code string) (string, error) {
	
	pool := database.GetInstance()
	
	url, err := database.GetURL(pool, code)
	if err != nil {
		return "", err
	}
	
	return url, nil
}
