package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func ServeErrorPage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"message": message,
	})
}
