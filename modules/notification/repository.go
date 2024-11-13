package notification

import (
	"context"
	"mailgo/lib"
	"mailgo/lib/db"
	"mailgo/lib/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrID = lib.NewValidationError().Add("id", "Invalid")

var collection *mongo.Collection

func dbCollection() *mongo.Collection {

	if collection == nil {
		database := db.Get()
		collection = database.Collection("notification")
	}

	return collection
}

func findAllNotificationsFromRelatedId(relatedId string, ctx ...interface{}) ([]Notification, error) {
	var notifications []Notification
	_id, err := primitive.ObjectIDFromHex(relatedId)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, ErrID
	}

	cursor, err := dbCollection().Find(context.TODO(), bson.M{"_id": _id})
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &notifications); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	return notifications, nil
}

func createNotification(notification *Notification, ctx context.Context) (string, error) {
	result, err := dbCollection().InsertOne(ctx, notification)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", ErrID
}
