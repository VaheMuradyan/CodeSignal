package main

import "codesignal.com/example/gin/router"

func main(){
	r := router.SetupRouter()

	r.Run(":8080")
}