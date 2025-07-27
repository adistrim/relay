package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/api", InitRouter)
	router.GET("/api/health", HealthCheck)
	router.POST("/api/shorten", CreateShortUrl)

	router.GET("/", ServeHomePage)
	router.GET("/:code", Forward)

	return router
}

func ServeHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func ServeErrorPage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"message": message,
	})
}
