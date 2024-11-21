package notification

import (
	"mailgo/lib"

	"github.com/gin-gonic/gin"
)

func GetNotificationsByUserController(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		restErr := lib.NewRestError(400, "User ID is required")
		c.JSON(restErr.Status(), restErr)
	}
	notifications, err := getNotificationsByUserService(userID, c)
	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	c.JSON(200, notifications)
}

func GetNotificationById(c *gin.Context) {
	userID := c.Param("userId")
	notificationID := c.Param("notificationId")
	if notificationID == "" {
		restErr := lib.NewRestError(400, "Notification ID is required")
		c.JSON(restErr.Status(), restErr)
	}
	notification, err := getNotificationByIdService(notificationID, userID, c)
	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	c.JSON(200, notification)
}
