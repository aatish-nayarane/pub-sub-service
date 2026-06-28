package models

type Alert struct {
	Project_id string `json:"project_id"`
	User_id    string `json:"user_id"`
	Email_enable bool `json:"email_enable"`
	Audit_enable bool `json:"audit_enable"`
}