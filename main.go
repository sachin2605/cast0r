package main

import (
	"flag"
	"fmt"

	"github.com/sachin2605/cast0r/configs"
	"github.com/sachin2605/cast0r/controller"
	"github.com/sachin2605/cast0r/svc"

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

	api := router.Group("api/v1/fruits")
	{
		fruitCtrl := new(controller.FruitController)
		api.GET("/", fruitCtrl.List)
		api.GET("/:id", fruitCtrl.Get)
		api.POST("/", fruitCtrl.Create)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
func main() {
	var envFlag string
	flag.StringVar(&envFlag, "env", "dev", "Environment flag: 'dev' or 'test'")
	flag.Parse()
	fmt.Println("Environment: ", envFlag)

	// Load environment variables based on the flag
	configs.LoadEnv(envFlag)
	svc.InitServices()
	router := SetupRouter()
	router.Run()
}
