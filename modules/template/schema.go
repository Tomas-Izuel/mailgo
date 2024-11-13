package template

type MailTemplate struct {
	Id       string `json:"id" bson:"_id"`
	Subject  string `json:"subject" bson:"subject"`
	BodyHTML string `json:"bodyHtml" bson:"bodyHtml"` // HTML del cuerpo del correo
	BodyText string `json:"bodyText" bson:"bodyText"` // Texto plano del correo
}
