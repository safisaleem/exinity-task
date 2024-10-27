package main

import (
	"exinity-task/pkg/config"
	"exinity-task/pkg/controller"
	"exinity-task/pkg/model"
	"exinity-task/pkg/repository"
	"exinity-task/pkg/router"
	"exinity-task/pkg/service"
)

func main() {
	db, err := config.InitPostgresDatabase()

	if err != nil {
		panic(err)
	}

	db.Table("transactions").AutoMigrate(&model.Transactions{})

	appRouter := router.InitRouter()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	router.InitTransactionRoutes(appRouter, transactionController)

	appRouter.Run(":8080")
}
