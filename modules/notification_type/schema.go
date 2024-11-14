package notificationtype

type NotificationType struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	TemplateId string `json:"templateId" bson:"templateId"`
}

type CreateNotificationTypeDto struct {
	Name       string `json:"name" bson:"name"`
	TemplateId string `json:"templateId" bson:"templateId"`
}
