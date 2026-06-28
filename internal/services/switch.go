package services


import "alert_and_notification/internal/models"



// create a service for switch
type  SwitchService struct {}

//  craete constructor for switch service
func NewSwitchService() *SwitchService {
	return &SwitchService{}
}

func (s *SwitchService) Create(payload *models.Switch) any {
	return  map[string]any{
		"message": "switch created successfully",
		"data": payload,
	}
}	


func (s *SwitchService) Update(project_id string, payload *models.Switch) any {
	return  map[string]any{
		"message": "switch updated successfully",
		"data": payload,
	}
}

func (s *SwitchService) Get(project_id string) any {
	return  map[string]any{
		"message": "switch fetched successfully",
		"data": map[string]any{
			"project_id": project_id,
			"user_id": "123",
			"email_enable": true,
			"audit_enable": false,
		},
	}
}

func (s *SwitchService) Delete(project_id string) any {
	return  map[string]any{
		"message": "switch deleted successfully",
	}
}