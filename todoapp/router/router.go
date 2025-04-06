package router

import (
	"codesignal.com/example/gin/todoapp/controllers"
	"codesignal.com/example/gin/todoapp/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	utils.RegisterValidators()
	r.GET("/api/todos/:id", controllers.GetTodo)
	r.GET("/api/todos", controllers.GetTodos)
	r.POST("/api/todos", controllers.CreateTodo2)
	r.POST("/api/todos/bulk", controllers.BulkUploadTodos)
	r.POST("/api/todos/:id/image", controllers.UploadImage)
	r.GET("/api/todos/:id/image", controllers.GetImage)
}
