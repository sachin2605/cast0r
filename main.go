package main

import (
	"github.com/sachin2605/cast0r/controller"
	"github.com/sachin2605/cast0r/svc"

	"github.com/joho/godotenv"
	_ "github.com/sachin2605/cast0r/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Castor Fruits  API
// @version 1.0
// @license.name Apache 2.0
// @description Fruits API Server.
// @host localhost:8080
// @contact.name Sachin Kumar
// @contact.name sachin26051@gmail.com
// @BasePath /api/v1

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("api/v1")
	{
		fruitCtrl := new(controller.FruitController)
		api.GET("/fruits", fruitCtrl.List)
		api.GET("/fruits/:id", fruitCtrl.Get)
		api.POST("/fruits", fruitCtrl.Create)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
func main() {
	godotenv.Load()
	svc.InitServices()
	router := SetupRouter()
	router.Run()
}
