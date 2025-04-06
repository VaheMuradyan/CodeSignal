package controllers

import (
	"net/http"
	"strconv"

	"codesignal.com/example/gin/todoapp/models"
	"codesignal.com/example/gin/todoapp/services"
	"codesignal.com/example/gin/todoapp/utils"
	"github.com/gin-gonic/gin"
)

var todos = make([]models.Todo, 0)

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, found := services.FindTodoById(id)
	if found {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}
}

func GetTodos(c *gin.Context) {
	var completed *bool
	completedParam := c.Query("completed")
	if completedParam != "" {
		parsedCompleted, err := strconv.ParseBool(completedParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter value"})
			return
		}
		completed = &parsedCompleted
	}

	c.JSON(http.StatusOK, services.FilterTodos(completed))
}

func CreateTodo(c *gin.Context) {
	newTodo, err := services.AddTodo(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func BulkUploadTodos(c *gin.Context) {
	var newTodos []models.Todo

	if err := c.ShouldBindJSON(&newTodos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json payload"})
		return
	}

	dublicates := utils.CheckForDuplicates(newTodos)
	if len(dublicates) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dublicate todos found", "dublicates": dublicates})
		return
	}

	addedTodos := services.AddTodos(&todos, newTodos)
	c.JSON(http.StatusCreated, addedTodos)

}

func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	todo, err := services.FindTodoByID2(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo2(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	todo, err := services.AddTodoService(newTodo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, todo)
}
