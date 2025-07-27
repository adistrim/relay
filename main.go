package main

import (
	"relay/config"
	"relay/database"
	"relay/routes"
)

func main() {
	redisClient := database.GetRedisInstance()
	router := routes.MainRouter(redisClient)
	router.Run(":" + config.ENV.Port)
}
