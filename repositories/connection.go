package repositories

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todoBackedAPI/models"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Warsaw",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	db.AutoMigrate(models.Task{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
