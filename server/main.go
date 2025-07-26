package main

import (
	"relay/config"
	"relay/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/api/", routes.InitRouter)
	router.GET("/api/health", routes.HealthCheck)
	
	router.POST("/api/shorten", routes.CreateShortUrl)
	
	router.Run(":" + config.ENV.Port)
}
