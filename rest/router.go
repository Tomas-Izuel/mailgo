package rest

import (
	"mailgo/modules/notification"
	notificationtype "mailgo/modules/notification_type"
	"mailgo/modules/template"
)

func initRouter() {
	if server == nil {
		panic("Server non existent")
	}

	//Notifications
	notificationGroup := server.Group("/notification")
	notificationGroup.GET("/", notification.GetNotificationsByUserController)

	//Notification Types
	notificationTypeGroup := server.Group("/notification-type")
	notificationTypeGroup.POST("/", notificationtype.CreateNotificationTypeController)
	notificationTypeGroup.GET("/:typeId", notificationtype.GetNotificationTypeByIDController)
	notificationTypeGroup.PUT("/:typeId", notificationtype.UpdateNotificationTypeController)
	notificationTypeGroup.DELETE("/:typeId", notificationtype.DeleteNotificationTypeController)

	//Email Templates
	templateGroup := server.Group("/template")
	templateGroup.POST("/", template.CreateTemplateController)
}
