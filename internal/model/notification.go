package model

import "encoding/json"

// How it needs to be stored in db
type CreateNotificationRequest struct{
	RecipientEmail 	string					`json:"recipient_email"`
	TemplateID 		int64					`json:"template_id"`
	Payload 		map[string]interface{}	`json:"payload"`
}

// Response
type Notification struct {
	ID 				int64
	RecipientEmail 	string
	TemplateID 		int64 
	PayloadJson 	json.RawMessage
	Status 			string
	ErrorMessage  	*string
}

