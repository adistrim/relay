package routes

import (
	"github.com/gin-gonic/gin"
)

func MainRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api", InitRouter)
	router.GET("/api/health", HealthCheck)

	router.POST("/api/shorten", CreateShortUrl)
	router.GET("/:code", Forward)

	return router
}
