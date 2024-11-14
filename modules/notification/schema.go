package notification

import "time"

type Notification struct {
	Id           string                 `json:"id" bson:"_id"`
	TypeId       string                 `json:"typeId" bson:"typeId"`
	Recipient    string                 `json:"recipient" bson:"recipient"`
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
