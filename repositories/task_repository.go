package repositories

import (
	"fmt"
	"todoBackedAPI/models"
)

func GetAllTasks() *[]models.Task {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		panic("AAAAAAA")
	}

	var tasks *[]models.Task

	db.Find(&tasks)

	return tasks
}
