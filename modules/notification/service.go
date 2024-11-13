package notification

import "context"

func getAllRelatedNotificationsService(relatedId string, ctx context.Context) ([]Notification, error) {
	related, err := findAllNotificationsFromRelatedId(relatedId, ctx)

	if err != nil {
		return nil, err
	}

	return related, nil
}
