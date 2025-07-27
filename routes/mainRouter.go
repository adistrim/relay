package routes

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"relay/middleware"
	"relay/templates"
)

func MainRouter(redisClient *redis.Client) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.SetTrustedProxies(nil)
	
	generalRateLimiter := middleware.RateLimitMiddleware(redisClient, 30, 1*time.Minute)
	strictRateLimiter := middleware.RateLimitMiddleware(redisClient, 10, 1*time.Minute)

	router.Use(generalRateLimiter)

	tmpl := template.Must(template.ParseFS(templates.TemplatesFS, "html/*"))
	router.SetHTMLTemplate(tmpl)

	router.GET("/api", InitRouter)
	router.GET("/api/health", HealthCheck)
	router.POST("/api/shorten", CreateShortUrl, strictRateLimiter)

	router.GET("/", templates.ServeHomePage)
	router.GET("/:code", Forward)

	return router
}
