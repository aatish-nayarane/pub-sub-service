package services

import (
	"alert_and_notification/internal/models"
	db "alert_and_notification/internal/storage/mongo"
	"context"
)

// create a service for switch
type SwitchService struct {
	repo *db.SwitchRepo
}

// craete constructor for switch service
func NewSwitchService(r *db.SwitchRepo) *SwitchService {
	return &SwitchService{
		repo: r,
	}
}

func (s *SwitchService) CreateSwitch(payload *models.Switch) error {
	return s.repo.Create(context.Background(), payload)

}

// func (s *SwitchService) Update(project_id string, payload *models.Switch) any {
// 	return map[string]any{
// 		"message": "switch updated successfully",
// 		"data":    payload,
// 	}
// }

// func (s *SwitchService) Get(project_id string) any {
// 	return map[string]any{
// 		"message": "switch fetched successfully",
// 		"data": map[string]any{
// 			"project_id":   project_id,
// 			"user_id":      "123",
// 			"email_enable": true,
// 			"audit_enable": false,
// 		},
// 	}
// }

// func (s *SwitchService) Delete(project_id string) any {
// 	return map[string]any{
// 		"message": "switch deleted successfully",
// 	}
// }
