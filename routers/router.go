package routers

import (
	"github.com/gin-gonic/gin"
	"todoBackedAPI/controllers"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/task/:id", controllers.GetTaskById)

	return router
}
