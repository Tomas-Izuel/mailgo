package notificationtype

import (
	"context"
	"fmt"
)

func getNotificationTypesService(ctx context.Context) ([]NotificationType,
	error) {
	notificationTypes, err := findNotificationTypes(ctx)
	if err != nil {
		return nil, err
	}

	return notificationTypes, nil
}

func createNotificationTypeService(
	notificationTypeDto *CreateNotificationTypeDto,
	ctx context.Context) (*ResponseNotificationTypeDto, error) {
	id, err := createNotificationType(notificationTypeDto, ctx)
	fmt.Print(id)
	if err != nil {
		return nil, err
	}

	createdNotificationType := &ResponseNotificationTypeDto{
		ID:         id,
		Name:       notificationTypeDto.Name,
		TemplateId: notificationTypeDto.TemplateId,
		EventKeys:  notificationTypeDto.EventKeys,
	}

	return createdNotificationType, nil
}

func updateNotificationTypeService(notificationType *NotificationType, ctx context.Context) (*NotificationType, error) {
	err := updateNotificationType(notificationType, ctx)
	if err != nil {
		return nil, err
	}

	return notificationType, nil
}

func deleteNotificationTypeService(typeId string, ctx context.Context) error {
	err := deleteNotificationType(typeId, ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetNotificationTypeByEventKeyService(eventKey string,
	ctx context.Context) (*NotificationType, error) {
	notificationType, err := findNotificationTypeByEventKey(eventKey, ctx)
	if err != nil {
		return nil, err
	}

	return notificationType, nil
}
