package template

import (
	"context"
	"mailgo/lib"
	"mailgo/lib/db"
	"mailgo/lib/log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrTemplateID = lib.NewValidationError().Add("templateId", "Invalid")

var collection *mongo.Collection

func dbCollection() *mongo.Collection {
	if collection == nil {
		database := db.Get()
		collection = database.Collection("mail_templates")
	}
	return collection
}

// FindTemplateByID busca un template por su ID
func FindTemplateByID(templateId string, ctx ...interface{}) (*MailTemplate, error) {
	var mailTemplate MailTemplate

	if err := dbCollection().FindOne(context.TODO(), bson.M{"templateId": templateId}).Decode(&mailTemplate); err != nil {
		log.Get(ctx...).Error(err)
		return nil, ErrTemplateID
	}

	return &mailTemplate, nil
}

// CreateTemplate crea un nuevo template en la base de datos
func CreateTemplate(mailTemplate *MailTemplate, ctx context.Context) (string, error) {
	result, err := dbCollection().InsertOne(ctx, mailTemplate)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", ErrTemplateID
}
