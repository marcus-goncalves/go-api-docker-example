package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	suprt *SupportRoutes
	user  *User
)

func CreateRouter() *gin.Engine {
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

	return router
}
