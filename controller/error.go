package controller

import (
	"fmt"

	"github.com/aspirin2ds/dungeon/model"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// only check for the first error if there are any
		if len(c.Errors) > 0 {
			last := c.Errors.Last()
			if last.IsType(gin.ErrorTypePrivate) {
				c.JSON(500, gin.H{"msg": "internal server error"})
			} else if last.IsType(gin.ErrorTypeBind) {
				c.JSON(c.Writer.Status(), gin.H{"msg": fmt.Sprintf("binding error: %s", last.Error())})
			} else if last.IsType(gin.ErrorTypeRender) {
				c.JSON(c.Writer.Status(), gin.H{"msg": fmt.Sprintf("rendering error: %s", last.Error())})
			} else {
				c.JSON(c.Writer.Status(), gin.H{"msg": last.Error()})
			}
			model.LogError(c.Errors.JSON())
		}
	}
}

type ErrorMessage struct {
	Message string `json:"msg" example:"error message"`
}
