package routes

import (
	_ "go-docker-example/v1/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	suprt *SupportRoutes
	user  *User
)

func CreateRouter() *gin.Engine {
	// @title GO Docker Example
	// @description This is an example of Gin and Swagger usage
	// @host localhost:3000
	router := gin.Default()

	// Support Routes
	sup := router.Group("/support")
	sup.GET("/ping", suprt.Ping)

	// User Routes
	us := router.Group("/user")
	us.GET("/list", user.List)
	us.GET("/find", user.Find)
	us.POST("/add", user.Add)
	us.DELETE("/del/:id", user.Delete)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
