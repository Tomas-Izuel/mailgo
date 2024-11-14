package template

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func findTemplateByID(templateId string, ctx ...interface{}) (*MailTemplate, error) {
	var mailTemplate MailTemplate

	if err := dbCollection().FindOne(context.TODO(), bson.M{"templateId": templateId}).Decode(&mailTemplate); err != nil {
		log.Get(ctx...).Error(err)
		return nil, ErrTemplateID
	}

	return &mailTemplate, nil
}

func createTemplate(mailTemplate *CreateMailTemplateDto, ctx context.Context) (string, error) {
	result, err := dbCollection().InsertOne(ctx, mailTemplate)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), ErrTemplateID
}
