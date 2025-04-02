package services

import "codesignal.com/example/gin/todoapp/models"

var todos = []models.Todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Master Gin", Completed: true},
}

func FindTodoById(id int) (models.Todo, bool){
	for _, todo := range todos{
		if todo.ID == id {
			return todo, true
		}
	}
	return models.Todo{}, false
}