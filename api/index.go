package handler

import (
	"net/http"
	"relay/database"
	"relay/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	redisClient := database.GetRedisInstance()
	router := routes.MainRouter(redisClient)
	router.ServeHTTP(w, r)
}
