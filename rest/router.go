package rest

import (
	"mailgo/modules/notification"
)

func initRouter() {
	if server == nil {
		panic("Server non existent")
	}

	notificationGroup := server.Group("/notifications")
	notificationGroup.GET("/:relatedId", notification.GetAllRelatedNotificationsController)
}
