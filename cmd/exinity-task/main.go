package main

import "exinity-task/pkg/router"

func main() {
	router := router.InitRouter()

	router.Run(":8080")
}
