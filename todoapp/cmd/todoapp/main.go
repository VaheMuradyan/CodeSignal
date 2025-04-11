package main

import (
	"codesignal.com/example/gin/todoapp/middleware"
	"codesignal.com/example/gin/todoapp/repositories/db"
	"codesignal.com/example/gin/todoapp/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDatabase()
	r := gin.Default()

	r.Use(middleware.RequestLoggerMiddleware())
	router.SetupRouter(r, database)

	r.Run(":8080")
}
