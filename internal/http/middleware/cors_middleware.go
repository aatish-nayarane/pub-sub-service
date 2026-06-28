package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCors() gin.HandlerFunc {

	return cors.New(
		cors.Config{
			AllowOrigins: []string{
				// "http://localhost:3000",
				"*",
			},

			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
			},

			AllowHeaders: []string{
				"Authorization",
				"Content-Type",
			},

			AllowCredentials: true,

			MaxAge: time.Hour,
		},
	)
}
