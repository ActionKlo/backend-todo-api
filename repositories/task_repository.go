package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"todoBackedAPI/models"
)

func Connect() *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	return db
}

func GetAllTasks() *[]models.Task {
	db := Connect()
	var tasks *[]models.Task

	db.Find(&tasks)

	return tasks
}

func GetTaskById(id string) (*models.Task, error) {
	db := Connect()
	var task *models.Task

	res := db.First(&task, id)

	fmt.Println(res.RowsAffected)

	if res.RowsAffected == 0 {
		return nil, errors.New("task not found") // add custom errors
	}

	return task, nil
}

func CreateTask(task models.Task) error {
	db := Connect()

	db.Create(&task)

	return nil
}
