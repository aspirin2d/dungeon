package router

import (
	"github.com/aspirin2ds/dungeon/controller"
	"github.com/aspirin2ds/dungeon/model"
	"github.com/gin-gonic/gin"
)

func InitApi(api *gin.RouterGroup) {
	model.Initialize()

	api.Use(controller.ErrorHandler())

	api.GET("/ping", controller.Ping)

	api.POST("/new", controller.NewUser)
	api.GET("/:uid", controller.GetUser)

	api.POST("/:uid/new", controller.NewCharacter)
	api.GET("/:uid/:cid", controller.GetCharacter)
}
