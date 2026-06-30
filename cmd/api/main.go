package main

import (
	"log"
	"pub-sub-service/internal/bootstrap"
	"pub-sub-service/internal/config"
	"pub-sub-service/internal/http/middleware"
	"pub-sub-service/internal/http/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

func main() {

	// Initialize logger
	logger :=
		config.SetLogger()

	defer logger.Sync()

	// Initialize Switch dependencies
	switchHandler,
		err :=
		bootstrap.SwitchBootstrap()

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
	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
		),
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
