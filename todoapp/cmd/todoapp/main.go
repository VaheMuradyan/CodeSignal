package main

import (
	"codesignal.com/example/gin/todoapp/middleware"
	"codesignal.com/example/gin/todoapp/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.RequestLoggerMiddleware())
	router.SetupRouter(r)

	r.Run()
}
