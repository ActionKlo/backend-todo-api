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

	router.GET("/drop-tasks", func(c *gin.Context) {
		db, err := repositories.ConnectDB()

		if err != nil {
			panic(err)
		}

		db.Migrator().DropTable("tasks")
	})

	return router
}
