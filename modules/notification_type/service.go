package notificationtype

import "context"

func getNotificationTypeByIDService(typeId string, ctx context.Context) (*NotificationType, error) {
	notificationType, err := getNotificationTypeByID(typeId, ctx)
	if err != nil {
		return nil, err
	}

	return notificationType, nil
}

func createNotificationTypeService(notificationTypeDto *CreateNotificationTypeDto, ctx context.Context) (*NotificationType, error) {
	id, err := createNotificationType(notificationTypeDto, ctx)
	if err != nil {
		return nil, err
	}

	createdNotificationType := &NotificationType{
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
