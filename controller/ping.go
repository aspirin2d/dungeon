package controller

import "github.com/gin-gonic/gin"

// Ping godoc
// @Summary		ping to the server
// @Tags			test
// @Produce		plain
// @Success		200
// @Router			/ping [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
}
