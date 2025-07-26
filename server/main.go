package main

import (
	"relay/config"
	"relay/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	// system routes
	router.GET("/api", routes.InitRouter)
	router.GET("/api/health", routes.HealthCheck)
	
	// URL routes
	router.POST("/api/shorten", routes.CreateShortUrl)
	router.GET("/:code", routes.Forward)
	
	router.Run(":" + config.ENV.Port)
}
