package middlewares

import (
	//"log"
	//"net/http"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)


func SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
	start := time.Now()
	header := c.Request.Header.Get("X-PING") 
	if header == "ping" {
		log.Println("бакаки")
		c.Writer.Header().Set("X-PONG", "pong")
	}
	time.Since(start)
	}
}