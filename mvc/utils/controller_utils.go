package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}

	c.JSON(status, body)
}


func RespondAppError(c *gin.Context, appErr *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(appErr.Code, appErr)
		return
	}

	c.JSON(appErr.Code, appErr)
}
