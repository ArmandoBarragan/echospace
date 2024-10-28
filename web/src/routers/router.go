package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Create a router that contains the application's endpoints
	var router *gin.Engine = gin.Default()
	router.GET("api/home", homePage)
	router.GET("api/login", login)
	router.POST("api/create_user", createUser)
	return router
}
