package todorepository

import (
	"codesignal.com/example/gin/todoapp/models"
	"gorm.io/gorm"
)

func FindAllTodos(db *gorm.DB) []models.Todo {
	var todo []models.Todo
	db.Find(&todo)
	return todo
}

func CreateTodo(db *gorm.DB, todo models.Todo) models.Todo {
	db.Create(&todo)
	return todo
}

func ResetTodos(db *gorm.DB) {
	db.Exec("DELETE FROM todos")
	db.Exec("ALTER TABLE todos AUTO_INCREMENT = 1")
}

// func ConnectDatabase() *gorm.DB {
//     db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database")
//     }

//     db.AutoMigrate(&models.Todo{})
//     return db
// }
