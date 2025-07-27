package routes

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"relay/templates"
)

func MainRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	tmpl := template.Must(template.ParseFS(templates.TemplatesFS, "html/*"))
	router.SetHTMLTemplate(tmpl)

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
