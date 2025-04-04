package router

import (
	"codesignal.com/example/gin/todoapp/controllers"
	"codesignal.com/example/gin/todoapp/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	utils.RegisterValidators()
	r.GET("/api/todos/:id", controllers.GetTodoById)
	r.GET("/api/todos", controllers.GetTodos)
	r.POST("/api/todos", controllers.CreateTodo)
	return r
}
