package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackedAPI/repositories"
)

func GetAllTasks(c *gin.Context) {
	tasks := repositories.GetAllTasks()

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Params[0].Value
	task, err := repositories.GetTaskById(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"err": "task not found",
		})
	} else {
		c.JSON(http.StatusOK, task)
	}
}
