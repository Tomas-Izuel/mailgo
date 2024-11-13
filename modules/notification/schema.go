package notification

import "time"

type Notification struct {
	Id           string                 `json:"id" bson:"_id"`
	TemplateId   string                 `json:"templateId" bson:"templateId"`
	Recipient    string                 `json:"recipient" bson:"recipient"`
	Type         string                 `json:"type" bson:"type"` //"feedback", "cart", "order"
	RelatedId    string                 `json:"relatedId" bson:"relatedId"`
	CreatedAt    time.Time              `json:"createdAt" bson:"createdAt"`
	EventDetails map[string]interface{} `json:"eventDetails" bson:"eventDetails"` // Datos adicionales según el tipo de evento
}

type CreateNotificationDto struct {
	Type         string                 `json:"type" bson:"type"`
	RelatedId    string                 `json:"relatedId" bson:"relatedId"`
	EventDetails map[string]interface{} `json:"eventDetails" bson:"eventDetails"`
}

type NotificationHandler interface {
	Process() error
}
