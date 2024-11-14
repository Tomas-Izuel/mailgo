package notificationtype

import (
	"mailgo/lib"

	"github.com/gin-gonic/gin"
)

func CreateNotificationTypeController(c *gin.Context) {
	notificationTypeDto := &CreateNotificationTypeDto{}
	if err := c.ShouldBindJSON(notificationTypeDto); err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	createdNotificationType, err := createNotificationTypeService(notificationTypeDto, c)
	if err != nil {
		restErr := lib.NewRestError(500, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, createdNotificationType)
}

func GetNotificationTypeByIDController(c *gin.Context) {
	typeId := c.Param("typeId")
	if typeId == "" {
		restErr := lib.NewRestError(400, "Type ID is required")
		c.JSON(restErr.Status(), restErr)
		return
	}
	notificationType, err := getNotificationTypeByIDService(typeId, c)
	if err != nil {
		restErr := lib.NewRestError(404, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, notificationType)
}

func UpdateNotificationTypeController(c *gin.Context) {
	notificationType := &NotificationType{}
	if err := c.ShouldBindJSON(notificationType); err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	updatedNotificationType, err := updateNotificationTypeService(notificationType, c)
	if err != nil {
		restErr := lib.NewRestError(500, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, updatedNotificationType)
}

func DeleteNotificationTypeController(c *gin.Context) {
	typeId := c.Param("typeId")
	if typeId == "" {
		restErr := lib.NewRestError(400, "Type ID is required")
		c.JSON(restErr.Status(), restErr)
		return
	}
	err := deleteNotificationTypeService(typeId, c)
	if err != nil {
		restErr := lib.NewRestError(404, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, nil)
}
