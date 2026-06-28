package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupEntryMiddlewares() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		// execute next middleware / handler
		c.Next()

		status := c.Writer.Status()

		duration := time.Since(start)

		fmt.Printf(
			"[RESPONSE] %s %s | Status: %d | Duration: %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			duration,
		)
	}
}