package models
type Switch struct {
	Project_id string `json:"project_id" bson:"project_id"`
	User_id    string `json:"user_id" bson:"user_id"`
	Email_enable bool `json:"email_enable" bson:"email_enable"`
	Audit_enable bool `json:"audit_enable" bson:"audit_enable"`
}

