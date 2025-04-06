package services

import (
	"errors"
	"path/filepath"

	"codesignal.com/example/gin/todoapp/models"
	"codesignal.com/example/gin/todoapp/utils"
	"github.com/gin-gonic/gin"
)

var (
	todos = []models.Todo{
		{ID: 1, Title: "Learn Go", Completed: false},
		{ID: 2, Title: "Master Gin", Completed: true},
	}
	idCounter int
)

func AddTodo(c *gin.Context) (models.Todo, map[string]string) {
	var newTodo models.Todo
	var validationErrors map[string]string

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		validationErrors = utils.EnhancedErrorMessages(err)
	} else {
		idCounter++
		newTodo.ID = idCounter
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

func AddTodos(todos *[]models.Todo, newTodos []models.Todo) []models.Todo {
	startId := len(*todos)

	for i, _ := range newTodos {
		newTodos[i].ID = startId + 1 + i
		*todos = append(*todos, newTodos[i])
	}
	return newTodos
}

func FindTodoByID2(id int) (models.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return models.Todo{}, errors.New("Todo not found")
}

func AddTodoService(newTodo models.Todo) (models.Todo, error) {
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	return newTodo, nil
}

func UploadTodoImage(c *gin.Context, id int) error {
	todo, err := findTodoByID(id)
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return errors.New("No file uploaded")
	}

	fileName := filepath.Base(file.Filename)
	filePath := filepath.Join("uploads", fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return errors.New("Failed to save file")
	}

	todo.ImagePath = filePath
	return nil
}

func GetTodoImage(c *gin.Context, id int) error {
	todo, err := findTodoByID(id)
	if err != nil {
		return err
	}

	if todo.ImagePath != "" {
		c.File(todo.ImagePath)
		return nil
	}

	return errors.New("Image not found")
}

func getTodos() []models.Todo {
	return todos
}

func findTodoByID(id int) (*models.Todo, error) {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}
