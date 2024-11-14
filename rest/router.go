package rest

import (
	notificationtype "mailgo/modules/notification_type"
)

func initRouter() {
	if server == nil {
		panic("Server non existent")
	}

	//Notification Types
	notificationTypeGroup := server.Group("/notification-type")
	notificationTypeGroup.POST("/", notificationtype.CreateNotificationTypeController)
	notificationTypeGroup.GET("/:typeId", notificationtype.GetNotificationTypeByIDController)
	notificationTypeGroup.PUT("/:typeId", notificationtype.UpdateNotificationTypeController)
	notificationTypeGroup.DELETE("/:typeId", notificationtype.DeleteNotificationTypeController)

	//Email Templates
}
