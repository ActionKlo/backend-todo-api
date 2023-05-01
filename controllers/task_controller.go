package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackedAPI/repositories"
)

func SayHello(c *gin.Context) {
	fmt.Println("SayHello")
	tasks := repositories.GetAllTasks()
	fmt.Println(tasks)
	c.JSON(http.StatusOK, gin.H{
		"message": tasks,
	})
}
