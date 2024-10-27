package router

import (
	"exinity-task/pkg/controller"

	"github.com/gin-gonic/gin"
)

func InitTransactionRoutes(router *gin.Engine, transactionController *controller.TransactionsController) {
	transactionRouter := router.Group("/transactions")

	transactionRouter.POST("", transactionController.Create)
}
