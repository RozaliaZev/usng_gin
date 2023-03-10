package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)


func SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
	start := time.Now()
	header := c.Request.Header.Get("X-PING") 
	if header == "ping" {
		c.Writer.Header().Set("X-PONG", "pong")
	}
	time.Since(start)
	}
}