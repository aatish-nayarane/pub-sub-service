package handlers

import (
	"net/http"

	"pub-sub-service/internal/models"
	"pub-sub-service/internal/services"

	"github.com/gin-gonic/gin"
)

// create a handler for switch
type SwitchHandler struct {
	switchService *services.SwitchService
}

func NewSwitchHandler(switchService *services.SwitchService) *SwitchHandler {
	return &SwitchHandler{
		switchService: switchService,
	}
}

// Create godoc
//
// @Summary Create switch
// @Description Create switch configuration
// @Tags Switch
// @Accept json
// @Produce json
// @Param request body models.Switch true "Switch"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/switch [post]
func (s SwitchHandler) Create(c *gin.Context) {
	var req models.Switch
	if err :=
		c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		},
		)
		return
	}
	c.JSON(
		201,
		gin.H{
			"message": "created",
		},
	)

}

func (h *SwitchHandler) CreateSwitch(c *gin.Context) {

	var payload models.Switch
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.switchService.CreateSwitch(&payload)
	c.JSON(http.StatusOK, response)
}

// func (h *SwitchHandler) Update(c *gin.Context) {
// 	project_id := c.Param("project_id")
// 	var payload models.Switch
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response := h.switchService.Update(project_id, &payload)
// 	c.JSON(http.StatusOK, response)
// }

// func (h *SwitchHandler) Get(c *gin.Context) {
// 	project_id := c.Param("project_id")

// 	response := h.switchService.Get(project_id)
// 	c.JSON(http.StatusOK, response)
// }

// func (h *SwitchHandler) Delete(c *gin.Context) {
// 	project_id := c.Param("project_id")

// 	response := h.switchService.Delete(project_id)
// 	c.JSON(http.StatusOK, response)
// }
