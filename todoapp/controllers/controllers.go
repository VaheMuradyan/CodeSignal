package controllers

import (
	"net/http"
	"strconv"

	"codesignal.com/example/gin/todoapp/services"
	"github.com/gin-gonic/gin"
)

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
