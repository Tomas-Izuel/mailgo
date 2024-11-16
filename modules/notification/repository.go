package notification

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"mailgo/lib"
	"mailgo/lib/db"
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

func getNotificationsByUser(userID string, ctx context.Context) ([]Notification, error) {
	var notifications []Notification
	cursor, err := dbCollection().Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var notification Notification
		if err := cursor.Decode(&notification); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
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
