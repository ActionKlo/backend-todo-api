package routers

import (
	"github.com/gin-gonic/gin"
	"todoBackedAPI/controllers"
	"todoBackedAPI/repositories"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/task/:id", controllers.GetTaskById)
	router.POST("/task", controllers.CreateTask)
	//router.PUT("/task/:id", controllers)
	router.DELETE("/task/:id", controllers.DeleteTaskById)

	router.GET("/drop-tasks", func(c *gin.Context) {
		db := repositories.ConnectDB()

		db.Migrator().DropTable("tasks")
	})

	return router
}
