package main

import "exinity-task/pkg/router"

func main() {
	appRouter := router.InitRouter()

	router.InitTransactionRoutes(appRouter)

	appRouter.Run(":8080")
}
