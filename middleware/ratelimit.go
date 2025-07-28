package middleware

import (
	"context"
	"net/http"
	"relay/templates"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	luaScript := `
		local key = KEYS[1]
		local now = tonumber(ARGV[1])
		local window = tonumber(ARGV[2])
		local limit = tonumber(ARGV[3])
		
		redis.call("ZREMRANGEBYSCORE", key, 0, now - window)
		
		local count = redis.call("ZCARD", key)
		
		if count < limit then
			redis.call("ZADD", key, now, now .. ":" .. math.random())
		
			if count == 0 then
				redis.call("EXPIRE", key, window)
			end
		
			return 1
		end
		return 0
	`

	script := redis.NewScript(luaScript)
	
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()
		key := "rl:" + ip

		now := time.Now().Unix()
		
		result, err := script.Run(ctx, redisClient, []string{key}, now, window.Seconds(), limit).Int()

		if err != nil {
			c.Next()
			return
		}

		if result == 0 {
			c.Status(http.StatusTooManyRequests)
			templates.ServeErrorPage(c, "You have sent too many requests. Please wait a moment and try again.")
			c.Abort()
			return
		}
		c.Next()
	}
}
