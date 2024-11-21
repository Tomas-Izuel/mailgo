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

type EventNotificationDto struct {
	EventKey     string                 `json:"EventKey" bson:"EventKey"`
	RelatedId    string                 `json:"RelatedId" bson:"RelatedId"`
	UserId       string                 `json:"UserId" bson:"UserId"`
	EventDetails map[string]interface{} `json:"EventDetails" bson:"EventDetails"`
}
