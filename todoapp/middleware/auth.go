package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const apiKey = "secret-key"

		providedKey := c.GetHeader("X-API-KEY")

		if providedKey == "" || providedKey != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid or missing API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
