package notificationtype

import "go.mongodb.org/mongo-driver/v2/bson"

type NotificationType struct {
	ID         bson.ObjectID `bson:"_id" json:"id"`
	Name       string        `json:"name" bson:"name"`
	TemplateId string        `json:"templateId" bson:"templateId"`
	EventKeys  []string      `json:"eventKeys" bson:"eventKeys"`
}

type ResponseNotificationTypeDto struct {
	ID         string   `json:"id" bson:"_id"`
	Name       string   `json:"name" bson:"name"`
	TemplateId string   `json:"templateId" bson:"templateId"`
	EventKeys  []string `json:"eventKeys" bson:"eventKeys"`
}

type CreateNotificationTypeDto struct {
	Name       string   `json:"name" bson:"name"`
	TemplateId string   `json:"templateId" bson:"templateId"`
	EventKeys  []string `json:"eventKeys" bson:"eventKeys"`
}
