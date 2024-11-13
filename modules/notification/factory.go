// factory.go
package notification

import "errors"

func NotificationFactory(notificationType string, baseNotification Notification) (NotificationHandler, error) {
	switch notificationType {
	case "feedback":
		return &FeedbackNotification{baseNotification}, nil
	case "cart":
		return &CartNotification{baseNotification}, nil
	case "order":
		return &OrderNotification{baseNotification}, nil
	default:
		return nil, errors.New("unsupported notification type")
	}
}
