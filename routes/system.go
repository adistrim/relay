package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "relay server",
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

