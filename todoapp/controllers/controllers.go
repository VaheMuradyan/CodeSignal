package controllers

import (
	"net/http"
	"strconv"

	"codesignal.com/example/gin/todoapp/services"
	"github.com/gin-gonic/gin"
)

func GetTodoById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

	todo, found := services.FindTodoById(id)
	if found {
		c.JSON(http.StatusOK, todo)
	}else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}
}