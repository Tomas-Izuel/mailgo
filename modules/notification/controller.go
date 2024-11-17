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
	notifications, err := getNotificationsByUser(userID, c)
	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	c.JSON(200, notifications)
}
