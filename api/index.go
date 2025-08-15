package handler

/*

This file is exists here to be compatible with the vercel serverless project structure.

https://vercel.com/docs/functions/runtimes/go

It is not used in the development.

*/

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
