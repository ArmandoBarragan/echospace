package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Create a router that contains the application's endpoints
	var router *gin.Engine = gin.Default()
	router.GET("/home", homePage)
	router.GET("/login", login)
	router.POST("/create_user", createUser)
	return router
}
