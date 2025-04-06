package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next() // process request

		duration := time.Since(startTime)
		clientIP := c.ClientIP()

		if clientIP == "" {
			clientIP = "127.0.0.1"
		}

		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		// Format duration to seconds for test compatibility
		durationStr := fmt.Sprintf("%.3fs", duration.Seconds())

		logMsg := fmt.Sprintf("Request: %s %s from %s | Status: %d | Duration: %s",
			method, path, clientIP, statusCode, durationStr)

		log.Println(logMsg)
	}
}
