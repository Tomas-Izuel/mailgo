package template

import (
	"mailgo/lib"
	"mailgo/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateTemplateController(c *gin.Context) {
	userRec := c.Request.Header.Get("Authorization")
	userRec = strings.TrimPrefix(userRec, "Bearer ")
	_, err := security.Validate(userRec)

	if err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
	}

	templateDto := &CreateMailTemplateDto{}
	if err := c.ShouldBindJSON(templateDto); err != nil {
		restErr := lib.NewRestError(400, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	createdTemplate, err := createTemplateService(templateDto, c)
	if err != nil {
		restErr := lib.NewRestError(500, err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(200, createdTemplate)
}
