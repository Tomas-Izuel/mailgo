package template

import "go.mongodb.org/mongo-driver/v2/bson"

type MailTemplate struct {
	ID       bson.ObjectID `bson:"_id" json:"id"`
	Subject  string        `json:"subject" bson:"subject"`
	BodyHTML string        `json:"bodyHtml" bson:"bodyHtml"` // HTML del cuerpo del correo
}

type ResponseMailTemplateDto struct {
	ID       string `json:"id" bson:"_id"`
	Subject  string `json:"subject" bson:"subject"`
	BodyHTML string `json:"bodyHtml" bson:"bodyHtml"`
}

type MailNotificationTemplate struct {
	Subject  string `json:"subject" bson:"subject"`
	BodyHTML string `json:"bodyHtml" bson:"bodyHtml"`
}

type CreateMailTemplateDto struct {
	Subject  string `json:"subject" bson:"subject"`
	BodyHTML string `json:"bodyHtml" bson:"bodyHtml"`
}
