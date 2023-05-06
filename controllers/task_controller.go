package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackedAPI/models"
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

func CreateTask(c *gin.Context) {
	var task models.Task
	fmt.Println(task.Title, task.Description)

	err := c.BindJSON(&task)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
		})
		return
	}

	err = repositories.CreateTask(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "created",
	})
}
