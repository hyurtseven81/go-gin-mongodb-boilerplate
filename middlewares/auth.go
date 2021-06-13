package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// config := config.GetConfig()
		authorizationHeaderValue := c.Request.Header.Get("Authorization")

		if authorizationHeaderValue == "" {
			c.AbortWithStatus(401)
		}

		c.Next()
	}
}
