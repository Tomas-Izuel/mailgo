package template

import (
	"mailgo/lib"

	"github.com/gin-gonic/gin"
)

func CreateTemplateController(c *gin.Context) {
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
