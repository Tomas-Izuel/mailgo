package notificationtype

import (
	"mailgo/lib"
	"mailgo/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateNotificationTypeController(c *gin.Context) {
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	_, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

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

func GetNotificationTypesController(c *gin.Context) {
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	_, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	notificationTypes, err := getNotificationTypesService(c)
	if err != nil {
		restErr := lib.NewRestError(500, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, notificationTypes)
}

func UpdateNotificationTypeController(c *gin.Context) {
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	_, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

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
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	_, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	typeId := c.Param("typeId")
	if typeId == "" {
		restErr := lib.NewRestError(400, "Type ID is required")
		c.JSON(restErr.Status(), restErr)
		return
	}
	err = deleteNotificationTypeService(typeId, c)
	if err != nil {
		restErr := lib.NewRestError(404, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, nil)
}
