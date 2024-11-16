package notification

import (
	"time"
)

type Notification struct {
	ID           string                 `json:"id" bson:"_id"`
	TypeId       string                 `json:"typeId" bson:"typeId"`
	UserId       string                 `json:"userId" bson:"userId"`
	Recipient    string                 `json:"recipient" bson:"recipient"`
	RelatedId    string                 `json:"relatedId" bson:"relatedId"`
	CreatedAt    time.Time              `json:"createdAt" bson:"createdAt"`
	EventDetails map[string]interface{} `json:"eventDetails" bson:"eventDetails"` // Datos adicionales según el tipo de evento
	Mail         struct {
		Subject  string `json:"subject" bson:"subject"`
		BodyHTML string `json:"bodyHtml" bson:"bodyHtml"`
		BodyText string `json:"bodyText" bson:"bodyText"`
	} `json:"mail"`
}

type CreateNotificationDto struct {
	Event        string                 `json:"type" bson:"type"`
	RelatedId    string                 `json:"relatedId" bson:"relatedId"`
	UserId       string                 `json:"userId" bson:"userId"`
	EventDetails map[string]interface{} `json:"eventDetails" bson:"eventDetails"`
}
