package routes

import (
	"alert_and_notification/internal/http/handlers"
	"alert_and_notification/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterV1(r *gin.Engine) {
	switchService := services.NewSwitchService()
	handlers := handlers.NewSwitchHandler(switchService)
	v1 := r.Group("/api/v1")
	switches := v1.Group("/switch")
	{
		switches.POST("", handlers.Create)
		switches.GET("/:project_id", handlers.Get)
		switches.PUT("/:project_id", handlers.Update)
		switches.DELETE("/:project_id", handlers.Delete)
	}

}
