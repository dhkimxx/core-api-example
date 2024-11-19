package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received: ", c.Request.URL.Path)
		c.Next()
		fmt.Println("Response sent: ", c.Writer.Status())
	}
}
