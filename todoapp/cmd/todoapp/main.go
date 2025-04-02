package main

import "codesignal.com/example/gin/todoapp/router"

func main(){
	r := router.SetupRouter()
	r.Run()
}