package notification

import (
	"mailgo/lib"
	"mailgo/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetNotificationsByUserController(c *gin.Context) {
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	user, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	notifications, err := getNotificationsByUserService(user.ID, c)
	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	c.JSON(200, notifications)
}

func GetNotificationById(c *gin.Context) {
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	user, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	notificationID := c.Param("notificationId")
	if notificationID == "" {
		restErr := lib.NewRestError(400, "Notification ID is required")
		c.JSON(restErr.Status(), restErr)
	}
	notification, err := getNotificationByIdService(notificationID, user.ID, c)
	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	c.JSON(200, notification)
}
