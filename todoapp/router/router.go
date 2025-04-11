package router

import (
	"codesignal.com/example/gin/todoapp/controllers"
	"codesignal.com/example/gin/todoapp/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {
	utils.RegisterValidators()
	// r.GET("/api/todos/:id", controllers.GetTodo)
	// r.GET("/api/todos", controllers.GetTodos)
	// r.POST("/api/todos", controllers.CreateTodo2)
	// r.POST("/api/todos/bulk", controllers.BulkUploadTodos)
	// r.POST("/api/todos/:id/image", controllers.UploadImage)
	// r.GET("/api/todos/:id/image", controllers.GetImage)
	router.GET("/api/todos", func(c *gin.Context) {
		controllers.GetTodosHandler(c, db)
	})

	router.POST("/api/todos", func(c *gin.Context) {
		controllers.CreateTodoHandler(c, db)
	})

	router.DELETE("/api/reset", func(c *gin.Context) {
		controllers.ResetTodosHandler(c, db)
	})
}
