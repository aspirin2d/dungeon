package main

import (
	_ "github.com/aspirin2ds/dungeon/docs"
	"github.com/aspirin2ds/dungeon/router"

	"github.com/gin-gonic/gin"

	// swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

//	@title			Dungeon API
//	@version		1.0
//	@description	This is an ai-driven dungeon server.

// @host		localhost:8080
// @BasePath	/api/v1
func main() {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	router.InitApi(v1)

	// swagger embed files
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
