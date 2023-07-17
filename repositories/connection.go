package repositories

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todoBackedAPI/models"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Warsaw", // TODO вынести в .env файл
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.AutoMigrate(models.Task{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}
