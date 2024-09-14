package model

import (
	"belajargo/database"
	"fmt"

	"gorm.io/gorm"
)

func EmDb() *gorm.DB {
	return database.ConnectDatabase()
}

func MigrateAll() {

	db := database.ConnectDatabase()
	db.AutoMigrate(
		&Product{},
		&User{},
	)

	fmt.Println("Migration Success!")
}
