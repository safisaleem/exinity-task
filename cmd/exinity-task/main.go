package main

import (
	"exinity-task/pkg/config"
	"exinity-task/pkg/model"
	"exinity-task/pkg/router"
)

func main() {
	db, err := config.InitPostgresDatabase()

	if err != nil {
		panic(err)
	}

	db.Table("transactions").AutoMigrate(&model.Transactions{})

	appRouter := router.InitRouter()

	router.InitTransactionRoutes(appRouter)

	appRouter.Run(":8080")
}
