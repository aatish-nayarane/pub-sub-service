package main

import (
	"alert_and_notification/internal/config"
	"alert_and_notification/internal/http/middleware"
	"alert_and_notification/internal/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize logger
	logger := config.SetLogger()

	// Create Gin Engine
	r := gin.Default()

	// Security
	r.SetTrustedProxies(nil)

	// Recovery middleware
	r.Use(gin.Recovery())

	// Inject logger into request context
	r.Use(
		middleware.InjectLogger(
			logger,
		),
	)

	// Custom middlewares
	r.Use(
		middleware.SetupEntryMiddlewares(),
	)

	r.Use(
		middleware.SetupCors(),
	)

	// Register routes
	routes.RegisterV1(r)

	// Start server
	logger.Info(
		"server started",
	)

	r.Run(":8080")
}
