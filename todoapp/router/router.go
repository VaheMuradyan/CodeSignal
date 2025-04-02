package router

import (
	"codesignal.com/example/gin/todoapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/api/todos/:id", controllers.GetTodoById)
	r.GET("/api/todos", controllers.GetTodos)
	return r
}