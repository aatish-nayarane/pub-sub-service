package routes

import (
	"alert_and_notification/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterV1(
	r *gin.Engine,
	switchHandler *handlers.SwitchHandler,
) {

	v1 :=
		r.Group(
			"/api/v1",
		)

	v1.POST(
		"/switch",
		switchHandler.Create,
	)

	// v1.GET(
	// 	"/switch/:project_id",
	// 	switchHandler.Get,
	// )

	// v1.PUT(
	// 	"/switch/:project_id",
	// 	switchHandler.Update,
	// )

	// v1.DELETE(
	// 	"/switch/:project_id",
	// 	switchHandler.Delete,
	// )
}