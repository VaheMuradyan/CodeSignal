package db

import (
	"codesignal.com/example/gin/todoapp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:java@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faild to connect database")
	}

	db.AutoMigrate(&models.Todo{})
	return db
}
