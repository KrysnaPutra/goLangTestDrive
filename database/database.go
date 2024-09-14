package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, errors := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if errors != nil {
		panic(errors)
	}
	return db
}
