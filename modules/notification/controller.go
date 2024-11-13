package notification

import "github.com/gin-gonic/gin"

func GetAllRelatedNotificationsController(c *gin.Context) {
	relatedId := c.Param("relatedId")
	notifications, err := getAllRelatedNotificationsService(relatedId, c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, notifications)
}
