package middleware

import (
	"context"
	"relay/templates"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()
		key := "rate-limit:" + ip

		pipe := redisClient.Pipeline()
		count := pipe.Incr(ctx, key)
		pipe.ExpireNX(ctx, key, window)
		_, err := pipe.Exec(ctx)

		if err != nil {
			c.Next()
			return
		}

		if count.Val() > int64(limit) {
			templates.ServeErrorPage(c, "You have sent too many requests. Please wait a moment and try again.")
			
			c.Abort()
			return
		}
		c.Next()
	}
}
