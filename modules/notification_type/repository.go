package notificationtype

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"mailgo/lib"
	"mailgo/lib/db"
	"mailgo/lib/log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrTypeID = lib.NewValidationError().Add("typeId", "Invalid")

var collection *mongo.Collection

func dbCollection() *mongo.Collection {
	if collection == nil {
		database := db.Get()
		collection = database.Collection("notification_type")
	}
	return collection
}

func findNotificationTypes(ctx ...interface{}) ([]NotificationType, error) {
	var notificationTypes []NotificationType

	cursor, err := dbCollection().Find(context.TODO(), bson.D{})
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	if err = cursor.All(context.TODO(), &notificationTypes); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	return notificationTypes, nil
}

func createNotificationType(notificationTypeDto *CreateNotificationTypeDto, ctx context.Context) (string, error) {
	result, err := dbCollection().InsertOne(ctx, notificationTypeDto)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", ErrTypeID
}

func updateNotificationType(notificationType *NotificationType, ctx context.Context) error {
	_, err := dbCollection().UpdateOne(ctx, bson.M{"typeId": notificationType.ID}, bson.M{"$set": notificationType})
	if err != nil {
		return err
	}

	return nil
}

func deleteNotificationType(typeId string, ctx context.Context) error {
	_, err := dbCollection().DeleteOne(ctx, bson.M{"typeId": typeId})
	if err != nil {
		return err
	}

	return nil
}

func findNotificationTypeByEventKey(eventKey string, ctx context.Context) (*NotificationType, error) {
	var notificationType NotificationType

	// Por el momento solo trae el último registro
	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})

	if err := dbCollection().FindOne(context.TODO(), bson.M{"eventKeys": eventKey}, opts).Decode(&notificationType); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Get(ctx).Error("No notification type found for event key: ", eventKey)
			return nil, ErrTypeID
		}
		log.Get(ctx).Error("Database error: ", err)
		return nil, err
	}

	return &notificationType, nil
}
