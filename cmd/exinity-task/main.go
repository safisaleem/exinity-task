package main

import (
	"context"
	"exinity-task/pkg/config"
	"exinity-task/pkg/controller"
	"exinity-task/pkg/gateway"
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

	// Initialize repository
	transactionRepository := repository.NewTransactionRepository(db)

	// Initialize factory
	paymentGatewayFactory := gateway.NewPaymentGatewayFactory()

	// Initialize services
	balanceService := service.NewBalanceService(transactionRepository)
	transactionService := service.NewTransactionService(transactionRepository, paymentGatewayFactory, balanceService)
	webhookService := service.NewWebhookService(transactionRepository)
	retryService := service.NewRetryService(transactionService)

	// start retry service to look for held transactions
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	retryService.Start(ctx)

	// Initialize controllers
	transactionController := controller.NewTransactionController(transactionService)
	webhookController := controller.NewWebhookController(webhookService)

	router.InitTransactionRoutes(appRouter, transactionController)
	router.InitWebhookRoutes(appRouter, webhookController)

	appRouter.Run(":8080")
}
