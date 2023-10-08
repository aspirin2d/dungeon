package main

import (
	"github.com/aspirin2ds/dungeon/controllers"
	_ "github.com/aspirin2ds/dungeon/docs"

	"github.com/gin-gonic/gin"

	// swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

//	@title			Dungeon API
//	@version		1.0
//	@description	This is an ai-driven dungeon server.

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	aspirin2ds@outlook.com

//	@license.name	GNU 3.0
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.en.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/ping", controllers.Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
