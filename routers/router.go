package routers

import (
	"github.com/gin-gonic/gin"
	"todoBackedAPI/controllers"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.SayHello)

	return router
}
