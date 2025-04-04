package services

import (
	"codesignal.com/example/gin/todoapp/models"
	"codesignal.com/example/gin/todoapp/utils"
	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Master Gin", Completed: true},
}

func AddTodo(c *gin.Context) (models.Todo, map[string]string) {
	var newTodo models.Todo
	var validationErrors map[string]string

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		validationErrors = utils.EnhancedErrorMessages(err)
	} else {
		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)
	}

	return newTodo, validationErrors
}

func FetchTodos() []models.Todo {
	return todos
}

func FindTodoById(id int) (models.Todo, bool) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, true
		}
	}
	return models.Todo{}, false
}

func FilterTodos(completed *bool) []models.Todo {
	var filtredTodos []models.Todo

	if completed == nil {
		return todos
	}

	for _, td := range todos {
		if td.Completed == *completed {
			filtredTodos = append(filtredTodos, td)
		}
	}

	return filtredTodos
}
