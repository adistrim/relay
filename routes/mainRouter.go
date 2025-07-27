package routes

import (
	"html/template"

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

	router.GET("/", templates.ServeHomePage)
	router.GET("/:code", Forward)

	return router
}
