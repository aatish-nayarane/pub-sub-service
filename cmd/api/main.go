package main

import (
	"log"
	"alert_and_notification/internal/bootstrap"
	"alert_and_notification/internal/config"
	"alert_and_notification/internal/http/middleware"
	"alert_and_notification/internal/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize logger
	logger :=
		config.SetLogger()

	defer logger.Sync()

	// Initialize Switch dependencies
	switchHandler,
	err :=
		bootstrap.InitSwitchHandler()

	if err != nil {

		logger.Error(
			"failed to initialize switch module",
		)

		log.Fatal(
			err,
		)
	}

	// Create Gin Engine
	r :=
		gin.New()

	// Security
	r.SetTrustedProxies(nil)

	// Recover panic
	r.Use(
		gin.Recovery(),
	)

	// Inject logger
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
	routes.RegisterV1(
		r,
		switchHandler,
	)

	logger.Info(
		"server started",
	)

	if err :=
		r.Run(
			":8080",
		); err != nil {

		logger.Error(
			"server startup failed",
		)

		log.Fatal(
			err,
		)
	}
}